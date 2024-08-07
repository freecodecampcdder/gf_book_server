package app

import (
	"github.com/gogf/gf/v2/frame/g"
	"goFrameMyServer/api/v1"
)

type CollectClickReq struct {
	g.Meta `path:"/collect" method:"post" tags:"收藏" summary:"收藏"`
	BookId int64 `json:"book_id"`
}
type CollectClickRes struct {
	Successful bool `json:"successful"`
}

type CollectListReq struct {
	g.Meta `path:"/collect/list" method:"get" tags:"收藏" summary:"收藏列表"`
	api.CommonPaginationReq
}
type CollectListRes struct {
	api.CommonPaginationRes
}

type CollectDelReq struct {
	g.Meta `path:"/collect/del" method:"delete" tags:"收藏" summary:"收藏删除"`
	Id     int `json:"id"`
}
type CollectDelRes struct {
	Successful bool `json:"successful"`
}
