package sort

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"goFrameMyServer/api/v1/admin"
	v1 "goFrameMyServer/api/v1/app"
	"goFrameMyServer/internal/dao"
	"goFrameMyServer/internal/model"
	"goFrameMyServer/internal/service"
)

type sSort struct {
}

func init() {
	service.RegisterSort(New())
}

func New() *sSort {
	return &sSort{}
}

func (s *sSort) List(ctx context.Context) (res *v1.SortListRes, err error) {
	res = new(v1.SortListRes)
	err = dao.Sort.Ctx(ctx).Scan(&res.List)
	if err != nil {
		return nil, errors.New("查询分类失败")
	}
	return
}

func (s *sSort) AdminList(ctx context.Context, req *admin.SortListReq) (res *admin.SortListRes, err error) {
	var sort []*model.AdminSortList
	var m = dao.Sort.Ctx(ctx)
	res = &admin.SortListRes{
		List: []*model.AdminTagList{},
	}
	if err = m.Where("id !=?", -1).Scan(&sort); err != nil {
		return res, err
	}
	if len(sort) == 0 {
		return res, err
	}
	res.List = sort
	return
}

func (s *sSort) AdminAdd(ctx context.Context, req *admin.SortAddReq) (res *admin.SortAddRes, err error) {
	count, err := dao.Sort.Ctx(ctx).Where("title = ?", req.Title).Count()
	if err != nil {
		return nil, errors.New("查询标签错误")
	}
	if count > 0 {
		return nil, errors.New("重复添加")
	}
	_, err = dao.Sort.Ctx(ctx).Insert(g.Map{"title": req.Title})
	if err != nil {
		return nil, errors.New("添加错误")
	}
	return
}

func (s *sSort) AdminUpd(ctx context.Context, req *admin.SortUpdReq) (res *admin.SortUpdRes, err error) {
	count, err := dao.Sort.Ctx(ctx).Where("title = ?", req.Title).Count()
	if err != nil {
		return nil, errors.New("查询标签错误")
	}
	if count > 0 {
		return nil, errors.New("重复添加")
	}
	_, err = dao.Sort.Ctx(ctx).WherePri(req.Id).Update(g.Map{"title": req.Title})
	if err != nil {
		return nil, errors.New("修改错误")
	}
	return
}

func (s *sSort) AdminDel(ctx context.Context, req *admin.SortDelReq) (res *admin.SortDelRes, err error) {
	_, err = dao.Sort.Ctx(ctx).WhereIn("id", req.Id).Delete()
	if err != nil {
		return nil, errors.New("删除失败")
	}
	return
}
