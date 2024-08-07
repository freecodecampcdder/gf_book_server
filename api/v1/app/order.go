package app

import (
	"github.com/gogf/gf/v2/frame/g"
	"goFrameMyServer/api/v1"
)

type OrderPreloadReq struct {
	g.Meta `path:"/order/preload" method:"post" tags:"订单" summary:"订单-预览"`
	BookId int64 `json:"book_id"`
	Way    int   `json:"way"`
}

type OrderPreloadRes struct {
	BookId int     `json:"book_id"`
	Cover  string  `json:"cover"`
	Title  string  `json:"title"`
	Price  float64 `json:"price"`
}

type OrderAddReq struct {
	g.Meta    `path:"/order/add" method:"post" tags:"订单" summary:"订单-添加"`
	AddressId int `json:"address_id"`
	BookId    int `json:"book_id"`
	Way       int `json:"way"`
}

type OrderAddRes struct {
	Id int64 `json:"id"`
}

type OrderDetailsReq struct {
	g.Meta  `path:"/order/details" method:"get" tags:"订单" summary:"订单-详情"`
	OrderId int `json:"order_id"`
}

type OrderDetailsRes struct {
	Id     int `json:"order_id"`
	Price  int `json:"price"`
	Way    int `json:"way"`
	Status int `json:"status"`
}

type OrderPayReq struct {
	g.Meta  `path:"/order/pay" method:"post" tags:"订单" summary:"订单-支付"`
	OrderId int `json:"order_id"`
}

type OrderPayRes struct {
	Successful bool `json:"successful"`
}

type OrderListReq struct {
	g.Meta `path:"/order/list" method:"get" tags:"订单" summary:"订单-列表"`
	Status int `json:"status"`
	api.CommonPaginationReq
}

type OrderListRes struct {
	api.CommonPaginationRes
}

type OrderReturnReq struct {
	g.Meta  `path:"/order/return" method:"post" tags:"订单" summary:"订单还书"`
	OrderId int `json:"order_id"`
}

type OrderReturnRes struct {
	Successful bool `json:"successful"`
}
