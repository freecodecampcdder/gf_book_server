package controller

import (
	"context"
	"goFrameMyServer/api/v1/admin"
	"goFrameMyServer/internal/service"
)

var Log = cLog{}

type cLog struct {
}

func (a *cLog) AdminLoginList(ctx context.Context, req *admin.LoginLogListReq) (res *admin.LoginLogListRes, err error) {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Size = 20
	}
	res, err = service.Log().AdminLoginList(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *cLog) AdminOpList(ctx context.Context, req *admin.OpLogListReq) (res *admin.OpLogListRes, err error) {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Size = 20
	}
	res, err = service.Log().AdminOpLogList(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
