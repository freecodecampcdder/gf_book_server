package consts

const (
	ErrLoginFailMsg      = "登录失败，账号或密码错误"
	GTokenFrontendPrefix = "User:"      //gtoken登录 前台用户 前缀区分
	GTokenAdminPrefix    = "Admin:"     //gtoken登录 管理后台 前缀区分
	ContextKey           = "ContextKey" // 上下文变量存储键名，前后端系统共享

	// TokenType 登录相关
	TokenType = "Bearer"

	CacheModeRedis     = 2
	BackendServerName  = "GXS图书管理系统"
	BackendServerV     = "1.0.0"
	ErrLoginFaulMsg    = "登录失败，账号或密码错误"
	GTokenExpireIn     = 10 * 24 * 60 * 60
	MultiLogin         = true
	FrontendMultiLogin = true

	//for admin
	CtxAdminId      = "CtxAdminId"
	CtxAdminName    = "CtxAdminName"
	CtxAdminIsAdmin = "CtxAdminIsAdmin"
	// CtxUserId for user
	CtxUserId     = "CtxUserId"
	CtxUserName   = "CtxUserName"
	CtxUserAvatar = "CtxUserAvatar"
	CtxUserSex    = "CtxUserSex"
	CtxUserSign   = "CtxUserSign"
	CtxUserStatus = "CtxUserStatus"

	QiNiuToken     = "QiNiuToken:"
	QiNiuTokenTime = 60 * 60
)
