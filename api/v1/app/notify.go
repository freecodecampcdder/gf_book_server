package app

import "github.com/gogf/gf/v2/frame/g"

type NotifyListReq struct {
	g.Meta `path:"/notify/list" method:"get" tags:"通知" summary:"通知列表(前五条)"`
}
type NotifyListRes struct {
	List interface{} `json:"list"`
}
