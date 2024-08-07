package language

import (
	"context"
	"goFrameMyServer/api/v1/admin"
	"goFrameMyServer/internal/dao"
	"goFrameMyServer/internal/model"
	"goFrameMyServer/internal/service"
)

type sLanguage struct {
}

func init() {
	service.RegisterLanguage(New())
}

func New() *sLanguage {
	return &sLanguage{}
}

func (s *sLanguage) AdminList(ctx context.Context, req *admin.LanguageListReq) (res *admin.LanguageListRes, err error) {
	var sort []*model.LanguageAdminList
	var m = dao.Language.Ctx(ctx)
	res = &admin.LanguageListRes{
		List: []*model.LanguageAdminList{},
	}
	if err = m.Scan(&sort); err != nil {
		return res, err
	}
	if len(sort) == 0 {
		return res, err
	}
	res.List = sort
	return
}
