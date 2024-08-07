package log

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"goFrameMyServer/api/v1"
	"goFrameMyServer/api/v1/admin"
	"goFrameMyServer/internal/dao"
	"goFrameMyServer/internal/model"
	"goFrameMyServer/internal/model/entity"
	"goFrameMyServer/internal/service"
)

type sLog struct {
}

func init() {
	service.RegisterLog(New())
}

func New() *sLog {
	return &sLog{}
}

func (s *sLog) AdminLoginList(ctx context.Context, req *admin.LoginLogListReq) (res *admin.LoginLogListRes, err error) {
	var list []*model.AdminLoginLogList

	var m = dao.LoginLog.Ctx(ctx).As("l").LeftJoin("user as u", "l.user_id = u.id")

	res = &admin.LoginLogListRes{
		CommonPaginationRes: api.CommonPaginationRes{
			Page: req.Page,
			Size: req.Size,
			List: []*model.AdminLoginLogList{},
		},
	}
	listModel := m.Page(req.Page, req.Size)
	if res.Total, err = listModel.Count(); err != nil {
		return res, err
	}
	if res.Total == 0 {
		return res, nil
	}
	if err = listModel.Order("l.id desc").Fields("l.id,u.id as user_id,l.ip,u.nickname,l.user_agent,l.created_at").Scan(&list); err != nil {
		return res, err
	}
	if len(list) == 0 {
		return res, err
	}
	res.List = list
	return
}

func (s *sLog) AdminOpLogAdd(ctx context.Context, way, url, ip string, t int64) (err error) {
	_, err = dao.OpLog.Ctx(ctx).Data(g.Map{"way": way, "url": url, "ip": ip, "t": t}).Insert()
	if err != nil {
		return errors.New("插入错误")
	}
	return
}

func (s *sLog) AdminOpLogList(ctx context.Context, req *admin.OpLogListReq) (res *admin.OpLogListRes, err error) {
	var list []*entity.OpLog

	var m = dao.OpLog.Ctx(ctx)

	res = &admin.OpLogListRes{
		CommonPaginationRes: api.CommonPaginationRes{
			Page: req.Page,
			Size: req.Size,
			List: []*entity.OpLog{},
		},
	}
	listModel := m.Page(req.Page, req.Size)
	if res.Total, err = listModel.Count(); err != nil {
		return res, err
	}
	if res.Total == 0 {
		return res, nil
	}
	if err = listModel.Order("id desc").Scan(&list); err != nil {
		return res, err
	}
	if len(list) == 0 {
		return res, err
	}
	res.List = list
	return
}
