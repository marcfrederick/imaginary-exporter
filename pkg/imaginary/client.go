package imaginary

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	BaseURL string
}

func NewClient(baseURL string) *Client {
	return &Client{BaseURL: baseURL}
}

// HealthStats represents the metrics provided by Imaginary.
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

// GetHealthStats reads metrics.
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
