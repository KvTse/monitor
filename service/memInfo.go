package service

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
	"monitor/monitor"
)

func GetMemInfo() monitor.MemoryInfo {
	memInfo, err := mem.VirtualMemory()

	if err != nil {
		fmt.Printf("get mem info error :%v %v", memInfo, err)
	}

	return monitor.MemoryInfo{
		Total:       memInfo.Total,
		Available:   memInfo.Available,
		Used:        memInfo.Used,
		UsedPercent: memInfo.UsedPercent,
		Free:        memInfo.Free,
		Active:      memInfo.Active,
		Inactive:    memInfo.Inactive,
	}
}
