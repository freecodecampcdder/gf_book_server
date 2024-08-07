package admin

import (
	"github.com/gogf/gf/v2/frame/g"
	"goFrameMyServer/api/v1"
)

// BookListReq 获取book的列表
type BookListReq struct {
	g.Meta `path:"/book/list" method:"get" tags:"后台-图书" summary:"后台-图书列表"`
	Name   string `json:"name"`
	api.CommonPaginationReq
}
type BookListRes struct {
	api.CommonPaginationRes
}

// BookAddReq 添加书
type BookAddReq struct {
	g.Meta `path:"/book/add" method:"post" tags:"后台-图书" summary:"后台-图书添加"`
	BookInfo
}
type BookAddRes struct {
	Successful bool `json:"successful"`
}

// BookUpdReq 修改书信息
type BookUpdReq struct {
	g.Meta `path:"/book/upd" method:"put" tags:"后台-图书" summary:"后台-图书修改"`
	Id     int64 `json:"id"`
	BookInfo
}
type BookUpdRes struct {
	Successful bool `json:"successful"`
}

// BookDelReq 删除book
type BookDelReq struct {
	g.Meta `path:"/book/del" method:"delete" tags:"后台-图书" summary:"后台-图书删除"`
	Id     []int64 `json:"id"`
}
type BookDelRes struct {
	Successful bool `json:"successful"`
}

// BookDetailReq 书的详情
type BookDetailReq struct {
	g.Meta `path:"/book/one" method:"get" tags:"后台-图书" summary:"后台-图书详情"`
	Id     int64 `json:"id"`
}
type BookDetailRes struct {
	Id int64 `json:"id"`
	BookInfo
}

type BookInfo struct {
	LanguageId   uint   `json:"language_id" dc:"语言" `
	Title        string `json:"title" dc:"书名"`
	Author       string `json:"author" dc:"作者"`
	Cover        string `json:"cover" dc:"封面"`
	Translator   string `json:"translator" dc:"译者"`
	Description  string `json:"description" dc:"描述"`
	Status       int    `json:"status" dc:"状态"`
	Isbn         string `json:"isbn" dc:"ISBN"`
	Press        string `json:"press" dc:"出版社"`
	PressTime    int    `json:"press_time" dc:"出版时间" `
	PageNum      int    `json:"page_num" dc:"页数"`
	Price        int    `json:"price" dc:"租价"`
	BuyPrice     int    `json:"buy_price" dc:"一口价"`
	InventoryNum int    `json:"inventory_num" dc:"库存"`
	Recommended  int    `json:"recommended" dc:"推荐指数"`
	Sort         []uint `json:"sort" dc:"分类"`
	Tag          []uint `json:"tag" dc:"标签"`
}
