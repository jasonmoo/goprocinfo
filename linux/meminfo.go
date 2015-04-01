package linux

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

type (
	MemInfo struct {
		MemTotal          uint64 `json:"mem_total"`
		MemFree           uint64 `json:"mem_free"`
		MemAvailable      uint64 `json:"mem_available"`
		Buffers           uint64 `json:"buffers"`
		Cached            uint64 `json:"cached"`
		SwapCached        uint64 `json:"swap_cached"`
		Active            uint64 `json:"active"`
		Inactive          uint64 `json:"inactive"`
		ActiveAnon        uint64 `json:"active_anon"`
		InactiveAnon      uint64 `json:"inactive_anon"`
		ActiveFile        uint64 `json:"active_file"`
		InactiveFile      uint64 `json:"inactive_file"`
		Unevictable       uint64 `json:"unevictable"`
		Mlocked           uint64 `json:"mlocked"`
		SwapTotal         uint64 `json:"swap_total"`
		SwapFree          uint64 `json:"swap_free"`
		Dirty             uint64 `json:"dirty"`
		Writeback         uint64 `json:"write_back"`
		AnonPages         uint64 `json:"anon_pages"`
		Mapped            uint64 `json:"mapped"`
		Shmem             uint64 `json:"shmem"`
		Slab              uint64 `json:"slab"`
		SReclaimable      uint64 `json:"s_reclaimable"`
		SUnreclaim        uint64 `json:"s_unclaim"`
		KernelStack       uint64 `json:"kernel_stack"`
		PageTables        uint64 `json:"page_tables"`
		NFS_Unstable      uint64 `json:"nfs_unstable"`
		Bounce            uint64 `json:"bounce"`
		WritebackTmp      uint64 `json:"writeback_tmp"`
		CommitLimit       uint64 `json:"commit_limit"`
		Committed_AS      uint64 `json:"committed_as"`
		VmallocTotal      uint64 `json:"vmalloc_total"`
		VmallocUsed       uint64 `json:"vmalloc_used"`
		VmallocChunk      uint64 `json:"vmalloc_chunk"`
		HardwareCorrupted uint64 `json:"hardware_corrupted"`
		AnonHugePages     uint64 `json:"anon_huge_pages"`
		HugePages_Total   uint64 `json:"huge_pages_total"`
		HugePages_Free    uint64 `json:"huge_pages_free"`
		HugePages_Rsvd    uint64 `json:"huge_pages_rsvd"`
		HugePages_Surp    uint64 `json:"huge_pages_surp"`
		Hugepagesize      uint64 `json:"hugepagesize"`
		DirectMap4k       uint64 `json:"direct_map_4k"`
		DirectMap2M       uint64 `json:"direct_map_2M"`
		DirectMap1G       uint64 `json:"direct_map_1G"`
	}
)

func ReadMemInfoFromBytes(data []byte) (*MemInfo, error) {

	info := &MemInfo{}
	v := reflect.ValueOf(info).Elem()
	IsSpaceOrColon := func(r rune) bool { return unicode.IsSpace(r) || r == ':' }

	for s := bufio.NewScanner(bytes.NewReader(data)); s.Scan(); {

		fields := strings.FieldsFunc(s.Text(), IsSpaceOrColon)
		if len(fields) < 2 {
			continue
		}

		name := fields[0]
		value, _ := strconv.ParseUint(fields[1], 10, 64)

		// store value in bytes
		if len(fields) == 3 && strings.ToLower(fields[2]) == "kb" {
			value *= 1024
		}

		switch name {
		case "Active(anon)":
			v.FieldByName("ActiveAnon").SetUint(value)
		case "Inactive(anon)":
			v.FieldByName("InactiveAnon").SetUint(value)
		case "Active(file)":
			v.FieldByName("ActiveFile").SetUint(value)
		case "Inactive(file)":
			v.FieldByName("InactiveFile").SetUint(value)
		default:
			if f := v.FieldByName(fields[0]); f.CanSet() {
				f.SetUint(value)
			}
		}

	}

	return info, nil

}

func ReadMemInfo(path string) (*MemInfo, error) {

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return ReadMemInfoFromBytes(data)

}
