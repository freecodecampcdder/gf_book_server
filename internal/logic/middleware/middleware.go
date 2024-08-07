package middleware

import (
	"fmt"
	"github.com/duke-git/lancet/v2/netutil"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"goFrameMyServer/internal/model"
	"goFrameMyServer/internal/service"
	"goFrameMyServer/utility/response"
	"time"
)

type sMiddleware struct {
	LoginUrl string // 登录路由地址
}

func init() {
	service.RegisterMiddleware(New())
}

func New() *sMiddleware {
	return &sMiddleware{
		LoginUrl: "/login",
	}
}
func (m *sMiddleware) ResponseHandler(r *ghttp.Request) {
	s := time.Now().UnixNano() / 1e6
	r.Middleware.Next()
	if r.Response.BufferLength() > 0 {
		return
	}
	end := time.Now().UnixNano() / 1e6
	t := end - s
	_ = service.Log().AdminOpLogAdd(r.Context(), r.Request.Method, r.Request.RequestURI, netutil.GetRequestPublicIp(r.Request), t)
	var (
		err             = r.GetError()
		res             = r.GetHandlerResponse()
		code gcode.Code = gcode.CodeOK
	)
	if err != nil {
		code = gerror.Code(err)
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		response.JsonExit(r, code.Code(), err.Error())
	} else {
		response.JsonExit(r, code.Code(), "success", res)
	}
}
func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// 自定义上下文对象
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	// 初始化，务必最开始执行
	customCtx := &model.Context{
		Session: r.Session,
		Data:    make(g.Map),
	}
	service.BizCtx().Init(r, customCtx)
	if userEntity := service.Session().GetUser(r.Context()); userEntity.Id > 0 {
		customCtx.User = &model.ContextUser{
			Id:       userEntity.Id,
			UserName: userEntity.UserName,
			Nickname: userEntity.Nickname,
			Avatar:   userEntity.Avatar,
			Role:     userEntity.Role,
		}
	}
	// 将自定义的上下文对象传递到模板变量中使用
	r.Assigns(g.Map{
		"Context": customCtx,
	})
	// 执行下一步请求逻辑
	r.Middleware.Next()
}

func (s *sMiddleware) OpTestCtx(r *ghttp.Request) {

	fmt.Println(r)
	r.Middleware.Next()
}
