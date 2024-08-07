package admin

import (
	"github.com/gogf/gf/v2/frame/g"
	"goFrameMyServer/api/v1"
)

// CommentListReq 评论列表
type CommentListReq struct {
	g.Meta `path:"/comment/list" method:"get" tags:"后台-评论" summary:"后台-评论列表"`
	api.CommonPaginationReq
}
type CommentListRes struct {
	api.CommonPaginationRes
}

// CommentDelReq 删除评论
type CommentDelReq struct {
	g.Meta `path:"/comment/del" method:"delete" tags:"后台-评论" summary:"后台-评论删除"`
	Id     []int64 `json:"id"`
}
type CommentDelRes struct {
	Successful bool `json:"successful"`
}
