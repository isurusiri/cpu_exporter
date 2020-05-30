package cpuclient

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

// CPUStat contains the metrics and other information that is going
// to exposed from the exporter.
type CPUStat struct {
	Idle                     uint64
	Total                    uint64
	LastIdle                 uint64
	LastTotal                uint64
	Utilization              float64
	ReadAt                   time.Time
	DurationOfTheUtilization string
}

func getCPUStats() (idle, total uint64) {
	contents, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		return
	}

	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if fields[0] == "cpu" {
			numFields := len(fields)
			for i := 0; i < numFields; i++ {
				val, err := strconv.ParseUint(fields[i], 10, 64)
				if err != nil {
					fmt.Println("Error: ", i, fields[i], err)
				}
				total += val
				if i == 4 {
					idle = val
				}
			}
			return
		}
	}
	return
}

// New create and return an instance of CPUStat
func New() *CPUStat {
	currentIdle, currentTotal := getCPUStats()

	cpuStat := &CPUStat{
		Idle:                     currentIdle,
		Total:                    currentTotal,
		LastIdle:                 0,
		LastTotal:                0,
		Utilization:              float64(0),
		ReadAt:                   time.Now(),
		DurationOfTheUtilization: "0",
	}

	return cpuStat
}

// GetCPUStats returns current CPU Stats
func (cpuStats *CPUStat) GetCPUStats() {
	currentIdle, currentTotal := getCPUStats()

	// switch previous current stats to last stats
	cpuStats.LastIdle = cpuStats.Idle
	cpuStats.LastTotal = cpuStats.Total

	// assigns current stats
	cpuStats.Idle = currentIdle
	cpuStats.Total = currentTotal

	cpuStats.Utilization = getCPUUtilization(cpuStats.Idle, cpuStats.Total, cpuStats.LastIdle, cpuStats.LastTotal)

	cpuStats.DurationOfTheUtilization = calcDuration(cpuStats.ReadAt)
	cpuStats.ReadAt = time.Now()

}

func getCPUUtilization(idle, total, lastIdle, lasstTotal uint64) float64 {
	idleTicks := float64(idle - lastIdle)
	totalTicks := float64(total - lasstTotal)

	return 100 * (totalTicks - idleTicks) / totalTicks
}

func calcDuration(lastReadTime time.Time) string {
	return time.Now().Sub(lastReadTime).String()
}
