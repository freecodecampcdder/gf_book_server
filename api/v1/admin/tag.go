package admin

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TagListReq struct {
	g.Meta `path:"/tag/list" method:"get" tags:"后台-标签" summary:"后台-标签列表"`
}
type TagListRes struct {
	List interface{} `json:"list"`
}
type TagAddReq struct {
	g.Meta `path:"/tag/add" method:"post" tags:"后台-标签" summary:"后台-标签添加"`
	Title  string `json:"title"`
}
type TagAddRes struct {
	Successful bool `json:"successful"`
}

type TagUpdReq struct {
	g.Meta `path:"/tag/upd" method:"put" tags:"后台-标签" summary:"后台-标签修改"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
}
type TagUpdRes struct {
	Successful bool `json:"successful"`
}

type TagDelReq struct {
	g.Meta `path:"/tag/del" method:"delete" tags:"后台-标签" summary:"后台-标签删除"`
	Id     []int `json:"id"`
}
type TagDelRes struct {
	Successful bool `json:"successful"`
}
