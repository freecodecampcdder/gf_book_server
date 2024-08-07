package admin

import (
	"github.com/gogf/gf/v2/frame/g"
)

type IndexReq struct {
	g.Meta `path:"/index/data" method:"get" tags:"后台-主页" summary:"后台-最上排数据"`
}
type IndexRes struct {
	BookTotal     int             `json:"book_total"`
	BookWeeks     int             `json:"book_weeks"`
	BorrowBook    int             `json:"borrow_book"`
	BorrowPeople  int             `json:"borrow_people"`
	AlsoBook      int             `json:"also_book"`
	AlsoPeople    int             `json:"also_people"`
	OverdueBook   int             `json:"overdue_book"`
	OverduePeople int             `json:"overdue_people"`
	LatelyLine    []InterviewData `json:"lately_line"`
	BookData      []*BookData     `json:"book_data"`
	SortData      []*SortData     `json:"sort_data"`
}

type InterviewData struct {
	Id             int    `json:"id"`
	T              string `json:"t"`
	InterviewCount int    `json:"interview_count"`
}

type BookData struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	BorrowNum int    `json:"borrow_num"`
}
type SortData struct {
	Id     int    `json:"id"`
	SortId int    `json:"sort_id"`
	Title  string `json:"title"`
	Num    int    `json:"num"`
}
