// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goFrameMyServer/api/v1/admin"
)

type (
	ILanguage interface {
		AdminList(ctx context.Context, req *admin.LanguageListReq) (res *admin.LanguageListRes, err error)
	}
)

var (
	localLanguage ILanguage
)

func Language() ILanguage {
	if localLanguage == nil {
		panic("implement not found for interface ILanguage, forgot register?")
	}
	return localLanguage
}

func RegisterLanguage(i ILanguage) {
	localLanguage = i
}
