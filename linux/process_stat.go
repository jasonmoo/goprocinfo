package linux

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// Status information about the process.
type (
	ProcessStat struct {
		Pid                 uint64 `json:"pid"`
		Comm                string `json:"comm"`
		State               string `json:"state"`
		Ppid                int64  `json:"ppid"`
		Pgrp                int64  `json:"pgrp"`
		Session             int64  `json:"session"`
		TtyNr               int64  `json:"tty_nr"`
		Tpgid               int64  `json:"tpgid"`
		Flags               uint64 `json:"flags"`
		Minflt              uint64 `json:"minflt"`
		Cminflt             uint64 `json:"cminflt"`
		Majflt              uint64 `json:"majflt"`
		Cmajflt             uint64 `json:"cmajflt"`
		Utime               uint64 `json:"utime"`
		Stime               uint64 `json:"stime"`
		Cutime              int64  `json:"cutime"`
		Cstime              int64  `json:"cstime"`
		Priority            int64  `json:"priority"`
		Nice                int64  `json:"nice"`
		NumThreads          int64  `json:"num_threads"`
		Itrealvalue         int64  `json:"itrealvalue"`
		Starttime           uint64 `json:"starttime"`
		Vsize               uint64 `json:"vsize"`
		Rss                 int64  `json:"rss"`
		Rsslim              uint64 `json:"rsslim"`
		Startcode           uint64 `json:"startcode"`
		Endcode             uint64 `json:"endcode"`
		Startstack          uint64 `json:"startstack"`
		Kstkesp             uint64 `json:"kstkesp"`
		Kstkeip             uint64 `json:"kstkeip"`
		Signal              uint64 `json:"signal"`
		Blocked             uint64 `json:"blocked"`
		Sigignore           uint64 `json:"sigignore"`
		Sigcatch            uint64 `json:"sigcatch"`
		Wchan               uint64 `json:"wchan"`
		Nswap               uint64 `json:"nswap"`
		Cnswap              uint64 `json:"cnswap"`
		ExitSignal          int64  `json:"exit_signal"`
		Processor           int64  `json:"processor"`
		RtPriority          uint64 `json:"rt_priority"`
		Policy              uint64 `json:"policy"`
		DelayacctBlkioTicks uint64 `json:"delayacct_blkio_ticks"`
		GuestTime           uint64 `json:"guest_time"`
		CguestTime          int64  `json:"cguest_time"`
		StartData           uint64 `json:"start_data"`
		EndData             uint64 `json:"end_data"`
		StartBrk            uint64 `json:"start_brk"`
		ArgStart            uint64 `json:"arg_start"`
		ArgEnd              uint64 `json:"arg_end"`
		EnvStart            uint64 `json:"env_start"`
		EnvEnd              uint64 `json:"env_end"`
		ExitCode            int64  `json:"exit_code"`
	}
)

var processStatRegExp = regexp.MustCompile("^(\\d+)( \\(.*?\\) )(.*)$")

