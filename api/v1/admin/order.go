package admin

import (
	"github.com/gogf/gf/v2/frame/g"
	"goFrameMyServer/api/v1"
)

type OrderListReq struct {
	g.Meta `path:"/order/list" method:"get" tags:"后台-订单" summary:"后台-订单列表"`
	Status int `json:"status"`
	api.CommonPaginationReq
}
type OrderListRes struct {
	api.CommonPaginationRes
}

type OrderDeliverReq struct {
	g.Meta  `path:"/order/deliver" method:"post" tags:"后台-订单" summary:"后台-订单发货"`
	OrderId int `json:"order_id"`
}
type OrderDeliverRes struct {
	Successful bool `json:"successful"`
}

type OrderReturnReq struct {
	g.Meta  `path:"/order/return" method:"post" tags:"后台-订单" summary:"后台-订单还书"`
	OrderId int `json:"order_id"`
}
type OrderReturnRes struct {
	Successful bool `json:"successful"`
}

type OrderPostponeReq struct {
	g.Meta  `path:"/order/postpone" method:"post" tags:"后台-订单" summary:"后台-延期"`
	OrderId int `json:"order_id"`
}
type OrderPostponeRes struct {
	Successful bool `json:"successful"`
}
