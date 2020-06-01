package cpustatcollector

import (
	"github.com/prometheus/client_golang/prometheus"
)

// CPUMetrics defines the CPU metrics going to expose from
// this exporter.
type CPUMetrics struct {
	cupIdle        prometheus.Gauge
	cpuTotal       prometheus.Gauge
	cpuUtilization prometheus.Gauge
}

// New craetes a new CPU metrics instance
func New() *CPUMetrics {
	// craete a new CPU metrics instance
	return &CPUMetrics{
		cupIdle:        prometheus.NewGauge(),
		cpuTotal:       prometheus.NewGauge(),
		cpuUtilization: prometheus.NewGauge(),
	}
}

// Init initializes the metrics scraping
func Init() {
	// initialize the metrics scraping
}

// Collect collects metrics from the client periodically
func Collect() {
	// collects metrics from the client periodically
}
