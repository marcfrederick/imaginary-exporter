package imaginary

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Client represents a client for accessing Imaginary services.
type Client struct {
	BaseURL string
}

// NewClient creates a new Client instance with the specified base URL.
func NewClient(baseURL string) *Client {
	return &Client{BaseURL: baseURL}
}

// HealthStats represents the health metrics provided by the Imaginary service.
type HealthStats struct {
	Uptime               int64   `json:"uptime"`
	AllocatedMemory      float64 `json:"allocatedMemory"`
	TotalAllocatedMemory float64 `json:"totalAllocatedMemory"`
	Goroutines           int     `json:"goroutines"`
	CompletedGCCycles    uint32  `json:"completedGCCycles"`
	CPUs                 int     `json:"cpus"`
	MaxHeapUsage         float64 `json:"maxHeapUsage"`
	HeapInUse            float64 `json:"heapInUse"`
	ObjectsInUse         uint64  `json:"objectsInUse"`
	OSMemoryObtained     float64 `json:"OSMemoryObtained"`
}

// GetHealthStats retrieves the health statistics from the Imaginary service.
func (c *Client) GetHealthStats() (*HealthStats, error) {
	resp, err := http.Get(c.BaseURL + "/health")
	if err != nil {
		return nil, fmt.Errorf("error getting health stats: %w", err)
	}
	defer resp.Body.Close()

	healthStats := &HealthStats{}
	if err = json.NewDecoder(resp.Body).Decode(healthStats); err != nil {
		return nil, fmt.Errorf("error decoding health stats: %w", err)
	}
	return healthStats, nil
}
