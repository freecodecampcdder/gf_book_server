package notify

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"goFrameMyServer/api/v1"
	"goFrameMyServer/api/v1/admin"
	v1 "goFrameMyServer/api/v1/app"
	"goFrameMyServer/internal/dao"
	"goFrameMyServer/internal/model"
	"goFrameMyServer/internal/model/entity"
	"goFrameMyServer/internal/service"
)

type sNotify struct {
}

func init() {
	service.RegisterNotify(New())
}

func New() *sNotify {
	return &sNotify{}
}

func (s *sNotify) AdminAdd(ctx context.Context, req *admin.NotifyAddReq) (res *admin.NotifyAddRes, err error) {

	_, err = dao.Notify.Ctx(ctx).Data(g.Map{
		"title":   req.Title,
		"content": req.Content,
	}).Insert()
	if err != nil {
		return nil, errors.New("插入失败")
	}
	return
}

func (s *sNotify) AdminUpd(ctx context.Context, req *admin.NotifyUpdReq) (res *admin.NotifyUpdRes, err error) {
	_, err = dao.Notify.Ctx(ctx).WherePri(req.Id).Data(g.Map{
		"title":   req.Title,
		"content": req.Content,
	}).Update()
	if err != nil {
		return nil, errors.New("修改失败")
	}
	return
}

func (s *sNotify) AdminDel(ctx context.Context, req *admin.NotifyDelReq) (res *admin.NotifyDelRes, err error) {
	_, err = dao.Notify.Ctx(ctx).WhereIn("id", req.Id).Delete()
	if err != nil {
		return nil, errors.New("删除")
	}
	return
}

func (s *sNotify) AdminList(ctx context.Context, req *admin.NotifyListReq) (res *admin.NotifyListRes, err error) {
	var list []*entity.Notify

	var m = dao.Notify.Ctx(ctx)

	res = &admin.NotifyListRes{
		CommonPaginationRes: api.CommonPaginationRes{
			Page: req.Page,
			Size: req.Size,
			List: []*entity.Notify{},
		},
	}
	listModel := m.Page(req.Page, req.Size)
	if res.Total, err = listModel.Count(); err != nil {
		return res, err
	}
	if res.Total == 0 {
		return res, nil
	}
	if err = listModel.Scan(&list); err != nil {
		return res, err
	}
	if len(list) == 0 {
		return res, err
	}
	res.List = list
	return
}

func (s *sNotify) List(ctx context.Context, req *v1.NotifyListReq) (res *v1.NotifyListRes, err error) {
	var list []*model.NotifyList
	err = dao.Notify.Ctx(ctx).Limit(6).Order("id desc").Scan(&list)
	if err != nil {
		return nil, errors.New("通知查询失败")
	}
	res = new(v1.NotifyListRes)
	res.List = list
	return
}
