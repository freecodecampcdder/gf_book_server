package model

type RegisterInput struct {
	UserName string `json:"user_name" description:"用户名" v:"required#用户名必填"`
	Password string `json:"password" description:"密码" v:"password"`
}

type RegisterOutput struct {
	Id uint `json:"id"`
}

type AdminUserList struct {
	Id        int64  `json:"id"`
	UserName  string `json:"user_name"`
	Nickname  string `json:"nickname"`
	Role      string `json:"role"`
	Status    int    `json:"status"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	CreatedAt string `json:"createdAt"`
}
