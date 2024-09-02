package worker

// TODO: Separated into 2 interfaces.

// JobLogger defines operations for logging job-related information.
type JobLogger interface {
	AddJob(jobID string)
	AddLog(jobID string, log []byte)
	GetLogs(jobID string) ([][]byte, error)
	GetOrCreateLogChannel(jobID string) chan []byte
}

// JobDetailsManagement defines operations for managing job details.
type JobDetailsManagement interface {
	GetJobDetails(jobID string) (JobDetails, error)
	UpdateJobDetails(jobID string, signal int, exitCode int)
	AddProcessId(jobID string, pid int)
	UpdateJobStatus(jobID string, status JobStatus)
	GetJobSummary(jobID string) (JobSummary, error)
}

// JobSummary represents a summary of a job's details.
type JobSummary struct {
	Status   JobStatus
	PID      int
	ExitCode int
	Signal   int
}

// JobDetails contains job summary details and the logs.
type JobDetails struct {
	JobSummary
	Logs [][]byte // store logs in their raw byte form, accommodating any non-UTF-8 characters.
}
