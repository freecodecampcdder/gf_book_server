package tag

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

type sTag struct {
}

func init() {
	service.RegisterTag(New())
}

func New() *sTag {
	return &sTag{}
}

func (s *sTag) List(ctx context.Context) (res *v1.TagListRes, err error) {
	res = new(v1.TagListRes)
	err = dao.Tag.Ctx(ctx).Scan(&res.List)
	if err != nil {
		return nil, errors.New("查询标签失败")
	}
	return
}

func (s *sTag) AdminList(ctx context.Context, req *admin.TagListReq) (res *admin.TagListRes, err error) {
	var tag []*model.AdminTagList
	var m = dao.Tag.Ctx(ctx)
	res = &admin.TagListRes{
		List: []*model.AdminTagList{},
	}
	if err = m.Scan(&tag); err != nil {
		return res, err
	}
	if len(tag) == 0 {
		return res, err
	}
	res.List = tag
	return
}

func (s *sTag) AdminAdd(ctx context.Context, req *admin.TagAddReq) (res *admin.TagAddRes, err error) {
	count, err := dao.Tag.Ctx(ctx).Where("title = ?", req.Title).Count()
	if err != nil {
		return nil, errors.New("查询标签错误")
	}
	if count > 0 {
		return nil, errors.New("重复添加")
	}
	_, err = dao.Tag.Ctx(ctx).Insert(g.Map{"title": req.Title})
	if err != nil {
		return nil, errors.New("添加错误")
	}
	return
}

func (s *sTag) AdminUpd(ctx context.Context, req *admin.TagUpdReq) (res *admin.TagUpdRes, err error) {
	count, err := dao.Tag.Ctx(ctx).Where("title = ?", req.Title).Count()
	if err != nil {
		return nil, errors.New("查询标签错误")
	}
	if count > 0 {
		return nil, errors.New("重复添加")
	}
	_, err = dao.Tag.Ctx(ctx).WherePri(req.Id).Update(g.Map{"title": req.Title})
	if err != nil {
		return nil, errors.New("修改错误")
	}
	return
}

func (s *sTag) AdminDel(ctx context.Context, req *admin.TagDelReq) (res *admin.TagDelRes, err error) {
	_, err = dao.Tag.Ctx(ctx).WhereIn("id", req.Id).Delete()
	if err != nil {
		return nil, errors.New("删除失败")
	}
	return
}
