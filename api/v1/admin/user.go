package admin

import (
	"github.com/gogf/gf/v2/frame/g"
	"goFrameMyServer/api/v1"
)

type UserListReq struct {
	g.Meta `path:"/user/list" method:"get" tags:"后台-用户" summary:"后台-用户列表"`
	Name   string `json:"name"`
	api.CommonPaginationReq
}

type UserListRes struct {
	api.CommonPaginationRes
}

type UserAddReq struct {
	g.Meta `path:"/user/add" method:"post" tags:"后台-用户" summary:"后台-用户添加"`
	UserInfo
}
type UserAddRes struct {
	Successful bool `json:"successful"`
}
type UserUpdReq struct {
	g.Meta `path:"/user/upd" method:"put" tags:"后台-用户" summary:"后台-用户修改"`
	Id     int64 `json:"Id"`
	UserInfo
}
type UserUpdRes struct {
	Successful bool `json:"successful"`
}
type UserDelReq struct {
	g.Meta `path:"/user/del" method:"delete" tags:"后台-用户" summary:"后台-用户删除"`
	Id     []int64 `json:"Id"`
}
type UserDelRes struct {
	Successful bool `json:"successful"`
}

type UserUpdPasswordReq struct {
	g.Meta   `path:"/user/upd/password" method:"put" tags:"后台-用户" summary:"后台-修改密码"`
	Id       int64  `json:"Id"`
	Password string `json:"password"`
}
type UserUpdPasswordRes struct {
	Successful bool `json:"successful"`
}

type UserGetInfoReq struct {
	g.Meta `path:"/user/one" method:"get" tags:"后台-用户" summary:"后台-用户详情"`
	Id     int64 `json:"Id"`
}
type UserGetInfoRes struct {
	Id int64 `json:"Id"`
	UserInfo
}

type UserInfo struct {
	Role     string `json:"role"`
	Status   int    `json:"status"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
}
