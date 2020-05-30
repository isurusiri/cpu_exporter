package cpuclient

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type CPUStat struct {
	Idle                  uint64
	Total                 uint64
	LastIdle              uint64
	LastTotal             uint64
	Utilization           uint64
	ReadAt                time.Time
	DureationFromLastRead time.Duration
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
