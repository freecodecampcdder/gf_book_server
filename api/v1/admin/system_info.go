package admin

import "github.com/gogf/gf/v2/frame/g"

type SystemInfoReq struct {
	g.Meta `path:"/system/info" method:"get" tags:"后台-系统" summary:"后台-系统消息"`
}
type SystemInfoRes struct {
	Name              string  `json:"name"`
	V                 string  `json:"v"`
	System            string  `json:"system"`
	CpuNum            int     `json:"cpu_num"`
	Cpu               string  `json:"cpu"`
	CpuLoad           float64 `json:"cpu_load"`
	InternalMemory    string  `json:"internal_memory"`
	UseMemory         string  `json:"use_memory"`
	MemoryUtilization float64 `json:"memory_utilization"`
	Language          string  `json:"language"`
	GoV               string  `json:"go_v"`
	SystemTimeZone    string  `json:"system_time_zone"`
}
