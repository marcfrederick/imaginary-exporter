package main

import (
	"encoding/json"
	"net/http"
)

type ImaginaryMetrics struct {
	Uptime               float64 `json:"uptime"`
	AllocatedMemory      float64 `json:"allocatedMemory"`
	TotalAllocatedMemory float64 `json:"totalAllocatedMemory"`
	Goroutines           float64 `json:"goroutines"`
	CompletedGCCycles    float64 `json:"completedGCCycles"`
	CPUs                 float64 `json:"cpus"`
	MaxHeapUsage         float64 `json:"maxHeapUsage"`
	HeapInUse            float64 `json:"heapInUse"`
	ObjectsInUse         float64 `json:"objectsInUse"`
	OSMemoryObtained     float64 `json:"OSMemoryObtained"`
}

// read json from the given url.
func NewFromURL(url string) (*ImaginaryMetrics, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	metrics := &ImaginaryMetrics{}
	err = json.NewDecoder(resp.Body).Decode(metrics)
	return metrics, err
}
