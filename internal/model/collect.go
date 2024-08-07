package model

type CollectList struct {
	Id         int64  `json:"id"`
	BookId     int64  `json:"book_id"`
	Title      string `json:"title"`
	Cover      string `json:"cover"`
	Author     string `json:"author"`
	Translator string `json:"translator"`
}
