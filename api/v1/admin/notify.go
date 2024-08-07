package admin

import (
	"github.com/gogf/gf/v2/frame/g"
	"goFrameMyServer/api/v1"
)

type NotifyAddReq struct {
	g.Meta  `path:"/notify/add" method:"post" tags:"后台-通知" summary:"后台-添加通知"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
type NotifyAddRes struct {
	Successful bool `json:"successful"`
}
type NotifyUpdReq struct {
	g.Meta  `path:"/notify/upd" method:"put" tags:"后台-通知" summary:"后台-修改通知"`
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
type NotifyUpdRes struct {
	Successful bool `json:"successful"`
}

type NotifyDelReq struct {
	g.Meta `path:"/notify/del" method:"delete" tags:"后台-通知" summary:"后台-删除通知"`
	Id     []int `json:"id"`
}
type NotifyDelRes struct {
	Successful bool `json:"successful"`
}

type NotifyListReq struct {
	g.Meta `path:"/notify/list" method:"get" tags:"后台-通知" summary:"后台-通知列表"`
	api.CommonPaginationReq
}
type NotifyListRes struct {
	api.CommonPaginationRes
}
