package controller

import (
	"context"
	"goFrameMyServer/api/v1/admin"
	v1 "goFrameMyServer/api/v1/app"
	"goFrameMyServer/internal/service"
)

var Order = cOrder{}

type cOrder struct {
}

func (a *cOrder) Preload(ctx context.Context, req *v1.OrderPreloadReq) (res *v1.OrderPreloadRes, err error) {
	res, err = service.Order().Preload(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}

func (a *cOrder) Add(ctx context.Context, req *v1.OrderAddReq) (res *v1.OrderAddRes, err error) {
	res, err = service.Order().Add(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}

func (a *cOrder) Details(ctx context.Context, req *v1.OrderDetailsReq) (res *v1.OrderDetailsRes, err error) {
	res, err = service.Order().Details(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}

func (a *cOrder) Pay(ctx context.Context, req *v1.OrderPayReq) (res *v1.OrderPayRes, err error) {
	res, err = service.Order().Pay(ctx, req)
	if err != nil {
		return nil, err
	}
	d := new(v1.OrderPayRes)
	d.Successful = true
	return d, nil
}

func (a *cOrder) AdminList(ctx context.Context, req *admin.OrderListReq) (res *admin.OrderListRes, err error) {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Size = 20
	}
	res, err = service.Order().AdminList(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}

// AdminDeliver 发货
func (a *cOrder) AdminDeliver(ctx context.Context, req *admin.OrderDeliverReq) (res *admin.OrderDeliverRes, err error) {
	err = service.Order().AdminUpdStatus(ctx, req.OrderId, 2)
	if err != nil {
		return nil, err
	}
	res = new(admin.OrderDeliverRes)
	res.Successful = true
	return
}

// AdminReturn 归还书
func (a *cOrder) AdminReturn(ctx context.Context, req *admin.OrderReturnReq) (res *admin.OrderReturnRes, err error) {
	err = service.Order().AdminUpdStatus(ctx, req.OrderId, 3)
	if err != nil {
		return nil, err
	}
	res = new(admin.OrderReturnRes)
	res.Successful = true
	return
}

// AdminPostpone 延期归还书 默认30天
func (a *cOrder) AdminPostpone(ctx context.Context, req *admin.OrderPostponeReq) (res *admin.OrderPostponeRes, err error) {
	res, err = service.Order().AdminPostpone(ctx, req)
	if err != nil {
		return nil, err
	}
	res = new(admin.OrderPostponeRes)
	res.Successful = true
	return
}

func (a *cOrder) List(ctx context.Context, req *v1.OrderListReq) (res *v1.OrderListRes, err error) {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Size = 20
	}
	res, err = service.Order().List(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}

// Return 归还书
func (a *cOrder) Return(ctx context.Context, req *v1.OrderReturnReq) (res *v1.OrderReturnRes, err error) {
	err = service.Order().AdminUpdStatus(ctx, req.OrderId, 3)
	if err != nil {
		return nil, err
	}
	res = new(v1.OrderReturnRes)
	res.Successful = true
	return
}
