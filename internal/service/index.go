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
	IIndex interface {
		IndexData(ctx context.Context, req *admin.IndexReq) (res *admin.IndexRes, err error)
	}
)

var (
	localIndex IIndex
)

func Index() IIndex {
	if localIndex == nil {
		panic("implement not found for interface IIndex, forgot register?")
	}
	return localIndex
}

func RegisterIndex(i IIndex) {
	localIndex = i
}
