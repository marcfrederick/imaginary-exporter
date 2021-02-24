package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const namePrefix = "imaginary_"

var (
	uptime               = newGauge("uptime", "The current uptime.")
	allocatedMemory      = newGauge("allocated_memory", "The currently allocated memory.")
	totalAllocatedMemory = newGauge("allocated_memory_total", "The total allocated memory.")
	goroutines           = newGauge("goroutines", "The number of running goroutines.")
	completedGCCycles    = newGauge("gc_cycles_total", "The number of garbage collection cycles.")
	cpus                 = newGauge("cpus_total", "The number of CPUs available.")
	maxHeapUsage         = newGauge("heap_usage_max", "The maximum heap usage.")
	heapInUse            = newGauge("heap_usage", "The current heap usage.")
	objectsInUse         = newGauge("objects", "The number of currently used objects.")
	oSMemoryObtained     = newGauge("os_memory", "The amount of OS memory obtained.")
)

func newGauge(name, help string) prometheus.Gauge {
	return promauto.NewGauge(prometheus.GaugeOpts{
		Name: namePrefix + name,
		Help: help,
	})
}
