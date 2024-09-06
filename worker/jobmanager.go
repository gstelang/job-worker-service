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
	// Set I/O weight to 500 (range is 1-1000, 100 is default)
	DefaultDiskIOWeight = 500
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
	if err == nil {
		return 0, 0
	}

	exitError, ok := err.(*exec.ExitError)

	if !ok {
		fmt.Printf("Error waiting for command: %v\n", err)
		return 0, 0
	}

	status, ok := exitError.Sys().(syscall.WaitStatus)
	if !ok {
		fmt.Printf("Error getting wait status: %v\n", err)
		return 0, 0
	}

	if status.Signaled() {
		signal = int(status.Signal())
		fmt.Printf("Process was killed by signal: %d\n", signal)
		return signal, 0
	}

	exitCode = status.ExitStatus()
	return 0, exitCode
}

func readAndLogPipe(jobID string, pipe io.ReadCloser, logger JobLogger) {
	defer pipe.Close()

	// Adjust
	buffer := make([]byte, 4096)
	for {
		n, err := pipe.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error reading from pipe: %v\n", err)
			break
		}
		if n > 0 {
			logger.AddLog(jobID, buffer[:n])
		}
	}
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

	jm.resource.SetLimits(jobID, DefaultCPULimit, DefaultMemoryLimit, DefaultDiskIOWeight)

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

	// Start the process in the cgroup
	err = jm.resource.StartProcessInCgroup(jobID, cmd)
	if err != nil {
		return "", fmt.Errorf("error starting process in cgroup: %w", err)
	}
	// Record the process ID
	jm.details.AddProcessId(jobID, cmd.Process.Pid)

	// Continuously read from the command's stdout
	go func() {
		// read stdout in a separate goroutine
		go readAndLogPipe(jobID, stdout, jm.logger)

		// read stderr in a separate goroutine
		go readAndLogPipe(jobID, stderr, jm.logger)

		defer func() {
			signal, exitCode := getJobEndStatus(cmd)
			status := StatusExited
			if signal != 0 {
				status = StatusSignaled
			}
			jm.details.UpdateJobStatus(jobID, status)
			jm.details.UpdateJobDetails(jobID, signal, exitCode)
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
