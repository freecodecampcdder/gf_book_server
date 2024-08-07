package app

import (
	"github.com/gogf/gf/v2/frame/g"
)

type SortListReq struct {
	g.Meta `path:"/sort/list" method:"get" tags:"分类" summary:"分类列表"`
}

type SortListRes struct {
	List []Sort `json:"list"`
}
type Sort struct {
	Id    int64  `json:"id"`
	Pid   int64  `json:"pid"`
	Title string `json:"title"`
}
