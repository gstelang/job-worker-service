package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"os"
	"slices"
	"sync"
	"time"

	pb "github.com/gstelang/job-worker-service/service/proto/v1"
	"github.com/gstelang/job-worker-service/worker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/peer"
)

type server struct {
	pb.UnimplementedJobWorkerServer
	logStore   *worker.JobLogStore
	jobManager *worker.JobManager
	clients    map[string][]pb.JobWorker_StreamLogsServer // jobID -> client streams
	streamsMu  sync.Mutex
}

// Convert StartJobRequest to Command
func convertToCommand(req *pb.StartJobRequest) worker.Command {
	return worker.Command{
		Name: req.GetCommandName(),
		Args: req.GetCommandArgs(),
	}
}

func (s *server) StartJob(ctx context.Context, req *pb.StartJobRequest) (*pb.StartJobResponse, error) {

	if err := authorize(ctx, "StopJob"); err != nil {
		return nil, err
	}
	command := convertToCommand(req)

	jobID, err := s.jobManager.Start(command)

	if err == nil {
		return &pb.StartJobResponse{JobId: jobID, Message: "start_job"}, nil
	} else {
		return &pb.StartJobResponse{JobId: "", Message: fmt.Sprintf("%v", err)}, err
	}
}

func (s *server) StopJob(ctx context.Context, req *pb.StopJobRequest) (*pb.StopJobResponse, error) {

	if err := authorize(ctx, "StopJob"); err != nil {
		return nil, err
	}
	jobID := req.GetJobId()

	ok, err := s.jobManager.Stop(jobID)
	if err != nil {
		return &pb.StopJobResponse{Success: false, Message: fmt.Sprintf("Failed to stop job: %v", err)}, err
	}

	return &pb.StopJobResponse{Success: ok, Message: "Stopped job!"}, nil
}

func (s *server) QueryJob(ctx context.Context, req *pb.QueryJobRequest) (*pb.QueryJobResponse, error) {
	if err := authorize(ctx, "QueryJob"); err != nil {
		return nil, err
	}
	jobID := req.GetJobId()

	jobSummary, err := s.jobManager.Query(jobID)
	if err != nil {
		return &pb.QueryJobResponse{Message: fmt.Sprintf("Failed to query job: %v", err)}, err
	}

	return &pb.QueryJobResponse{
		Status:   jobSummary.Status.String(),
		Pid:      int32(jobSummary.PID),
		ExitCode: int32(jobSummary.ExitCode),
		Signal:   int32(jobSummary.Signal),
		Message:  "Sucess!",
	}, nil
}

// Test
// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8  -> client 1, logCh (100)
// 9
// 10 -> client 2, logCh
// 11
// 12
// 13
func (s *server) StreamLogs(req *pb.StreamLogsRequest, stream pb.JobWorker_StreamLogsServer) error {

	if err := authorize(stream.Context(), "StreamLogs"); err != nil {
		return err
	}

	jobID := req.GetJobId()

	// Register this client
	s.addClient(jobID, stream)
	defer s.removeClient(jobID, stream)

	// Fetch existing logs and create/get the log channel
	existingLogs, logChannel, err := s.jobManager.Stream(stream.Context(), jobID)
	if err != nil {
		return fmt.Errorf("failed to stream logs: %w", err)
	}

	// Send existing logs to the client
	// TODO: Test this with high volume and large concurrent connections with real time updates.
	for _, logEntry := range existingLogs {
		if err := stream.Send(&pb.StreamLogsResponse{Message: logEntry}); err != nil {
			return fmt.Errorf("failed to send existing log entry: %w", err)
		}
	}

	for {
		select {
		case logEntry := <-logChannel:
			s.sendLogToClient(jobID, logEntry)
		case <-stream.Context().Done():
			return nil
		}
	}
}

