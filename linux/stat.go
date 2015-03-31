package linux

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type (
	Stat struct {
		CPUStatAll      CPUStat   `json:"cpu_all"`
		CPUStats        []CPUStat `json:"cpus"`
		Interrupts      uint64    `json:"intr"`
		ContextSwitches uint64    `json:"ctxt"`
		BootTime        time.Time `json:"btime"`
		Processes       uint64    `json:"processes"`
		ProcsRunning    uint64    `json:"procs_running"`
		ProcsBlocked    uint64    `json:"procs_blocked"`
	}

	CPUStat struct {
		Id        string `json:"id"`
		User      uint64 `json:"user"`
		Nice      uint64 `json:"nice"`
		System    uint64 `json:"system"`
		Idle      uint64 `json:"idle"`
		IOWait    uint64 `json:"iowait"`
		IRQ       uint64 `json:"irq"`
		SoftIRQ   uint64 `json:"softirq"`
		Steal     uint64 `json:"steal"`
		Guest     uint64 `json:"guest"`
		GuestNice uint64 `json:"guest_nice"`
	}
)

func createCPUStat(fields []string) *CPUStat {

	for i := 0; i < 12-len(fields); i++ {
		fields = append(fields, "")
	}

	s := &CPUStat{}

	s.Id = fields[0]
	s.User, _ = strconv.ParseUint(fields[1], 10, 64)
	s.Nice, _ = strconv.ParseUint(fields[2], 10, 64)
	s.System, _ = strconv.ParseUint(fields[3], 10, 64)
	s.Idle, _ = strconv.ParseUint(fields[4], 10, 64)
	s.IOWait, _ = strconv.ParseUint(fields[5], 10, 64)
	s.IRQ, _ = strconv.ParseUint(fields[6], 10, 64)
	s.SoftIRQ, _ = strconv.ParseUint(fields[7], 10, 64)
	s.Steal, _ = strconv.ParseUint(fields[8], 10, 64)
	s.Guest, _ = strconv.ParseUint(fields[9], 10, 64)
	s.GuestNice, _ = strconv.ParseUint(fields[10], 10, 64)

	return s

}

func ReadStatFromBytes(content []byte) (*Stat, error) {

	stat := &Stat{}

	for s := bufio.NewScanner(bytes.NewReader(content)); s.Scan(); {

		fields := strings.Fields(s.Text())
		if len(fields) == 0 {
			continue
		}

		switch fields[0] {
		case "intr":
			stat.Interrupts, _ = strconv.ParseUint(fields[1], 10, 64)
		case "ctxt":
			stat.ContextSwitches, _ = strconv.ParseUint(fields[1], 10, 64)
		case "btime":
			seconds, _ := strconv.ParseInt(fields[1], 10, 64)
			stat.BootTime = time.Unix(seconds, 0)
		case "processes":
			stat.Processes, _ = strconv.ParseUint(fields[1], 10, 64)
		case "procs_running":
			stat.ProcsRunning, _ = strconv.ParseUint(fields[1], 10, 64)
		case "procs_blocked":
			stat.ProcsBlocked, _ = strconv.ParseUint(fields[1], 10, 64)
		case "cpu":
			if cpuStat := createCPUStat(fields); cpuStat != nil {
				stat.CPUStatAll = *cpuStat
			}
		default:
			if strings.HasPrefix(fields[0], "cpu") {
				if cpuStat := createCPUStat(fields); cpuStat != nil {
					stat.CPUStats = append(stat.CPUStats, *cpuStat)
				}
			}
		}

	}

	return stat, nil

}

func ReadStat(path string) (*Stat, error) {

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return ReadStatFromBytes(b)

}
