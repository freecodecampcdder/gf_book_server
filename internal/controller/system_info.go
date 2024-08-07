package controller

import (
	"context"
	"goFrameMyServer/api/v1/admin"
	"goFrameMyServer/internal/service"
)

var SystemInfo = cSystemInfo{}

type cSystemInfo struct {
}

func (s *cSystemInfo) Info(ctx context.Context, req *admin.SystemInfoReq) (res *admin.SystemInfoRes, err error) {
	res, err = service.SystemInfo().Info(ctx)
	if err != nil {
		return nil, err
	}
	return
}
