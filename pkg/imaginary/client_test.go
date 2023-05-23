package imaginary_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"imaginary-exporter/pkg/imaginary"
)

func TestClient_GetHealthStats(t *testing.T) {
	healthJSON, _ := os.ReadFile("testdata/health.json")
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/health" {
			t.Errorf("unexpected path: want = /health, got = %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(healthJSON)
	}))
	defer s.Close()

	c := imaginary.NewClient(s.URL)
	stats, err := c.GetHealthStats()
	if err != nil {
		t.Fatalf("unexpected error: got = %s", err)
	}

	if stats.Uptime != 12 {
		t.Errorf("unexpected uptime: want = 12, got = %d", stats.Uptime)
	}
	if stats.CPUs != 5 {
		t.Errorf("unexpected CPUs: want = 5, got = %d", stats.CPUs)
	}
	if stats.OSMemoryObtained != 69.45 {
		t.Errorf("unexpected OSMemoryObtained: want = 69.45, got = %d", stats.CPUs)
	}
}
