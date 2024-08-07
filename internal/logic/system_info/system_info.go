package system_info

import (
	"context"
	"github.com/duke-git/lancet/v2/mathutil"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"goFrameMyServer/api/v1/admin"
	"goFrameMyServer/internal/consts"
	"goFrameMyServer/internal/service"
	"runtime"
	"time"
)

type sSystemInfo struct {
}

func init() {
	service.RegisterSystemInfo(New())
}

func New() *sSystemInfo {
	return &sSystemInfo{}
}

func (s *sSystemInfo) Info(ctx context.Context) (res *admin.SystemInfoRes, err error) {
	res = new(admin.SystemInfoRes)
	cpuInfo, _ := cpu.Info()
	percent, _ := cpu.Percent(time.Second, false)
	memInfo, _ := mem.VirtualMemory()
	res.Name = consts.BackendServerName
	res.V = consts.BackendServerV
	res.System = runtime.GOOS + " " + runtime.GOARCH
	res.CpuNum = runtime.NumCPU()
	res.Cpu = cpuInfo[0].ModelName
	res.CpuLoad = mathutil.RoundToFloat(percent[0], 2)
	res.InternalMemory = gconv.String(memInfo.Total/1024/1024/1024) + "G"
	res.UseMemory = gconv.String(memInfo.Total/1024/1024/1024-memInfo.Available/1024/1024/1024) + "G"
	res.MemoryUtilization = memInfo.UsedPercent
	res.Language = "zh_CN"
	res.GoV = runtime.Version()
	res.SystemTimeZone = "中国标准时间"

	return
}
