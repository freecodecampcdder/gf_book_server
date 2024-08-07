package controller

import (
	"context"
	"goFrameMyServer/api/v1/admin"
	v1 "goFrameMyServer/api/v1/app"
	"goFrameMyServer/internal/service"
)

var Sort = cSort{}

type cSort struct {
}

func (s *cSort) List(ctx context.Context, req *v1.SortListReq) (res *v1.SortListRes, err error) {
	res, err = service.Sort().List(ctx)
	if err != nil {
		return nil, err
	}
	return
}

func (s *cSort) AdminList(ctx context.Context, req *admin.SortListReq) (res *admin.SortListRes, err error) {
	res, err = service.Sort().AdminList(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *cSort) AdminAdd(ctx context.Context, req *admin.SortAddReq) (res *admin.SortAddRes, err error) {
	res, err = service.Sort().AdminAdd(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *cSort) AdminUpd(ctx context.Context, req *admin.SortUpdReq) (res *admin.SortUpdRes, err error) {
	res, err = service.Sort().AdminUpd(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *cSort) AdminDel(ctx context.Context, req *admin.SortDelReq) (res *admin.SortDelRes, err error) {
	res, err = service.Sort().AdminDel(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}
