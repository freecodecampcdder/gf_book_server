package cmd

import (
	"context"
	"github.com/duke-git/lancet/v2/netutil"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"goFrameMyServer/api/v1/admin"
	"goFrameMyServer/api/v1/app"
	"goFrameMyServer/internal/consts"
	"goFrameMyServer/internal/dao"
	"goFrameMyServer/internal/model/entity"
	"goFrameMyServer/utility"
	"goFrameMyServer/utility/response"
	"strconv"
)

// 管理后台相关
func StartBackendGToken() (gfAdminToken *gtoken.GfToken, err error) {
	gfAdminToken = &gtoken.GfToken{
		CacheMode:        consts.CacheModeRedis,
		ServerName:       consts.BackendServerName,
		LoginPath:        "/login",
		LoginBeforeFunc:  loginFunc,
		LoginAfterFunc:   loginAfterFunc,
		LogoutPath:       "/user/logout",
		AuthPaths:        g.SliceStr{"/backend/admin/info"},
		AuthExcludePaths: g.SliceStr{"/admin/user/info", "/admin/system/user/info"}, // 不拦截路径 /user/info,/system/user/info,/system/user,
		AuthAfterFunc:    authAfterFunc,
		MultiLogin:       consts.MultiLogin,
	}
	//todo 去掉全局校验，只用cmd中的路由组校验
	err = gfAdminToken.Start()
	return
}

// 前台登录gtoken相关
func StartFrontendGToken() (gfFrontendToken *gtoken.GfToken, err error) {
	gfFrontendToken = &gtoken.GfToken{
		CacheMode:       consts.CacheModeRedis,
		ServerName:      consts.BackendServerName,
		LoginPath:       "/login",
		LoginBeforeFunc: loginFuncFrontend,
		LoginAfterFunc:  loginAfterFuncFrontend,
		LogoutPath:      "/user/logout",
		//AuthPaths:       g.SliceStr{"/app/book/details"},
		AuthAfterFunc: authAfterFuncFrontend,
		MultiLogin:    consts.FrontendMultiLogin,
	}

	return
}

func loginFunc(r *ghttp.Request) (string, interface{}) {
	name := r.Get("user_name").String()
	password := r.Get("password").String()
	ctx := context.TODO()

	if name == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFaulMsg))
		r.ExitAll()
	}

	//验证账号密码是否正确
	adminInfo := entity.User{}
	err := dao.User.Ctx(ctx).Where(dao.User.Columns().UserName, name).Where(dao.User.Columns().Role, "admin").Scan(&adminInfo)
	if err != nil {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFaulMsg))
		r.ExitAll()
	}
	if utility.EncryptPassword(password, adminInfo.UserSalt) != adminInfo.Password {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFaulMsg))
		r.ExitAll()
	}
	// 唯一标识，扩展参数user data
	return consts.GTokenAdminPrefix + strconv.Itoa(int(adminInfo.Id)), adminInfo
}

// 自定义的登录之后的函数
func loginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		respData.Code = 0
		r.Response.WriteJson(respData)
		return
	} else {
		respData.Code = 1
		//获得登录用户id
		userKey := respData.GetString("userKey")
		adminId := gstr.StrEx(userKey, consts.GTokenAdminPrefix)
		//根据id获得登录用户其他信息
		adminInfo := entity.User{}
		err := dao.User.Ctx(context.TODO()).WherePri(adminId).Scan(&adminInfo)
		if err != nil {
			return
		}
		//获取ip以及访问头
		ip := netutil.GetRequestPublicIp(r.Request)
		userAgent := r.UserAgent()
		_, _ = dao.LoginLog.Ctx(context.TODO()).Data(g.Map{"ip": ip, "user_agent": userAgent, "user_id": adminId}).Insert()
		data := &admin.LoginRes{
			Type:     consts.TokenType,
			Token:    respData.GetString("token"),
			ExpireIn: consts.GTokenExpireIn, //单位秒,
			IsAdmin:  1,
			Name:     adminInfo.Nickname,
			Avatar:   adminInfo.Avatar,
		}
		response.JsonExit(r, 0, "", data)
	}
	return
}

// for 前台项目
func loginFuncFrontend(r *ghttp.Request) (string, interface{}) {
	name := r.Get("user_name").String()
	password := r.Get("password").String()
	ctx := context.TODO()

	if name == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFailMsg))
		r.ExitAll()
	}

	//验证账号密码是否正确
	userInfo := entity.User{}
	err := dao.User.Ctx(ctx).Where(dao.User.Columns().UserName, name).Scan(&userInfo)
	if err != nil {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFailMsg))
		r.ExitAll()
	}
	if utility.EncryptPassword(password, userInfo.UserSalt) != userInfo.Password {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFailMsg))
		r.ExitAll()
	}
	// 唯一标识，扩展参数user data
	return consts.GTokenFrontendPrefix + strconv.Itoa(int(userInfo.Id)), userInfo
}

// 自定义的登录之后的函数 for前台项目
func loginAfterFuncFrontend(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		respData.Code = 0
		r.Response.WriteJson(respData)
		return
	} else {
		respData.Code = 1
		//获得登录用户id
		userKey := respData.GetString("userKey")
		userId := gstr.StrEx(userKey, consts.GTokenFrontendPrefix)
		//根据id获得登录用户其他信息
		userInfo := entity.User{}
		err := dao.User.Ctx(context.TODO()).WherePri(userId).Scan(&userInfo)
		if err != nil {
			return
		}
		//获取ip以及访问头
		ip := netutil.GetRequestPublicIp(r.Request)
		userAgent := r.UserAgent()
		_, _ = dao.LoginLog.Ctx(context.TODO()).Data(g.Map{"ip": ip, "user_agent": userAgent, "user_id": userId}).Insert()
		data := &app.LoginRes{
			Type:     consts.TokenType,
			Token:    respData.GetString("token"),
			ExpireIn: consts.GTokenExpireIn, //单位秒,
		}
		data.Id = userInfo.Id
		data.Name = userInfo.Nickname
		data.Avatar = userInfo.Avatar
		data.Sign = userInfo.Description
		data.Status = uint8(userInfo.Status)
		response.JsonExit(r, 0, "", data)
	}
	return
}

// 登录鉴权中间件for后台
func authAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	var adminInfo entity.User
	err := gconv.Struct(respData.GetString("data"), &adminInfo)
	if err != nil {
		response.Auth(r)
		return
	}
	//todo 这里可以写账号前置校验、是否被拉黑、有无权限等逻辑
	r.SetCtxVar(consts.CtxAdminId, adminInfo.Id)
	r.SetCtxVar(consts.CtxAdminName, adminInfo.UserName)
	r.SetCtxVar(consts.CtxAdminIsAdmin, adminInfo.Role)
	r.Middleware.Next()
}

// 登录鉴权中间件for前台
func authAfterFuncFrontend(r *ghttp.Request, respData gtoken.Resp) {
	var userInfo entity.User
	err := gconv.Struct(respData.GetString("data"), &userInfo)
	if err != nil {
		response.Auth(r)
		return
	}
	//todo 这里可以写账号前置校验、是否被拉黑、有无权限等逻辑
	r.SetCtxVar(consts.CtxUserId, userInfo.Id)
	r.SetCtxVar(consts.CtxUserName, userInfo.Nickname)
	r.SetCtxVar(consts.CtxUserAvatar, userInfo.Avatar)
	r.SetCtxVar(consts.CtxUserSign, userInfo.Description)
	r.SetCtxVar(consts.CtxUserStatus, userInfo.Status)
	r.Middleware.Next()
}
