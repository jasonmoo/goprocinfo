package linux

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

type LoadAvg struct {
	Last1Min       float64 `json:"last1min"`
	Last5Min       float64 `json:"last5min"`
	Last15Min      float64 `json:"last15min"`
	ProcessRunning uint64  `json:"process_running"`
	ProcessTotal   uint64  `json:"process_total"`
	LastPID        uint64  `json:"last_pid"`
}

func ReadLoadAvgFromBytes(data []byte) (*LoadAvg, error) {

	loadavg := LoadAvg{}

	if _, err := fmt.Fscan(
		bytes.NewReader(data),
		&loadavg.Last1Min,
		&loadavg.Last5Min,
		&loadavg.Last15Min,
		&loadavg.ProcessRunning,
		&loadavg.ProcessTotal,
		&loadavg.LastPID,
	); err != nil {
		return nil, err
	}

	return &loadavg, nil
}

func ReadLoadAvg(path string) (*LoadAvg, error) {

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return ReadLoadAvgFromBytes(data)

}
