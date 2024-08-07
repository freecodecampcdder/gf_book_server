package admin

import (
	"github.com/gogf/gf/v2/frame/g"
)

type SortListReq struct {
	g.Meta `path:"/sort/list" method:"get" tags:"后台-分类" summary:"后台-分类列表"`
}
type SortListRes struct {
	List interface{} `json:"list"`
}
type SortAddReq struct {
	g.Meta `path:"/sort/add" method:"post" tags:"后台-分类" summary:"后台-分类添加"`
	Title  string `json:"title"`
}
type SortAddRes struct {
	Successful bool `json:"successful"`
}

type SortUpdReq struct {
	g.Meta `path:"/sort/upd" method:"put" tags:"后台-分类" summary:"后台-分类修改"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
}
type SortUpdRes struct {
	Successful bool `json:"successful"`
}

type SortDelReq struct {
	g.Meta `path:"/sort/del" method:"delete" tags:"后台-分类" summary:"后台-标签删除"`
	Id     []int `json:"id"`
}
type SortDelRes struct {
	Successful bool `json:"successful"`
}
