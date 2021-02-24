package main

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"net/http"
	"net/url"
	"time"
)

const defaultFetchRate = 15 * time.Second

func readMetrics(url string, fetchRate *time.Duration) {
	go func() {
		for {
			log.Info("refreshing metrics")
			metrics, err := NewFromURL(url)
			if err != nil {
				log.Errorf("failed to fetch metrics from '%s': %+v", url, err)
			} else {
				uptime.Set(metrics.Uptime)
				allocatedMemory.Set(metrics.AllocatedMemory)
				totalAllocatedMemory.Set(metrics.TotalAllocatedMemory)
				goroutines.Set(metrics.Goroutines)
				completedGCCycles.Set(metrics.CompletedGCCycles)
				cpus.Set(metrics.CPUs)
				maxHeapUsage.Set(metrics.MaxHeapUsage)
				heapInUse.Set(metrics.HeapInUse)
				objectsInUse.Set(metrics.ObjectsInUse)
				oSMemoryObtained.Set(metrics.OSMemoryObtained)
			}
			time.Sleep(*fetchRate)
		}
	}()
}

// Ensure the entered url is valid.
func isURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func main() {
	imaginaryURL := flag.String("url", "", "url of the imaginary instance")
	fetchRate := flag.Duration("fetch-rate", defaultFetchRate, "interval in milliseconds in which to fetch metrics")
	flag.Parse()

	if !isURL(*imaginaryURL) {
		log.Fatalf("Given URL '%s' is invalid.", *imaginaryURL)
	}

	healthURL := *imaginaryURL + "/health"
	readMetrics(healthURL, fetchRate)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
