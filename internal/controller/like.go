package controller

import (
	"context"
	v1 "goFrameMyServer/api/v1/app"
	"goFrameMyServer/internal/service"
)

var Like = cLike{}

type cLike struct {
}

func (a *cLike) Like(ctx context.Context, req *v1.LikeClickReq) (res *v1.LikeClickRes, err error) {
	res, err = service.Like().LikeClick(ctx, req)
	if err != nil {
		return nil, err
	}
	d := new(v1.LikeClickRes)
	d.Successful = true
	return d, nil
}

func (a *cLike) List(ctx context.Context, req *v1.LikeListReq) (res *v1.LikeListRes, err error) {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Size = 20
	}
	res, err = service.Like().List(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *cLike) Del(ctx context.Context, req *v1.LikeDelReq) (res *v1.LikeDelRes, err error) {
	res, err = service.Like().Del(ctx, req)
	if err != nil {
		return nil, err
	}
	d := new(v1.LikeDelRes)
	d.Successful = true
	return d, nil
}
