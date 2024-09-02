package worker

import (
	"fmt"
	"sync"
)

type ResourceController interface {
	SetLimits(jobID string, cpuLimit int, memoryLimit int, diskIOLimit int) error
	GetLimits(jobID string) (cpuLimit int, memoryLimit int, diskIOLimit int, err error)
}

type ResourceLimits struct {
	CPULimit    int
	MemoryLimit int
	DiskIOLimit int
}

type ResourceManager struct {
	limits map[string]ResourceLimits
	mu     sync.RWMutex
}

func NewResourceController() *ResourceManager {
	return &ResourceManager{
		limits: make(map[string]ResourceLimits),
	}
}

func (rm *ResourceManager) SetLimits(jobID string, cpuLimit int, memoryLimit int, diskIOLimit int) error {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	rm.limits[jobID] = ResourceLimits{cpuLimit, memoryLimit, diskIOLimit}
	// TODO: enforce using cgroups
	return nil
}

func (rm *ResourceManager) GetLimits(jobID string) (int, int, int, error) {
	rm.mu.RLock()
	defer rm.mu.RUnlock()
	limits, exists := rm.limits[jobID]
	if !exists {
		return 0, 0, 0, fmt.Errorf("job ID %s not found", jobID)
	}
	return limits.CPULimit, limits.MemoryLimit, limits.DiskIOLimit, nil
}
