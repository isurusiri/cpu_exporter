package cpustatcollector

import (
	"github.com/isurusiri/cpu_exporter/cpuclient"
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
	cpuMetrics *CPUMetrics
	cpuStats   *cpuclient.CPUStat
}

func newCPUMetrics() *CPUMetrics {
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

// New craetes a new CPU collector instance
func New() *CPUCollector {
	cpuMetrics := newCPUMetrics()
	cpuStats := cpuclient.New()

	return &CPUCollector{
		cpuMetrics: cpuMetrics,
		cpuStats:   cpuStats,
	}
}

// Init initializes the metrics scraping
func (cpuCollector *CPUCollector) Init() {
	prometheus.MustRegister(cpuCollector.cpuMetrics.cpuTotal)
	prometheus.MustRegister(cpuCollector.cpuMetrics.cpuUtilization)
	prometheus.MustRegister(cpuCollector.cpuMetrics.cupIdle)
}

// Collect collects metrics from the client periodically
func (cpuCollector *CPUCollector) Collect() {
	cpuCollector.cpuStats.GetCPUStats()

	cpuCollector.cpuMetrics.cpuTotal.Set(float64(cpuCollector.cpuStats.Total))
	cpuCollector.cpuMetrics.cupIdle.Set(float64(cpuCollector.cpuStats.Idle))
	cpuCollector.cpuMetrics.cpuUtilization.Set(cpuCollector.cpuStats.Utilization)
}
