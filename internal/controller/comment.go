package controller

import (
	"context"
	"goFrameMyServer/api/v1/admin"
	v1 "goFrameMyServer/api/v1/app"
	"goFrameMyServer/internal/service"
)

var Comment = cComment{}

type cComment struct {
}

func (a *cComment) Add(ctx context.Context, req *v1.CommentAddReq) (res *v1.CommentAddRes, err error) {
	res, err = service.Comment().Add(ctx, req)
	if err != nil {
		return nil, err
	}
	d := new(v1.CommentAddRes)
	d.Successful = true
	return d, nil
}

func (a *cComment) List(ctx context.Context, req *v1.CommentListReq) (res *v1.CommentListRes, err error) {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Size = 20
	}
	res, err = service.Comment().List(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}

func (a *cComment) Like(ctx context.Context, req *v1.CommentLikeReq) (res *v1.CommentLikeRes, err error) {
	res, err = service.Comment().Like(ctx, req)
	if err != nil {
		return nil, err
	}
	d := new(v1.CommentLikeRes)
	d.Successful = true
	return d, nil
}

func (a *cComment) AdminList(ctx context.Context, req *admin.CommentListReq) (res *admin.CommentListRes, err error) {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Size = 20
	}
	res, err = service.Comment().AdminList(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}

func (a *cComment) AdminDel(ctx context.Context, req *admin.CommentDelReq) (res *admin.CommentDelRes, err error) {
	res, err = service.Comment().AdminDel(ctx, req)
	if err != nil {
		return nil, err
	}
	d := new(admin.CommentDelRes)
	d.Successful = true
	return d, nil
}

func (a *cComment) MyList(ctx context.Context, req *v1.MyCommentListReq) (res *v1.MyCommentListRes, err error) {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Size = 20
	}
	res, err = service.Comment().MyList(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}
