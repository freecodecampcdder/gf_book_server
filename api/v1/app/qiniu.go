package app

import "github.com/gogf/gf/v2/frame/g"

type QiNiuUpReq struct {
	g.Meta `path:"/qiniu/token" method:"get" tags:"七牛" summary:"获取上传token"`
}
type QiNiuUpRes struct {
	Token string `json:"token"`
}
