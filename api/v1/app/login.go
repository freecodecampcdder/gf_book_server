package app

import "github.com/gogf/gf/v2/frame/g"

type UserLoginReq struct {
	g.Meta   `path:"/login" method:"post" tags:"用户登录" summary:"用户登录"`
	UserName string `json:"user_name" v:"required#用户名不可为空" dc:"用户名"`
	Password string `json:"password" v:"required#密码不可为空" dc:"密码"`
}
type UserLoginRes struct {
	Token string `json:"token"`
}
