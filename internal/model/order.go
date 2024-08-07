package model

type OrderAdminList struct {
	Id           int64  `json:"id"`
	UserId       int64  `json:"user_id"`
	Nickname     string `json:"nickname"`
	AddressName  string `json:"address_name"`
	AddressPhone string `json:"address_phone"`
	BookId       int    `json:"book_id"`
	BookTitle    string `json:"book_title"`
	Price        int    `json:"price"`
	Way          int    `json:"way"`
	Status       int    `json:"status"`
	LendAt       string `json:"lend_at"`
	ReturnAt     string `json:"return_at"`
}

type OrderList struct {
	Id        int64  `json:"id"`
	BookId    int    `json:"book_id"`
	BookTitle string `json:"book_title"`
	Cover     string `json:"cover"`
	Price     int    `json:"price"`
	Way       int    `json:"way"`
	Status    int    `json:"status"`
	LendAt    string `json:"lend_at"`
	ReturnAt  string `json:"return_at"`
}
