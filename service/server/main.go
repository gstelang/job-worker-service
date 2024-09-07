package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net"
	"os"
	"time"

	pb "github.com/gstelang/job-worker-service/service/proto/v1"
	"github.com/gstelang/job-worker-service/worker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
)

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
