package controller

import (
	"context"
	"goFrameMyServer/api/v1/admin"
	v1 "goFrameMyServer/api/v1/app"
	"goFrameMyServer/internal/service"
)

var Notify = cNotify{}

type cNotify struct {
}

func (a *cNotify) AdminAdd(ctx context.Context, req *admin.NotifyAddReq) (res *admin.NotifyAddRes, err error) {
	res, err = service.Notify().AdminAdd(ctx, req)
	if err != nil {
		return nil, err
	}
	d := new(admin.NotifyAddRes)
	d.Successful = true
	return d, nil
}

func (a *cNotify) AdminUpd(ctx context.Context, req *admin.NotifyUpdReq) (res *admin.NotifyUpdRes, err error) {
	res, err = service.Notify().AdminUpd(ctx, req)
	if err != nil {
		return nil, err
	}
	d := new(admin.NotifyUpdRes)
	d.Successful = true
	return d, nil
}

func (a *cNotify) AdminDel(ctx context.Context, req *admin.NotifyDelReq) (res *admin.NotifyDelRes, err error) {
	res, err = service.Notify().AdminDel(ctx, req)
	if err != nil {
		return nil, err
	}
	d := new(admin.NotifyDelRes)
	d.Successful = true
	return d, nil
}

func (a *cNotify) AdminList(ctx context.Context, req *admin.NotifyListReq) (res *admin.NotifyListRes, err error) {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Size = 20
	}
	res, err = service.Notify().AdminList(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *cNotify) List(ctx context.Context, req *v1.NotifyListReq) (res *v1.NotifyListRes, err error) {
	res, err = service.Notify().List(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
