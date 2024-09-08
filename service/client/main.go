package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/gstelang/job-worker-service/service/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	// Define command-line flags
	command := flag.String("command", "", "Command to execute (start, stop, query, stream)")
	jobID := flag.String("jobid", "", "Job ID for stop, query, or stream")

	flag.Parse()

	// Load client certificates
	cert, err := tls.LoadX509KeyPair("./certs/client/admin/admin.crt", "./certs/client/admin/admin.key")
	if err != nil {
		log.Fatalf("failed to load client certificate: %v", err)
	}

	// Load CA certificate
	caCert, err := os.ReadFile("./certs/ca/ca.crt")
	if err != nil {
		log.Fatalf("failed to read CA certificate: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Set up TLS credentials
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	})

	// Connect to the gRPC server
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewJobWorkerClient(conn)

	// Handle commands
	switch *command {
	case "start":
		remainingArgs := flag.Args()

		commandName := remainingArgs[0]
		commandArgs := remainingArgs[1:]

		startJob(client, commandName, commandArgs)

	case "stop":
		stopJob(client, *jobID)
	case "query":
		queryJob(client, *jobID)
	case "stream":
		streamLogs(client, *jobID)
	default:
		log.Fatalf("Unknown command: %s", *command)
	}
}

func startJob(client pb.JobWorkerClient, commandName string, args []string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &pb.StartJobRequest{
		CommandName: commandName,
		CommandArgs: args,
	}

	res, err := client.StartJob(ctx, req)
	if err != nil {
		log.Fatalf("Failed to start job: %v", err)
	}

	fmt.Printf("Job started with ID: %s\n", res.JobId)
}

func stopJob(client pb.JobWorkerClient, jobID string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &pb.StopJobRequest{
		JobId: jobID,
	}

	res, err := client.StopJob(ctx, req)
	if err != nil {
		log.Fatalf("Failed to stop job: %v", err)
	}

	fmt.Printf("Job stopped: %v\n", res.Message)
}

func queryJob(client pb.JobWorkerClient, jobID string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &pb.QueryJobRequest{
		JobId: jobID,
	}

	res, err := client.QueryJob(ctx, req)
	if err != nil {
		log.Fatalf("Failed to query job: %v", err)
	}

	resJSON, _ := json.MarshalIndent(res, "", "  ")
	fmt.Printf("Job query result: %s\n", resJSON)
}

func streamLogs(client pb.JobWorkerClient, jobID string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	req := &pb.StreamLogsRequest{
		JobId: jobID,
	}

	stream, err := client.StreamLogs(ctx, req)
	if err != nil {
		log.Fatalf("Failed to stream logs: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error receiving log: %v", err)
		}

		// printing string representation.
		// adjust if binary
		fmt.Printf("Log entry: %s\n", res.Message)
	}
}
