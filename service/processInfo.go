package service

import (
	"github.com/shirou/gopsutil/process"
	"monitor/monitor"
	"time"
)

func GetProcess() []monitor.ProcessInfo {

	processes, _ := process.Processes()
	var processInfos []monitor.ProcessInfo
	for _, p := range processes {

		processName, _ := p.Name()
		pid := p.Pid
		status, _ := p.Status()
		cpuPercent, _ := p.CPUPercent()
		memoryInfo, _ := p.MemoryInfo()
		createTime, _ := p.CreateTime()
		isRunning, _ := p.IsRunning()
		runningTime := time.Now().UnixMilli() - createTime
		if createTime == 0 {
			runningTime = 0
		}
		var processMemoryInfo monitor.ProcessMemoryInfo
		if memoryInfo != nil {
			processMemoryInfo = monitor.ProcessMemoryInfo{
				Rss: memoryInfo.RSS, Vms: memoryInfo.VMS}
		}
		processInfo := monitor.ProcessInfo{ProcessName: processName, Pid: pid, Status: status, CpuPercent: cpuPercent, MemoryInfo: processMemoryInfo, CreateTime: createTime, IsRunning: isRunning, RunningTime: runningTime}
		processInfos = append(processInfos, processInfo)
	}
	return processInfos
}
