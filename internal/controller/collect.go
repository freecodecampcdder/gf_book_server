package controller

import (
	"context"
	v1 "goFrameMyServer/api/v1/app"
	"goFrameMyServer/internal/service"
)

var Collect = cCollect{}

type cCollect struct {
}

func (a *cCollect) List(ctx context.Context, req *v1.CollectListReq) (res *v1.CollectListRes, err error) {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Size = 20
	}
	res, err = service.Collect().List(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *cCollect) Collect(ctx context.Context, req *v1.CollectClickReq) (res *v1.CollectClickRes, err error) {
	res, err = service.Collect().CollectClick(ctx, req)
	if err != nil {
		return nil, err
	}
	d := new(v1.CollectClickRes)
	d.Successful = true
	return d, nil
}

func (a *cCollect) Del(ctx context.Context, req *v1.CollectDelReq) (res *v1.CollectDelRes, err error) {
	res, err = service.Collect().Del(ctx, req)
	if err != nil {
		return nil, err
	}
	d := new(v1.CollectDelRes)
	d.Successful = true
	return d, nil
}
