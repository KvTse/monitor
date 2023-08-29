package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"monitor/config"
	"monitor/monitor"
	"monitor/service"
	"monitor/timer"
	"net/http"
	"time"
)

func main() {
	config.Init()
	if config.Conf.Enable {
		fmt.Println("start a task to report metrics info ")
		gt := timer.NewMyTick(config.Conf.ReportInterval, GetMonitorInfo)
		gt.Start()
	}
}
func GetMonitorInfo() error {
	go asyncReportMonitorInfo()

	return nil
}
func asyncReportMonitorInfo() {
	memInfo := service.GetMemInfo()
	cpuInfo := service.GetCpuInfo()
	processInfo := service.GetProcess()
	monitorInfo := monitor.Monitor{
		CurrentTime:  time.Now().UnixMilli(),
		MemoryInfo:   memInfo,
		CpuInfo:      cpuInfo,
		ProcessInfos: processInfo,
	}
	marshal, _ := json.Marshal(monitorInfo)
	req, _ := http.NewRequest("POST", config.Conf.ReportUrl, bytes.NewBuffer(marshal))
	req.Header.Set("Content-Type", "application/json")
	//timeout := config.Conf.HttpTimeout

	client := &http.Client{Timeout: 10 * time.Second}

	resp, _ := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		fmt.Println("response Body:", string(body))
	}
}
