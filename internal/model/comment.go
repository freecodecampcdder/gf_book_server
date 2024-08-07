package model

type CommentList struct {
	Id        int    `json:"id"`
	UserId    int64  `json:"user_id"`
	Avatar    string `json:"avatar"`
	Nickname  string `json:"nickname"`
	Context   string `json:"context"`
	Like      int    `json:"like"`
	CreatedAt string `json:"created_at"`
}

type AdminCommentList struct {
	Id        int    `json:"id"`
	UserId    int64  `json:"user_id"`
	Nickname  string `json:"nickname"`
	Context   string `json:"context"`
	BookId    int    `json:"book_id"`
	BookTitle string `json:"book_title"`
	CreatedAt string `json:"created_at"`
}

type MyCommentList struct {
	CommentList
	Title string `json:"title"`
}
