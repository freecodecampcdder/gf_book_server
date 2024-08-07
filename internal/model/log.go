package model

type AdminLoginLogList struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Nickname  string `json:"nickname"`
	Ip        string `json:"ip"`
	UserAgent string `json:"user_agent"`
	CreatedAt string `json:"created_at"`
}
