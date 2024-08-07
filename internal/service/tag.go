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
	ITag interface {
		List(ctx context.Context) (res *v1.TagListRes, err error)
		AdminList(ctx context.Context, req *admin.TagListReq) (res *admin.TagListRes, err error)
		AdminAdd(ctx context.Context, req *admin.TagAddReq) (res *admin.TagAddRes, err error)
		AdminUpd(ctx context.Context, req *admin.TagUpdReq) (res *admin.TagUpdRes, err error)
		AdminDel(ctx context.Context, req *admin.TagDelReq) (res *admin.TagDelRes, err error)
	}
)

var (
	localTag ITag
)

func Tag() ITag {
	if localTag == nil {
		panic("implement not found for interface ITag, forgot register?")
	}
	return localTag
}

func RegisterTag(i ITag) {
	localTag = i
}
