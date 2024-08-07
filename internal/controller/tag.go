package controller

import (
	"context"
	"goFrameMyServer/api/v1/admin"
	v1 "goFrameMyServer/api/v1/app"
	"goFrameMyServer/internal/service"
)

var Tag = cTag{}

type cTag struct {
}

func (s *cTag) List(ctx context.Context, req *v1.TagListReq) (res *v1.TagListRes, err error) {
	res, err = service.Tag().List(ctx)
	if err != nil {
		return nil, err
	}
	return
}

func (s *cTag) AdminList(ctx context.Context, req *admin.TagListReq) (res *admin.TagListRes, err error) {
	res, err = service.Tag().AdminList(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *cTag) AdminAdd(ctx context.Context, req *admin.TagAddReq) (res *admin.TagAddRes, err error) {
	res, err = service.Tag().AdminAdd(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *cTag) AdminUpd(ctx context.Context, req *admin.TagUpdReq) (res *admin.TagUpdRes, err error) {
	res, err = service.Tag().AdminUpd(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}

func (s *cTag) AdminDel(ctx context.Context, req *admin.TagDelReq) (res *admin.TagDelRes, err error) {
	res, err = service.Tag().AdminDel(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}
