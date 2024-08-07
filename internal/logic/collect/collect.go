package collect

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"goFrameMyServer/api/v1"
	v1 "goFrameMyServer/api/v1/app"
	"goFrameMyServer/internal/consts"
	"goFrameMyServer/internal/dao"
	"goFrameMyServer/internal/model"
	"goFrameMyServer/internal/service"
)

type sCollect struct {
}

func init() {
	service.RegisterCollect(New())
}

func New() *sCollect {
	return &sCollect{}
}
func (s *sCollect) CollectClick(ctx context.Context, req *v1.CollectClickReq) (res *v1.CollectClickRes, err error) {
	userId := gconv.Uint(ctx.Value(consts.CtxUserId))
	sql := dao.Collect.Ctx(ctx).Where("user_id = ? and book_id = ?", userId, req.BookId)
	bookCount, err := dao.Book.Ctx(ctx).WherePri(req.BookId).Count()
	if err != nil {
		return nil, errors.New("查询失败")
	}
	if bookCount == 0 {
		return nil, errors.New("这本书不存在")
	}
	count, err := sql.Count()
	if err != nil {
		return nil, errors.New("查询失败")
	}
	if count > 0 {
		_, err = sql.Delete()
		if err != nil {
			return nil, errors.New("删除错误")
		}
		_, err = dao.Book.Ctx(ctx).WherePri(req.BookId).Decrement("collect_num", 1)
		if err != nil {
			return nil, errors.New("减少错误")
		}

	} else {
		_, err = dao.Collect.Ctx(ctx).OmitEmpty().Insert(g.Map{
			"book_id": req.BookId,
			"user_id": userId,
		})
		if err != nil {
			return nil, errors.New("插入失败")
		}
		_, err = dao.Book.Ctx(ctx).WherePri(req.BookId).Increment("collect_num", 1)
		if err != nil {
			return nil, errors.New("增加错误")
		}
	}
	return
}

func (s *sCollect) List(ctx context.Context, req *v1.CollectListReq) (res *v1.CollectListRes, err error) {
	var list []*model.CollectList
	userId := gconv.Uint(ctx.Value(consts.CtxUserId))
	sql := dao.Collect.Ctx(ctx).As("c").LeftJoin("book as b", "c.book_id = b.id").Where("c.user_id = ?", userId)
	res = &v1.CollectListRes{
		CommonPaginationRes: api.CommonPaginationRes{
			Page: req.Page,
			Size: req.Size,
			List: []*model.CollectList{},
		},
	}
	sql = sql.Page(req.Page, req.Size)
	if res.Total, err = sql.Count(); err != nil {
		return res, err
	}
	if res.Total == 0 {
		return res, nil
	}
	sql = sql.Fields("c.id,c.book_id,b.title,b.cover,b.author,b.translator")
	if err = sql.Scan(&list); err != nil {
		return res, err
	}
	if len(list) == 0 {
		return res, err
	}
	res.List = list
	return
}

func (s *sCollect) Del(ctx context.Context, req *v1.CollectDelReq) (res *v1.CollectDelRes, err error) {
	_, err = dao.Collect.Ctx(ctx).WherePri(req.Id).Delete()
	if err != nil {
		return nil, errors.New("删除错误")
	}
	return
}
