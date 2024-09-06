package worker

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"sync"
	"syscall"
)

type ResourceLimits struct {
	CPULimit     int
	MemoryLimit  int
	DiskIOWeight int
}

type ResourceManager struct {
	limits map[string]ResourceLimits
	mu     sync.RWMutex
}

func NewResourceManager() *ResourceManager {
	return &ResourceManager{
		limits: make(map[string]ResourceLimits),
	}
}

func (rm *ResourceManager) SetLimits(jobID string, cpuLimit int, memoryLimit int, diskIOWeight int) error {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	rm.limits[jobID] = ResourceLimits{
		CPULimit:     cpuLimit,
		MemoryLimit:  memoryLimit,
		DiskIOWeight: diskIOWeight,
	}
	return nil
}

func (rm *ResourceManager) GetLimits(jobID string) (int, int, int, error) {
	rm.mu.RLock()
	defer rm.mu.RUnlock()
	limits, exists := rm.limits[jobID]
	if !exists {
		return 0, 0, 0, fmt.Errorf("job ID %s not found", jobID)
	}
	return limits.CPULimit, limits.MemoryLimit, limits.DiskIOWeight, nil
}

func (rm *ResourceManager) CreateCgroup(jobID string) error {
	cgroupName := fmt.Sprintf("job_%s", jobID)
	cgroupPath := filepath.Join("/sys/fs/cgroup", cgroupName)
	err := os.Mkdir(cgroupPath, 0755)
	if err != nil {
		if !os.IsExist(err) {
			return fmt.Errorf("error creating cgroup: %w", err)
		}
		// Cgroup already exists, continue
	}

	return nil
}

// Additionally we might need a cronjob to clean up cgroups
// Example: rmdir job_* inside /sys/fs/cgroup folder
func (rm *ResourceManager) CleanupCgroup(jobID string) error {
	cgroupName := fmt.Sprintf("job_%s", jobID)
	cgroupPath := filepath.Join("/sys/fs/cgroup", cgroupName)
	err := os.RemoveAll(cgroupPath)
	if err != nil {
		return fmt.Errorf("error removing cgroup: %w", err)
	}
	return nil
}

func (rm *ResourceManager) StartProcessInCgroup(jobID string, cmd *exec.Cmd) error {
	cgroupName := fmt.Sprintf("job_%s", jobID)
	cgroupPath := filepath.Join("/sys/fs/cgroup", cgroupName)

	cgroupDir := filepath.Join("/sys/fs/cgroup", cgroupName)
	cgroupFD, err := os.Open(cgroupDir)
	if err != nil {
		return fmt.Errorf("error opening cgroup directory: %w", err)
	}
	defer cgroupFD.Close()

	// Set resource limits for the job
	cpuLimit, memoryLimit, diskIOWeight, err := rm.GetLimits(jobID)
	if err != nil {
		return fmt.Errorf("error getting resource limits: %w", err)
	}

	// set memory limit
	err = os.WriteFile(filepath.Join(cgroupPath, "memory.max"), []byte(strconv.Itoa(memoryLimit)), 0644)
	if err != nil {
		return fmt.Errorf("error setting memory limit: %w", err)
	}

	// Set CPU limit
	err = os.WriteFile(filepath.Join(cgroupPath, "cpu.max"), []byte(fmt.Sprintf("%d 100000", cpuLimit)), 0644)
	if err != nil {
		return fmt.Errorf("error setting CPU limit: %w", err)
	}

	// Set disk I/O limit
	err = os.WriteFile(filepath.Join(cgroupPath, "io.weight"), []byte(strconv.Itoa(diskIOWeight)), 0644)
	if err != nil {
		return fmt.Errorf("failed to set disk I/O limit: %w", err)
	}

	// Set the cgroup for the new process before starting it
	cmd.SysProcAttr = &syscall.SysProcAttr{
		UseCgroupFD: true,
		CgroupFD:    int(cgroupFD.Fd()),
	}

	// Start the command
	err = cmd.Start()
	if err != nil {
		return fmt.Errorf("error starting command: %w", err)
	}

	return nil
}
