package model

type NotifyList struct {
	Id        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}
