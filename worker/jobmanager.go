package worker

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/google/uuid"
)

// JobManager is composed of jobLogger and implements the controller interface
type JobManager struct {
	logger  JobLogger
	details JobDetailsManagement
	// TODO: implementation
	// resource ResourceController
}

func NewJobManager(store *JobLogStore) *JobManager {
	return &JobManager{
		logger:  store,
		details: store,
	}
}

func getJobEndStatus(cmd *exec.Cmd) (signal, exitCode int) {
	err := cmd.Wait()
	signal = 0
	exitCode = 0
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			if status, ok := exitError.Sys().(syscall.WaitStatus); ok {
				if status.Signaled() {
					signal = int(status.Signal())
					fmt.Printf("Process was killed by signal: %d\n", status.Signal())
				} else {
					exitCode = status.ExitStatus()
					fmt.Printf("Process exited with status code: %d\n", status.ExitStatus())
				}
			}
		} else {
			fmt.Printf("Error waiting for command: %v\n", err)
		}
	} else {
		exitCode = 0
	}

	return signal, exitCode
}

// Start starts a command and logs the output
func (jm *JobManager) Start(command Command) (jobID string, err error) {
	jobID = generateUUID()
	jm.logger.AddJob(jobID)
	jm.logger.AddLog(jobID, []byte(fmt.Sprintf("Starting job with ID: %s", jobID)))

	// Create the command and attach stdout pipe
	cmd := exec.Command(command.Name, command.Args...)
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		jm.logger.AddLog(jobID, []byte(fmt.Sprintf("Failed to create stdout pipe: %v", err)))
		return "", fmt.Errorf("failed to create stdout pipe: %w", err)
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		jm.logger.AddLog(jobID, []byte(fmt.Sprintf("Failed to start the command: %v", err)))
		return "", fmt.Errorf("failed to start command: %w", err)
	}

	// Record the process ID
	jm.details.AddProcessId(jobID, cmd.Process.Pid)

	// Continuously read from the command's stdout
	go func() {
		defer func() {
			signal, exitCode := getJobEndStatus(cmd)
			status := StatusExited
			if signal != 0 {
				status = StatusSignaled
			}
			jm.details.UpdateJobStatus(jobID, status)
			jm.details.UpdateJobDetails(jobID, signal, exitCode)
		}()

		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			line := scanner.Bytes()
			// Copy line to avoid reuse
			jm.logger.AddLog(jobID, append([]byte(nil), line...))
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading from command output: %v", err)
		}
	}()

	return jobID, nil
}

// Query retrieves the logs for a given jobID
func (jm *JobManager) Query(jobID string) (JobSummary, error) {

	// Retrieve job details
	jobSummary, err := jm.details.GetJobSummary(jobID)
	if err != nil {
		return jobSummary, err
	}
	return jobSummary, nil
}

// Stop stops a running job
func (store *JobManager) Stop(jobID string) (bool, error) {
	// Retrieve job details
	jobDetails, err := store.details.GetJobDetails(jobID)
	if err != nil {
		return false, err
	}

	// Send termination signal to the process
	process, err := os.FindProcess(jobDetails.PID)
	if err != nil {
		return false, fmt.Errorf("failed to find process: %v", err)
	}

	if err := process.Signal(syscall.SIGTERM); err != nil {
		return false, fmt.Errorf("failed to stop process: %v", err)
	}

	// Update job status
	store.details.UpdateJobStatus(jobID, StatusSignaled)
	return true, nil
}

// Stream retrieves existing logs and returns a channel for real-time logs
func (jm *JobManager) Stream(ctx context.Context, jobID string) ([][]byte, chan []byte, error) {
	// Retrieve existing logs
	existingLogs, err := jm.logger.GetLogs(jobID)
	if err != nil {
		return nil, nil, err
	}

	// Create or get the log channel for real-time logs
	logChannel := jm.logger.GetOrCreateLogChannel(jobID)

	return existingLogs, logChannel, err
}

// generateUUID generates a unique identifier for jobs
func generateUUID() string {
	return uuid.New().String()
}
