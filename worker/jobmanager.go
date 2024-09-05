package worker

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"

	"github.com/google/uuid"
)

const (
	DefaultCPULimit    = 50000     // 50% of 1 core
	DefaultMemoryLimit = 104857600 // 2^20 bytes = 100 MB.
	// Set I/O limit to 50MB read and 20MB write per second for a device with major:minor 8:0
	DefaultDiskIOLimit = "8:0 52428800 20971520"
)

// JobManager is composed of jobLogger and implements the controller interface
type JobManager struct {
	logger   JobLogger
	details  JobDetailsManagement
	resource ResourceController
}

func NewJobManager(store *JobLogStore) *JobManager {
	return &JobManager{
		logger:   store,
		details:  store,
		resource: NewResourceManager(),
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

	// Create cgroup for the job
	err = jm.resource.CreateCgroup(jobID)
	if err != nil {
		return "", fmt.Errorf("error creating cgroup: %w", err)
	}

	defer func() {
		if err != nil {
			jm.resource.CleanupCgroup(jobID)
		}
	}()

	jm.resource.SetLimits(jobID, DefaultCPULimit, DefaultMemoryLimit, DefaultDiskIOLimit)

	// Create the command and attach stdout pipe
	cmd := exec.Command(command.Name, command.Args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		jm.logger.AddLog(jobID, []byte(fmt.Sprintf("Failed to create stdout pipe: %v", err)))
		return "", fmt.Errorf("failed to create stdout pipe: %w", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", fmt.Errorf("error creating stderr pipe: %w", err)
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		jm.logger.AddLog(jobID, []byte(fmt.Sprintf("Failed to start the command: %v", err)))
		return "", fmt.Errorf("failed to start command: %w", err)
	}

	// Start the process in the cgroup
	err = jm.resource.StartProcessInCgroup(jobID, cmd)
	if err != nil {
		return "", fmt.Errorf("error starting process in cgroup: %w", err)
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

		// read stdout in chunks
		go func() {
			buffer := make([]byte, 1024)
			for {
				n, err := stdout.Read(buffer)
				if err != nil {
					if err == io.EOF {
						break
					}
					fmt.Printf("Error reading from stdout: %v\n", err)
					break
				}
				jm.logger.AddLog(jobID, buffer[:n])
			}
		}()

		// read stderr in chunks
		go func() {
			buffer := make([]byte, 1024)
			for {
				n, err := stderr.Read(buffer)
				if err != nil {
					if err == io.EOF {
						break
					}
					fmt.Printf("Error reading from stderr: %v\n", err)
					break
				}
				jm.logger.AddLog(jobID, buffer[:n])
			}
		}()

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
