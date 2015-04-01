package linux

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type Uptime struct {
	Total time.Duration `json:"total"`
	Idle  time.Duration `json:"idle"`
}

func (self *Uptime) GetTotalDuration() time.Duration {
	return self.Total
}

func (self *Uptime) GetIdleDuration() time.Duration {
	return self.Idle
}

func (self *Uptime) CalculateIdle() float64 {
	// XXX
	// num2/(num1*N)     # N = SMP CPU numbers
	return 0
}

func ReadUptimeFromBytes(data []byte) (*Uptime, error) {

	fields := strings.Fields(string(data))

	if len(fields) != 2 {
		return nil, fmt.Errorf("Expected 2 fields, got %d", len(fields))
	}

	total, _ := strconv.ParseFloat(fields[0], 64)
	idle, _ := strconv.ParseFloat(fields[1], 64)

	return &Uptime{
		Total: time.Duration(total) * time.Second,
		Idle:  time.Duration(idle) * time.Second,
	}, nil

}

func ReadUptime(path string) (*Uptime, error) {

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return ReadUptimeFromBytes(data)

}
