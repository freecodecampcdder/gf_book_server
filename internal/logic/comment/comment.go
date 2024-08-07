package comment

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/util/gconv"
	"goFrameMyServer/api/v1"
	"goFrameMyServer/api/v1/admin"
	v1 "goFrameMyServer/api/v1/app"
	"goFrameMyServer/internal/consts"
	"goFrameMyServer/internal/dao"
	"goFrameMyServer/internal/model"
	"goFrameMyServer/internal/model/entity"
	"goFrameMyServer/internal/service"
)

type sComment struct {
}

func init() {
	service.RegisterComment(New())
}

func New() *sComment {
	return &sComment{}
}

func (a *sComment) Add(ctx context.Context, req *v1.CommentAddReq) (res *v1.CommentAddRes, err error) {
	userId := gconv.Int(ctx.Value(consts.CtxUserId))
	comment := &entity.Comment{
		UserId:  userId,
		BookId:  req.BookId,
		Context: req.Context,
	}
	_, err = dao.Comment.Ctx(ctx).Data(comment).InsertAndGetId()
	if err != nil {
		return nil, errors.New("评论失败")
	}
	return
}

func (a *sComment) List(ctx context.Context, req *v1.CommentListReq) (res *v1.CommentListRes, err error) {
	var list []*model.CommentList
	sql := dao.Comment.Ctx(ctx).
		As("c").
		LeftJoin("user as u", "c.user_id = u.id").
		Where("c.book_id = ?", req.BookId)
	res = &v1.CommentListRes{
		CommonPaginationRes: api.CommonPaginationRes{
			Page: req.Page,
			Size: req.Size,
			List: []*model.CommentList{},
		},
	}
	if res.Total, err = sql.Count(); err != nil {
		return res, err
	}
	if res.Total == 0 {
		return res, nil
	}
	if req.Order == 2 {
		sql = sql.Order("c.like desc")
	} else {
		sql = sql.Order("c.id desc")
	}
	err = sql.Fields("c.id as id,c.user_id as user_id,u.avatar as avatar,u.nickname as nickname,c.context as context,c.created_at as created_at,c.like").
		Page(req.Page, req.Size).
		Scan(&list)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	res.List = list
	return
}

func (a *sComment) Like(ctx context.Context, req *v1.CommentLikeReq) (res *v1.CommentLikeRes, err error) {
	_, err = dao.Comment.Ctx(ctx).WherePri(req.Id).Increment("like", 1)
	if err != nil {
		return nil, errors.New("点赞失败")
	}
	return
}

func (a *sComment) AdminList(ctx context.Context, req *admin.CommentListReq) (res *admin.CommentListRes, err error) {
	var list []*model.AdminCommentList
	sql := dao.Comment.Ctx(ctx).
		As("c").
		LeftJoin("user as u", "c.user_id = u.id").
		LeftJoin("book as b", "c.book_id = b.id")
	res = &admin.CommentListRes{
		CommonPaginationRes: api.CommonPaginationRes{
			Page: req.Page,
			Size: req.Size,
			List: []*model.AdminCommentList{},
		},
	}
	if res.Total, err = sql.Count(); err != nil {
		return res, err
	}
	if res.Total == 0 {
		return res, nil
	}
	err = sql.Fields("c.id as id,c.user_id as user_id,u.nickname as nickname,c.context as context,c.created_at as created_at,b.id as book_id,b.title as book_title").
		Page(req.Page, req.Size).
		Scan(&list)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	res.List = list
	return
}
func (a *sComment) AdminDel(ctx context.Context, req *admin.CommentDelReq) (res *admin.CommentDelRes, err error) {
	_, err = dao.Comment.Ctx(ctx).WhereIn("id", req.Id).Delete()
	if err != nil {
		return nil, errors.New("删除失败")
	}
	return
}

func (a *sComment) MyList(ctx context.Context, req *v1.MyCommentListReq) (res *v1.MyCommentListRes, err error) {
	var list []*model.MyCommentList
	sql := dao.Comment.Ctx(ctx).
		As("c").
		LeftJoin("user as u", "c.user_id = u.id").
		LeftJoin("book as b", "c.book_id = b.id")
	res = &v1.MyCommentListRes{
		CommonPaginationRes: api.CommonPaginationRes{
			Page: req.Page,
			Size: req.Size,
			List: []*model.MyCommentList{},
		},
	}
	if res.Total, err = sql.Count(); err != nil {
		return res, err
	}
	if res.Total == 0 {
		return res, nil
	}
	err = sql.Fields("c.id as id,c.user_id as user_id,u.avatar as avatar,u.nickname as nickname,c.context as context,c.created_at as created_at,c.like,b.title").
		Page(req.Page, req.Size).
		Scan(&list)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	res.List = list
	return
}