func ReadProcessStatFromBytes(data []byte) (*ProcessStat, error) {

	e := processStatRegExp.FindStringSubmatch(strings.TrimSpace(string(data)))

	if len(e) != 4 {
		return nil, fmt.Errorf("Expected 3 fields in regexp, got %d", len(e))
	}

	// Inject process Pid/Comm
	f := append([]string{e[1], strings.TrimSpace(e[2])}, strings.Fields(e[3])...)

	// Ensure total number of possible fields
	f = append(f, make([]string, 52-len(f))...)

	stat := &ProcessStat{}

	stat.Pid, _ = strconv.ParseUint(f[0], 10, 64)
	stat.Comm = f[1]
	stat.State = f[2]
	stat.Ppid, _ = strconv.ParseInt(f[3], 10, 64)
	stat.Pgrp, _ = strconv.ParseInt(f[4], 10, 64)
	stat.Session, _ = strconv.ParseInt(f[5], 10, 64)
	stat.TtyNr, _ = strconv.ParseInt(f[6], 10, 64)
	stat.Tpgid, _ = strconv.ParseInt(f[7], 10, 64)
	stat.Flags, _ = strconv.ParseUint(f[8], 10, 64)
	stat.Minflt, _ = strconv.ParseUint(f[9], 10, 64)
	stat.Cminflt, _ = strconv.ParseUint(f[10], 10, 64)
	stat.Majflt, _ = strconv.ParseUint(f[11], 10, 64)
	stat.Cmajflt, _ = strconv.ParseUint(f[12], 10, 64)
	stat.Utime, _ = strconv.ParseUint(f[13], 10, 64)
	stat.Stime, _ = strconv.ParseUint(f[14], 10, 64)
	stat.Cutime, _ = strconv.ParseInt(f[15], 10, 64)
	stat.Cstime, _ = strconv.ParseInt(f[16], 10, 64)
	stat.Priority, _ = strconv.ParseInt(f[17], 10, 64)
	stat.Nice, _ = strconv.ParseInt(f[18], 10, 64)
	stat.NumThreads, _ = strconv.ParseInt(f[19], 10, 64)
	stat.Itrealvalue, _ = strconv.ParseInt(f[20], 10, 64)
	stat.Starttime, _ = strconv.ParseUint(f[21], 10, 64)
	stat.Vsize, _ = strconv.ParseUint(f[22], 10, 64)
	stat.Rss, _ = strconv.ParseInt(f[23], 10, 64)
	stat.Rsslim, _ = strconv.ParseUint(f[24], 10, 64)
	stat.Startcode, _ = strconv.ParseUint(f[25], 10, 64)
	stat.Endcode, _ = strconv.ParseUint(f[26], 10, 64)
	stat.Startstack, _ = strconv.ParseUint(f[27], 10, 64)
	stat.Kstkesp, _ = strconv.ParseUint(f[28], 10, 64)
	stat.Kstkeip, _ = strconv.ParseUint(f[29], 10, 64)
	stat.Signal, _ = strconv.ParseUint(f[30], 10, 64)
	stat.Blocked, _ = strconv.ParseUint(f[31], 10, 64)
	stat.Sigignore, _ = strconv.ParseUint(f[32], 10, 64)
	stat.Sigcatch, _ = strconv.ParseUint(f[33], 10, 64)
	stat.Wchan, _ = strconv.ParseUint(f[34], 10, 64)
	stat.Nswap, _ = strconv.ParseUint(f[35], 10, 64)
	stat.Cnswap, _ = strconv.ParseUint(f[36], 10, 64)
	stat.ExitSignal, _ = strconv.ParseInt(f[37], 10, 64)
	stat.Processor, _ = strconv.ParseInt(f[38], 10, 64)
	stat.RtPriority, _ = strconv.ParseUint(f[39], 10, 64)
	stat.Policy, _ = strconv.ParseUint(f[40], 10, 64)
	stat.DelayacctBlkioTicks, _ = strconv.ParseUint(f[41], 10, 64)
	stat.GuestTime, _ = strconv.ParseUint(f[42], 10, 64)
	stat.CguestTime, _ = strconv.ParseInt(f[43], 10, 64)
	stat.StartData, _ = strconv.ParseUint(f[44], 10, 64)
	stat.EndData, _ = strconv.ParseUint(f[45], 10, 64)
	stat.StartBrk, _ = strconv.ParseUint(f[46], 10, 64)
	stat.ArgStart, _ = strconv.ParseUint(f[47], 10, 64)
	stat.ArgEnd, _ = strconv.ParseUint(f[48], 10, 64)
	stat.EnvStart, _ = strconv.ParseUint(f[49], 10, 64)
	stat.EnvEnd, _ = strconv.ParseUint(f[50], 10, 64)
	stat.ExitCode, _ = strconv.ParseInt(f[51], 10, 64)

	return stat, nil

}

func ReadProcessStat(path string) (*ProcessStat, error) {

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return ReadProcessStatFromBytes(data)

}
