package app

import (
	"github.com/gogf/gf/v2/frame/g"
	"goFrameMyServer/api/v1"
)

type CommentAddReq struct {
	g.Meta  `path:"/comment/add" method:"post" tags:"评论" summary:"添加评论"`
	BookId  int    `json:"book_id"`
	Context string `json:"context"`
}
type CommentAddRes struct {
	Successful bool `json:"successful"`
}

type CommentListReq struct {
	g.Meta `path:"/comment/list" method:"get" tags:"评论" summary:"评论列表"`
	BookId int `json:"book_id"`
	Order  int `json:"order"`
	api.CommonPaginationReq
}
type CommentListRes struct {
	api.CommonPaginationRes
}

type CommentLikeReq struct {
	g.Meta `path:"/comment/like" method:"post" tags:"评论" summary:"点赞评论"`
	Id     int `json:"id"`
}
type CommentLikeRes struct {
	Successful bool `json:"successful"`
}

type MyCommentListReq struct {
	g.Meta `path:"/comment/my/list" method:"get" tags:"评论" summary:"我的评论"`
	api.CommonPaginationReq
}
type MyCommentListRes struct {
	api.CommonPaginationRes
}