func (s *server) addClient(jobID string, stream pb.JobWorker_StreamLogsServer) {
	s.streamsMu.Lock()
	defer s.streamsMu.Unlock()
	if s.clients == nil {
		s.clients = make(map[string][]pb.JobWorker_StreamLogsServer)
	}
	s.clients[jobID] = append(s.clients[jobID], stream)
}

func (s *server) removeClient(jobID string, stream pb.JobWorker_StreamLogsServer) {
	s.streamsMu.Lock()
	defer s.streamsMu.Unlock()
	if clients, ok := s.clients[jobID]; ok {
		s.clients[jobID] = slices.DeleteFunc(clients, func(currentStream pb.JobWorker_StreamLogsServer) bool {
			return currentStream == stream
		})
	}
}

func (s *server) sendLogToClient(jobID string, logEntry []byte) {
	s.streamsMu.Lock()
	clients := s.clients[jobID]
	s.streamsMu.Unlock()
	for _, client := range clients {
		if err := client.Send(&pb.StreamLogsResponse{Message: logEntry}); err != nil {
			s.removeClient(jobID, client)
		}
	}
}

func isAuthorized(role, methodName string) bool {
	roleAccess := map[string][]string{
		"StartJob":   {"Admin"},
		"StopJob":    {"Admin"},
		"QueryJob":   {"Admin", "User"},
		"StreamLogs": {"Admin", "User"},
	}

	if allowedRoles, ok := roleAccess[methodName]; ok {
		for _, r := range allowedRoles {
			if role == r {
				return true
			}
		}
	}
	return false
}

func authorize(ctx context.Context, methodName string) error {
	OIDRoleMapping := map[string]string{
		"1.3.6.1.4.1.12345.1.1": "Admin",
		"1.3.6.1.4.1.12345.1.2": "User",
	}

	peer, ok := peer.FromContext(ctx)
	if !ok {
		return fmt.Errorf("no peer found")
	}

	tlsInfo, ok := peer.AuthInfo.(credentials.TLSInfo)
	if !ok {
		return fmt.Errorf("unexpected peer transport credentials")
	}

	for _, cert := range tlsInfo.State.PeerCertificates {
		for _, ext := range cert.Extensions {
			oid := ext.Id.String()
			if role, exists := OIDRoleMapping[oid]; exists {
				if isAuthorized(role, methodName) {
					return nil
				}
				return fmt.Errorf("unauthorized: permission denied for role %s for action %s", role, methodName)
			}
		}
	}
	return fmt.Errorf("unknown role")
}

func main() {
	// Load server's certificate and private key
	certFile := "./certs/server/server.crt"
	certKey := "./certs/server/server.key"
	caFile := "./certs/ca/ca.crt"

	// Load CA certificate
	caCert, err := os.ReadFile(caFile)
	if err != nil {
		log.Fatalf("Failed to read CA certificate: %v", err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Load server's certificate and private key
	cert, err := tls.LoadX509KeyPair(certFile, certKey)
	if err != nil {
		panic(err)
	}

	// Set up TLS credentials
	creds := credentials.NewTLS(&tls.Config{
		MinVersion:   tls.VersionTLS13, // Enforce TLS 1.3
		MaxVersion:   tls.VersionTLS13, // Enforce TLS 1.3
		Certificates: []tls.Certificate{cert},
		ClientCAs:    caCertPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	})

	// Create a listener on TCP port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	// Initialize logStore and jobManager
	logStore := worker.NewJobLogStore()
	jobManager := worker.NewJobManager(logStore)

	// Create a gRPC server object with TLS credentials
	s := grpc.NewServer(
		grpc.Creds(creds),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 30 * time.Minute, // Disconnect after 30 minutes of inactivity
			Time:              10 * time.Minute, // Send keepalive pings every 10 minutes
			Timeout:           20 * time.Second, // Wait 20 seconds for keepalive ping ack
		}),
	)
	pb.RegisterJobWorkerServer(s, &server{
		logStore:   logStore,
		jobManager: jobManager,
		clients:    make(map[string][]pb.JobWorker_StreamLogsServer),
	})

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
