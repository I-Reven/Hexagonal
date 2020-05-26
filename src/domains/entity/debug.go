package entity

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"runtime"
	"time"
)

type (
	Debug struct {
		Message   string   `cql:"message" json:"message"`
		Data      []string `cql:"data" json:"data"`
		Memory    string   `cql:"memory" json:"memory"`
		CPU       string   `cql:"cpu" json:"cpu"`
		Timestamp int64    `cql:"timestamp" json:"timestamp"`
	}

	CPU struct {
		CPUser float64 `json:"cPUser"`
		CPNice float64 `json:"cPNice"`
		CPSys  float64 `json:"cPSys"`
		CPIntr float64 `json:"cPIntr"`
		//CPIdle    float64 `json:"cPIdle"`
		//CPUStates float64 `json:"cPUStates"`
	}

	Memory struct {
		Alloc        uint64 `json:"alloc"`
		TotalAlloc   uint64 `json:"totalAlloc"`
		Sys          uint64 `json:"sys"`
		Lookups      uint64 `json:"lookups"`
		Mallocs      uint64 `json:"mallocs"`
		Frees        uint64 `json:"frees"`
		HeapAlloc    uint64 `json:"heapAlloc"`
		HeapSys      uint64 `json:"heapSys"`
		HeapIdle     uint64 `json:"heapIdle"`
		HeapInuse    uint64 `json:"heapInuse"`
		HeapReleased uint64 `json:"heapReleased"`
		HeapObjects  uint64 `json:"heapObjects"`
		StackInuse   uint64 `json:"stackInuse"`
		StackSys     uint64 `json:"stackSys"`
		MSpanInuse   uint64 `json:"mSpanInuse"`
		MSpanSys     uint64 `json:"mSpanSys"`
		MCacheInuse  uint64 `json:"mCacheInuse"`
		MCacheSys    uint64 `json:"mCacheSys"`
		BuckHashSys  uint64 `json:"buckHashSys"`
		GCSys        uint64 `json:"gCSys"`
		OtherSys     uint64 `json:"otherSys"`
		NextGC       uint64 `json:"nextGC"`
		LastGC       uint64 `json:"lastGC"`
		PauseTotalNs uint64 `json:"pauseTotalNs"`
		//PauseNs       [256]uint64 `json:"pauseNs"`
		//PauseEnd      [256]uint64 `json:"pauseEnd"`
		NumGC         uint32  `json:"numGC"`
		NumForcedGC   uint32  `json:"numForcedGC"`
		GCCPUFraction float64 `json:"gCCPUFraction"`
		EnableGC      bool    `json:"enableGC"`
		DebugGC       bool    `json:"debugGC"`
	}
)

func (d *Debug) GetMessage() string           { return d.Message }
func (d *Debug) SetMessage(message string)    { d.Message = message }
func (d *Debug) GetData() []string            { return d.Data }
func (d *Debug) SetData(data []string)        { d.Data = data }
func (d *Debug) AddData(data interface{})     { d.Data = append(d.Data, fmt.Sprintf("%v", data)) }
func (d *Debug) GetTimestamp() int64          { return d.Timestamp }
func (d *Debug) SetTimestamp(timestamp int64) { d.Timestamp = timestamp }

func CreateDebugger(message string, data ...interface{}) Debug {
	debugger := Debug{
		Message: message,
	}

	debugger.SetMemory()
	debugger.SetCPU()

	for _, info := range data {
		debugger.AddData(info)
	}

	return debugger
}

func (d *Debug) GetMemory() Memory {
	memory := Memory{}
	_ = json.Unmarshal([]byte(d.Memory), &memory)

	return memory
}

func (d *Debug) SetMemory() {
	memory := runtime.MemStats{}
	runtime.ReadMemStats(&memory)
	Memo := Memory{
		memory.Alloc,
		memory.TotalAlloc,
		memory.Sys,
		memory.Lookups,
		memory.Mallocs,
		memory.Frees,
		memory.HeapAlloc,
		memory.HeapSys,
		memory.HeapIdle,
		memory.HeapInuse,
		memory.HeapReleased,
		memory.HeapObjects,
		memory.StackInuse,
		memory.StackSys,
		memory.MSpanInuse,
		memory.MSpanSys,
		memory.MCacheInuse,
		memory.MCacheSys,
		memory.BuckHashSys,
		memory.GCSys,
		memory.OtherSys,
		memory.NextGC,
		memory.LastGC,
		memory.PauseTotalNs,
		//memory.PauseNs,
		//memory.PauseEnd,
		memory.NumGC,
		memory.NumForcedGC,
		memory.GCCPUFraction,
		memory.EnableGC,
		memory.DebugGC,
	}

	memo, err := json.Marshal(Memo)

	if err == nil {
		d.Memory = string(memo)
	}
}

func (d *Debug) GetUPU() CPU {
	cpu := CPU{}
	_ = json.Unmarshal([]byte(d.CPU), &cpu)

	return cpu
}

func (d *Debug) SetCPU() {
	percent, _ := cpu.Percent(time.Nanosecond, true)
	Cpu := CPU{
		percent[cpu.CPUser],
		percent[cpu.CPNice],
		percent[cpu.CPSys],
		percent[cpu.CPIntr],
		//percent[cpu.CPIdle],
		//percent[cpu.CPUStates],
	}
	cpu, err := json.Marshal(Cpu)

	if err == nil {
		d.CPU = string(cpu)
	}
}
