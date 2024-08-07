package app

import (
	"github.com/gogf/gf/v2/frame/g"
	"goFrameMyServer/api/v1"
)

type BookListReq struct {
	g.Meta `path:"/book/list" method:"get" tags:"图书" summary:"图书列表"`
	Name   string `json:"name"`
	Order  int    `json:"order"`
	Sort   int    `json:"sort"`
	Tag    int    `json:"tag"`

	api.CommonPaginationReq
}

type BookListRes struct {
	api.CommonPaginationRes
}

type BookDetailsReq struct {
	g.Meta `path:"/book/details" method:"get" tags:"图书" summary:"图书详情"`
	Id     int64 `json:"id"`
	UserId int64 `json:"user_id"`
}

type BookDetailsRes struct {
	Id           int64   `json:"id"`
	LanguageId   int64   `json:"language_id"`
	Language     string  `json:"language"`
	Title        string  `json:"title"`
	Author       string  `json:"author"`
	Cover        string  `json:"cover"`
	Translator   string  `json:"translator"`
	Description  string  `json:"description"`
	Status       int     `json:"status"`
	Isbn         string  `json:"isbn"`
	Press        string  `json:"press"`
	PressTime    int64   `json:"press_time"`
	PageNum      int     `json:"page_num"`
	WishNum      int     `json:"wish_num"`
	CollectNum   int     `json:"collect_num"`
	BorrowNum    int     `json:"borrow_num"`
	Price        float64 `json:"price"`
	BuyPrice     float64 `json:"buy_price"`
	InventoryNum int     `json:"inventory_num"`
	Sort         []Sort  `json:"sort"`
	Tag          []Tag   `json:"tag"`
	IsWish       bool    `json:"is_wish"`
	IsCollect    bool    `json:"is_collect"`
}

type BookRecommendedListReq struct {
	g.Meta `path:"/book/recommended" method:"get" tags:"图书" summary:"推荐图书"`
	api.CommonPaginationReq
}

type BookRecommendedListRes struct {
	api.CommonPaginationRes
}
