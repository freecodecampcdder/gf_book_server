package app

import (
	"github.com/gogf/gf/v2/frame/g"
	"goFrameMyServer/api/v1"
)

type AddressAddReq struct {
	g.Meta `path:"/address/add" method:"post" tags:"地址" summary:"添加地址"`
	AddressData
}
type AddressAddRes struct {
	Id int64 `json:"id"`
}

type AddressData struct {
	ReceiverName   string `json:"receiver_name"`
	ReceiverPhone  string `json:"receiver_phone"`
	AddressContent string `json:"address_content"`
	Status         int    `json:"status"`
}
type AddressUpdReq struct {
	g.Meta `path:"/address/upd" method:"put" tags:"地址" summary:"修改地址"`
	Id     int64 `json:"id"`
	AddressData
}

type AddressUpdRes struct {
	Successful bool `json:"successful"`
}

type AddressDelReq struct {
	g.Meta `path:"/address/del" method:"delete" tags:"地址" summary:"删除地址"`
	Id     int64 `json:"id"`
}
type AddressDelRes struct {
	Successful bool `json:"successful"`
}

type AddressListReq struct {
	g.Meta `path:"/address/list" method:"get" tags:"地址" summary:"地址列表"`
	//api.CommonPaginationReq
}
type AddressListRes struct {
	api.CommonPaginationRes
}
