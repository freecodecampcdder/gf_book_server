package controller

import (
	"context"
	"goFrameMyServer/api/v1/admin"
	"goFrameMyServer/internal/service"
)

var Index = cIndex{}

type cIndex struct {
}

func (a *cIndex) IndexData(ctx context.Context, req *admin.IndexReq) (res *admin.IndexRes, err error) {
	res, err = service.Index().IndexData(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}
