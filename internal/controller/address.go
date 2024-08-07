package controller

import (
	"context"
	v1 "goFrameMyServer/api/v1/app"
	"goFrameMyServer/internal/service"
)

var Address = cAddress{}

type cAddress struct {
}

func (a *cAddress) Add(ctx context.Context, req *v1.AddressAddReq) (res *v1.AddressAddRes, err error) {
	id, err := service.Address().Add(ctx, req)
	if err != nil {
		return nil, err
	}
	d := new(v1.AddressAddRes)
	d.Id = id
	return d, nil
}

func (a *cAddress) Upd(ctx context.Context, req *v1.AddressUpdReq) (res *v1.AddressUpdRes, err error) {
	err = service.Address().Upd(ctx, req)
	if err != nil {
		return nil, err
	}
	d := new(v1.AddressUpdRes)
	d.Successful = true
	return d, nil
}

func (a *cAddress) Del(ctx context.Context, req *v1.AddressDelReq) (res *v1.AddressDelRes, err error) {
	err = service.Address().Del(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	d := new(v1.AddressDelRes)
	d.Successful = true
	return d, nil
}

func (a *cAddress) List(ctx context.Context, req *v1.AddressListReq) (res *v1.AddressListRes, err error) {
	//if req.Page == 0 {
	//	req.Page = 1
	//}
	//if req.Size == 0 {
	//	req.Size = 20
	//}
	res, err = service.Address().List(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
