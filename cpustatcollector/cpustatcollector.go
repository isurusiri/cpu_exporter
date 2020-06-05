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

// CPUCollector represents CPU Collector object
type CPUCollector struct {
	cpuMetrics CPUMetrics
	// cpuClient  cpu_client.CPUStat
}

// New craetes a new CPU metrics instance
func New() *CPUMetrics {
	return &CPUMetrics{
		cupIdle: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "cpu_exporter_cpu_idle_time",
			Help: "Shows the CPU idle time",
		}),
		cpuTotal: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "cpu_exporter_cpu_total_usage",
			Help: "Shows the total CPU availability",
		}),
		cpuUtilization: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "cpu_exporter_cpu_utilization",
			Help: "Shows the total CPU utilization",
		}),
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
