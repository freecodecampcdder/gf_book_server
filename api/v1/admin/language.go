package admin

import "github.com/gogf/gf/v2/frame/g"

type LanguageListReq struct {
	g.Meta `path:"/language/list" method:"get" tags:"后台-语种" summary:"后台-语种列表"`
}
type LanguageListRes struct {
	List interface{} `json:"list"`
}
