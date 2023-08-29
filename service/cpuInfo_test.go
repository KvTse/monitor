package service

import (
	"fmt"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"testing"
)

func TestGetCpuInfo(t *testing.T) {
	info := GetCpuInfo()
	fmt.Printf("%v\n", info)
}
func TestGetCpuLoad(t *testing.T) {
	info, _ := load.Avg()
	misc, _ := load.Misc()
	fmt.Printf("%v\n %v", info, misc)
}

// host info
func TestGetHostInfo(t *testing.T) {
	hInfo, _ := host.Info()
	fmt.Printf("host info:%v uptime:%v boottime:%v\n", hInfo, hInfo.Uptime, hInfo.BootTime)
}

// mem info
func TestGetMemInfo(t *testing.T) {
	memInfo, _ := mem.VirtualMemory()
	fmt.Printf("mem info:%v\n", memInfo)
}
