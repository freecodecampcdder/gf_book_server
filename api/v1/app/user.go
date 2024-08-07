package app

import "github.com/gogf/gf/v2/frame/g"

type UserRegisterReq struct {
	g.Meta   `path:"/user/register" method:"post" tags:"注册" summary:"用户注册接口"`
	UserName string `json:"user_name" v:"required#用户名不可为空" dc:"用户名"`
	Password string `json:"password" v:"required#密码不可为空" dc:"密码"`
}

type UserRegisterRes struct {
	UserId uint `json:"user_id"`
}

// LoginRes for token
type LoginRes struct {
	Type     string `json:"type"`
	Token    string `json:"token"`
	ExpireIn int    `json:"expire_in"`
	UserInfoBase
}

// UserInfoBase 可以复用的,一定要抽取出来
type UserInfoBase struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Sex    uint8  `json:"sex"`
	Sign   string `json:"sign"`
	Status uint8  `json:"status"`
}

type UserInfoReq struct {
	g.Meta `path:"/user/info" method:"get" tags:"用户" summary:"用户详情接口"`
}

type UserInfoRes struct {
	Id          uint   `json:"id"`
	Avatar      string `json:"avatar"`
	Nickname    string `json:"nickname"`
	Mobile      string `json:"mobile"`
	Email       string `json:"email"`
	Description string `json:"description"`
	PushSwitch  int    `json:"push_switch"`
}

type UserUpdInfoReq struct {
	g.Meta      `path:"/user/info" method:"put" tags:"用户" summary:"用户更换头像"`
	Avatar      string `json:"avatar"`
	Nickname    string `json:"nickname"`
	Mobile      string `json:"mobile"`
	Email       string `json:"email"`
	Description string `json:"description"`
}
type UserUpdAvatarRes struct {
	Successful bool `json:"successful"`
}

type UserUpdPasswordReq struct {
	g.Meta      `path:"/user/password" method:"put" tags:"用户" summary:"用户更换密码"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}
type UserUpdPasswordRes struct {
	Successful bool `json:"successful"`
}
type UserPushSwitchReq struct {
	g.Meta `path:"/user/push" method:"put" tags:"用户" summary:"用户推送开关"`
}
type UserPushSwitchRes struct {
	Successful bool `json:"successful"`
}

type UserPointReq struct {
	g.Meta `path:"/user/point" method:"get" tags:"用户" summary:"获取用户积分"`
}

type UserPointRes struct {
	Points int `json:"points"`
}
