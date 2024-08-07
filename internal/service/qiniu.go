// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IQiniu interface {
		GetToken(ctx context.Context) (string, error)
	}
)

var (
	localQiniu IQiniu
)

func Qiniu() IQiniu {
	if localQiniu == nil {
		panic("implement not found for interface IQiniu, forgot register?")
	}
	return localQiniu
}

func RegisterQiniu(i IQiniu) {
	localQiniu = i
}
