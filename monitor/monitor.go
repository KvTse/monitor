package monitor

type Monitor struct {
	CurrentTime  int64         `json:"currentTime"`
	MemoryInfo   MemoryInfo    `json:"memoryInfo"`
	CpuInfo      CpuInfo       `json:"cpuInfo"`
	ProcessInfos []ProcessInfo `json:"processInfos"`
}

type CpuInfo struct {
	PhysicalCnt  int       `json:"physicalCnt"`
	LogicalCnt   int       `json:"logicalCnt"`
	TotalPercent float64   `json:"totalPercent"`
	PerPercents  []float64 `json:"perPercents"`
}
type MemoryInfo struct {
	Total uint64 `json:"total"`

	Available uint64 `json:"available"`

	Used uint64 `json:"used"`

	UsedPercent float64 `json:"usedPercent"`

	Free uint64 `json:"free"`

	Active uint64 `json:"active"`

	Inactive uint64 `json:"inactive"`
}
type ProcessInfo struct {
	ProcessName string            `json:"processName"`
	Pid         int32             `json:"pid"`
	Status      string            `json:"status"`
	CpuPercent  float64           `json:"cpuPercent"`
	MemoryInfo  ProcessMemoryInfo `json:"memoryInfo"`
	CreateTime  int64             `json:"createTime"`
	IsRunning   bool              `json:"isRunning"`
	RunningTime int64             `json:"runningTime"`
}
type ProcessMemoryInfo struct {
	Rss uint64 `json:"rss"`
	Vms uint64 `json:"vms"`
}
