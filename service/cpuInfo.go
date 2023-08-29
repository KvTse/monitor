package service

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"monitor/monitor"
	"time"
)

func GetCpuInfo() monitor.CpuInfo {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err:%v", err)
	}
	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}
	physicalCnt, _ := cpu.Counts(false)
	logicalCnt, _ := cpu.Counts(true)
	totalPercents, _ := cpu.Percent(time.Second, false)
	var totalPercent float64 = 0
	if len(totalPercents) >= 1 {
		totalPercent = totalPercents[0]
	}
	perPercents, _ := cpu.Percent(time.Second, true)
	return monitor.CpuInfo{
		PhysicalCnt:  physicalCnt,
		LogicalCnt:   logicalCnt,
		TotalPercent: totalPercent,
		PerPercents:  perPercents}
}
func GetCpuLoad() {
	info, _ := load.Avg()
	fmt.Printf("%v\n", info)
}
