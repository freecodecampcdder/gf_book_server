package admin

import (
	"github.com/gogf/gf/v2/frame/g"
	"goFrameMyServer/api/v1"
)

type LoginLogListReq struct {
	g.Meta `path:"/log/login" method:"get" tags:"后台-登录日志" summary:"后台-登录日志列表"`
	api.CommonPaginationReq
}
type LoginLogListRes struct {
	api.CommonPaginationRes
}

type OpLogListReq struct {
	g.Meta `path:"/log/op" method:"get" tags:"后台-操作日志" summary:"后台-操作日志列表"`
	api.CommonPaginationReq
}
type OpLogListRes struct {
	api.CommonPaginationRes
}
