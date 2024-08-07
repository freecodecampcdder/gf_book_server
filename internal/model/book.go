package model

type BookList struct {
	Id         string  `json:"id"`
	Language   string  `json:"language"`
	Title      string  `json:"title"`
	Author     string  `json:"author"`
	Cover      string  `json:"cover"`
	Translator string  `json:"translator"`
	Price      float64 `json:"price"`
}

type AdminBookList struct {
	Id           int64   `json:"id"`
	Title        string  `json:"title"`
	Author       string  `json:"author"`
	Cover        string  `json:"cover"`
	Translator   string  `json:"translator"`
	Press        string  `json:"press"`
	PageNum      int     `json:"page_num"`
	Price        float64 `json:"price"`
	BuyPrice     float64 `json:"buy_price"`
	Isbn         string  `json:"isbn"`
	PressTime    string  `json:"press_time"`
	Status       int     `json:"status"`
	InventoryNum int64   `json:"inventory_num"`
	Description  string  `json:"description"`
}

type BookRecommendedList struct {
	Id         int    `json:"id"`
	Cover      string `json:"cover"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	Translator string `json:"translator"`
}
