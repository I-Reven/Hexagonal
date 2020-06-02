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
		Message   string   `cql:"message" json:"message" faker:"sentence"`
		Data      []string `cql:"data" json:"data"`
		Memory    string   `cql:"memory" json:"memory" faker:"sentence"`
		CPU       string   `cql:"cpu" json:"cpu" faker:"sentence"`
		Timestamp int64    `cql:"timestamp" json:"timestamp"`
	}

	CPU struct {
		CPUser float64 `json:"cPUser"`
		CPNice float64 `json:"cPNice"`
		CPSys  float64 `json:"cPSys"`
		CPIntr float64 `json:"cPIntr"`
	}

	Memory struct {
		Alloc         uint64  `json:"alloc"`
		TotalAlloc    uint64  `json:"totalAlloc"`
		Sys           uint64  `json:"sys"`
		Lookups       uint64  `json:"lookups"`
		Mallocs       uint64  `json:"mallocs"`
		Frees         uint64  `json:"frees"`
		HeapAlloc     uint64  `json:"heapAlloc"`
		HeapSys       uint64  `json:"heapSys"`
		HeapIdle      uint64  `json:"heapIdle"`
		HeapInuse     uint64  `json:"heapInuse"`
		HeapReleased  uint64  `json:"heapReleased"`
		HeapObjects   uint64  `json:"heapObjects"`
		StackInuse    uint64  `json:"stackInuse"`
		StackSys      uint64  `json:"stackSys"`
		MSpanInuse    uint64  `json:"mSpanInuse"`
		MSpanSys      uint64  `json:"mSpanSys"`
		MCacheInuse   uint64  `json:"mCacheInuse"`
		MCacheSys     uint64  `json:"mCacheSys"`
		BuckHashSys   uint64  `json:"buckHashSys"`
		GCSys         uint64  `json:"gCSys"`
		OtherSys      uint64  `json:"otherSys"`
		NextGC        uint64  `json:"nextGC"`
		LastGC        uint64  `json:"lastGC"`
		PauseTotalNs  uint64  `json:"pauseTotalNs"`
		NumGC         uint32  `json:"numGC"`
		NumForcedGC   uint32  `json:"numForcedGC"`
		GCCPUFraction float64 `json:"gCCPUFraction"`
		EnableGC      bool    `json:"enableGC"`
		DebugGC       bool    `json:"debugGC"`
	}
)

func (e *Debug) GetMessage() string           { return e.Message }
func (e *Debug) SetMessage(message string)    { e.Message = message }
func (e *Debug) GetData() []string            { return e.Data }
func (e *Debug) SetData(data []string)        { e.Data = data }
func (e *Debug) AddData(data interface{})     { e.Data = append(e.Data, fmt.Sprintf("%v", data)) }
func (e *Debug) GetTimestamp() int64          { return e.Timestamp }
func (e *Debug) SetTimestamp(timestamp int64) { e.Timestamp = timestamp }

func CreateDebugger(message string, data ...interface{}) Debug {
	debugger := Debug{
		Message: message,
	}

	debugger.SetMemory()
	//debugger.SetCPU()
	debugger.SetTimestamp(time.Now().UnixNano() / int64(time.Millisecond))

	for _, info := range data {
		debugger.AddData(info)
	}

	return debugger
}

func (e *Debug) GetMemory() Memory {
	memory := Memory{}
	_ = json.Unmarshal([]byte(e.Memory), &memory)

	return memory
}

func (e *Debug) SetMemory() {
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
		memory.NumGC,
		memory.NumForcedGC,
		memory.GCCPUFraction,
		memory.EnableGC,
		memory.DebugGC,
	}

	memo, err := json.Marshal(Memo)

	if err == nil {
		e.Memory = string(memo)
	}
}

func (e *Debug) GetUPU() CPU {
	cpu := CPU{}
	_ = json.Unmarshal([]byte(e.CPU), &cpu)

	return cpu
}

func (e *Debug) SetCPU() {
	percent, _ := cpu.Percent(time.Nanosecond, true)
	Cpu := CPU{
		percent[0],
		percent[1],
		percent[2],
		percent[3],
	}
	cpu, err := json.Marshal(Cpu)

	if err == nil {
		e.CPU = string(cpu)
	}
}
