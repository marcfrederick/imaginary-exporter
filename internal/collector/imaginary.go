package collector

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/marcfrederick/imaginary-exporter/internal/imaginary"
)

// ImaginaryCollector collects metrics from a given Imaginary instance.
type ImaginaryCollector struct {
	client                     *imaginary.Client
	uptimeMetric               *prometheus.Desc
	allocatedMemoryMetric      *prometheus.Desc
	totalAllocatedMemoryMetric *prometheus.Desc
	goroutinesMetric           *prometheus.Desc
	completedGCCyclesMetric    *prometheus.Desc
	cpusMetric                 *prometheus.Desc
	maxHeapUsageMetric         *prometheus.Desc
	heapInUseMetric            *prometheus.Desc
	objectsInUseMetric         *prometheus.Desc
	oSMemoryObtainedMetric     *prometheus.Desc
}

var _ prometheus.Collector = (*ImaginaryCollector)(nil)

// NewImaginaryCollector creates a new ImaginaryCollector and initializes it.
func NewImaginaryCollector(client *imaginary.Client) *ImaginaryCollector {
	return &ImaginaryCollector{
		client:                     client,
		uptimeMetric:               prometheus.NewDesc("imaginary_uptime", "The current uptime.", nil, nil),
		allocatedMemoryMetric:      prometheus.NewDesc("imaginary_allocated_memory", "The currently allocated memory.", nil, nil),
		totalAllocatedMemoryMetric: prometheus.NewDesc("imaginary_allocated_memory_total", "The total allocated memory.", nil, nil),
		goroutinesMetric:           prometheus.NewDesc("imaginary_goroutines", "The number of running goroutines.", nil, nil),
		completedGCCyclesMetric:    prometheus.NewDesc("imaginary_gc_cycles_total", "The number of garbage collection cycles.", nil, nil),
		cpusMetric:                 prometheus.NewDesc("imaginary_cpus_total", "The number of CPUs available.", nil, nil),
		maxHeapUsageMetric:         prometheus.NewDesc("imaginary_heap_usage_max", "The maximum heap usage.", nil, nil),
		heapInUseMetric:            prometheus.NewDesc("imaginary_heap_usage", "The current heap usage.", nil, nil),
		objectsInUseMetric:         prometheus.NewDesc("imaginary_objects", "The number of currently used objects.", nil, nil),
		oSMemoryObtainedMetric:     prometheus.NewDesc("imaginary_os_memory", "The amount of OS memory obtained.", nil, nil),
	}
}

func (c *ImaginaryCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.uptimeMetric
	ch <- c.uptimeMetric
	ch <- c.allocatedMemoryMetric
	ch <- c.totalAllocatedMemoryMetric
	ch <- c.goroutinesMetric
	ch <- c.completedGCCyclesMetric
	ch <- c.cpusMetric
	ch <- c.maxHeapUsageMetric
	ch <- c.heapInUseMetric
	ch <- c.objectsInUseMetric
	ch <- c.oSMemoryObtainedMetric
}

func (c *ImaginaryCollector) Collect(ch chan<- prometheus.Metric) {
	res, err := c.client.GetHealthStats()
	if err != nil {
		log.Printf("error getting metrics: %s", err)
		return
	}
	ch <- prometheus.MustNewConstMetric(c.uptimeMetric, prometheus.CounterValue, float64(res.Uptime))
	ch <- prometheus.MustNewConstMetric(c.allocatedMemoryMetric, prometheus.GaugeValue, res.AllocatedMemory)
	ch <- prometheus.MustNewConstMetric(c.totalAllocatedMemoryMetric, prometheus.GaugeValue, res.TotalAllocatedMemory)
	ch <- prometheus.MustNewConstMetric(c.goroutinesMetric, prometheus.GaugeValue, float64(res.Goroutines))
	ch <- prometheus.MustNewConstMetric(c.completedGCCyclesMetric, prometheus.GaugeValue, float64(res.CompletedGCCycles))
	ch <- prometheus.MustNewConstMetric(c.cpusMetric, prometheus.GaugeValue, float64(res.CPUs))
	ch <- prometheus.MustNewConstMetric(c.maxHeapUsageMetric, prometheus.GaugeValue, res.MaxHeapUsage)
	ch <- prometheus.MustNewConstMetric(c.heapInUseMetric, prometheus.GaugeValue, res.HeapInUse)
	ch <- prometheus.MustNewConstMetric(c.objectsInUseMetric, prometheus.GaugeValue, float64(res.ObjectsInUse))
	ch <- prometheus.MustNewConstMetric(c.oSMemoryObtainedMetric, prometheus.GaugeValue, res.OSMemoryObtained)
}
