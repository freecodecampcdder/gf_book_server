package cron

import (
	"context"
	"github.com/duke-git/lancet/v2/datetime"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"goFrameMyServer/internal/dao"
	"time"
)

func InitCron() {
	_, _ = gcron.Add(context.TODO(), "1 * * * * ?", func(ctx context.Context) {
		datetime.BeginOfDay(datetime.AddDay(time.Now(), -1))
		count, _ := dao.LoginLog.Ctx(ctx).Where("created_at >= ? and created_at <= ?", datetime.BeginOfDay(datetime.AddDay(time.Now(), -1)), datetime.EndOfDay(datetime.AddDay(time.Now(), -1))).Group("ip").Count()
		_, _ = dao.Interview.Ctx(ctx).Insert(g.Map{
			"t":               datetime.BeginOfDay(datetime.AddDay(time.Now(), -1)),
			"interview_count": count,
		})
	})
}
