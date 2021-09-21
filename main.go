package main

import (
	"encoding/json"
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"net/url"
)

// ImaginaryMetrics represents the metrics provided by Imaginary.
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

// getMetrics reads metrics from the given Imaginary url.
func getMetrics(url string) (*ImaginaryMetrics, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	metrics := &ImaginaryMetrics{}
	err = json.NewDecoder(resp.Body).Decode(metrics)
	return metrics, err
}

// ImaginaryCollector collects metrics from a given Imaginary instance.
type ImaginaryCollector struct {
	url                        string
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

// newImaginaryCollector creates a new ImaginaryCollector and initializes it.
func newImaginaryCollector(url string) *ImaginaryCollector {
	return &ImaginaryCollector{
		url:                        url,
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

func (c ImaginaryCollector) Describe(ch chan<- *prometheus.Desc) {
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

func (c ImaginaryCollector) Collect(ch chan<- prometheus.Metric) {
	res, err := getMetrics(c.url)
	if err != nil {
		log.Printf("failed to fetch metrics from '%s': %+v", c.url, err)
		return
	}
	ch <- prometheus.MustNewConstMetric(c.uptimeMetric, prometheus.CounterValue, res.Uptime)
	ch <- prometheus.MustNewConstMetric(c.allocatedMemoryMetric, prometheus.GaugeValue, res.AllocatedMemory)
	ch <- prometheus.MustNewConstMetric(c.totalAllocatedMemoryMetric, prometheus.GaugeValue, res.TotalAllocatedMemory)
	ch <- prometheus.MustNewConstMetric(c.goroutinesMetric, prometheus.GaugeValue, res.Goroutines)
	ch <- prometheus.MustNewConstMetric(c.completedGCCyclesMetric, prometheus.GaugeValue, res.CompletedGCCycles)
	ch <- prometheus.MustNewConstMetric(c.cpusMetric, prometheus.GaugeValue, res.CPUs)
	ch <- prometheus.MustNewConstMetric(c.maxHeapUsageMetric, prometheus.GaugeValue, res.MaxHeapUsage)
	ch <- prometheus.MustNewConstMetric(c.heapInUseMetric, prometheus.GaugeValue, res.HeapInUse)
	ch <- prometheus.MustNewConstMetric(c.objectsInUseMetric, prometheus.GaugeValue, res.ObjectsInUse)
	ch <- prometheus.MustNewConstMetric(c.oSMemoryObtainedMetric, prometheus.GaugeValue, res.OSMemoryObtained)
}

// isURL checks whether the given string is a valid url
func isURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func main() {
	addr := flag.String("addr", ":8080", "address to listen on")
	imaginaryURL := flag.String("url", "", "url of the imaginary instance")
	flag.Parse()
	if !isURL(*imaginaryURL) {
		log.Fatalf("Given URL '%s' is invalid.", *imaginaryURL)
	}

	c := newImaginaryCollector(*imaginaryURL + "/health")
	if err := prometheus.Register(c); err != nil {
		log.Fatalf("failed to register imaginary collector: %+v", err)
	}

	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal(err)
	}
}
