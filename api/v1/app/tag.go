package app

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TagListReq struct {
	g.Meta `path:"/tag/list" method:"get" tags:"标签" summary:"标签列表"`
}

type TagListRes struct {
	List []Tag `json:"list"`
}
type Tag struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}
