package controller

import (
	"context"
	"goFrameMyServer/api/v1/admin"
	"goFrameMyServer/internal/service"
)

var Language = cLanguage{}

type cLanguage struct {
}

func (s *cLanguage) AdminList(ctx context.Context, req *admin.LanguageListReq) (res *admin.LanguageListRes, err error) {
	res, err = service.Language().AdminList(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}
