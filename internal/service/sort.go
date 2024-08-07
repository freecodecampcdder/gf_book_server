// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goFrameMyServer/api/v1/admin"
	v1 "goFrameMyServer/api/v1/app"
)

type (
	ISort interface {
		List(ctx context.Context) (res *v1.SortListRes, err error)
		AdminList(ctx context.Context, req *admin.SortListReq) (res *admin.SortListRes, err error)
		AdminAdd(ctx context.Context, req *admin.SortAddReq) (res *admin.SortAddRes, err error)
		AdminUpd(ctx context.Context, req *admin.SortUpdReq) (res *admin.SortUpdRes, err error)
		AdminDel(ctx context.Context, req *admin.SortDelReq) (res *admin.SortDelRes, err error)
	}
)

var (
	localSort ISort
)

func Sort() ISort {
	if localSort == nil {
		panic("implement not found for interface ISort, forgot register?")
	}
	return localSort
}

func RegisterSort(i ISort) {
	localSort = i
}
