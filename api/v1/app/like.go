package app

import (
	"github.com/gogf/gf/v2/frame/g"
	"goFrameMyServer/api/v1"
)

type LikeClickReq struct {
	g.Meta `path:"/like" method:"post" tags:"喜欢" summary:"添加喜欢"`
	BookId int64 `json:"book_id"`
}
type LikeClickRes struct {
	Successful bool `json:"successful"`
}

type LikeListReq struct {
	g.Meta `path:"/like/list" method:"get" tags:"喜欢" summary:"喜欢列表"`
	api.CommonPaginationReq
}
type LikeListRes struct {
	api.CommonPaginationRes
}

type LikeDelReq struct {
	g.Meta `path:"/like/del" method:"delete" tags:"喜欢" summary:"喜欢删除"`
	Id     int `json:"id"`
}
type LikeDelRes struct {
	Successful bool `json:"successful"`
}
