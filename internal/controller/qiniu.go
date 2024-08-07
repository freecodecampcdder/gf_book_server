package controller

import (
	"context"
	v1 "goFrameMyServer/api/v1/app"
	"goFrameMyServer/internal/service"
)

var QiNiu = cQiNiu{}

type cQiNiu struct {
}

func (u *cQiNiu) GetToken(ctx context.Context, req *v1.QiNiuUpReq) (res *v1.QiNiuUpRes, err error) {
	out, err := service.Qiniu().GetToken(ctx)
	if err != nil {
		return nil, err
	}
	res = &v1.QiNiuUpRes{
		Token: out,
	}
	return res, nil
}
