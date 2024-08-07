package cmd

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"goFrameMyServer/internal/cron"
	"goFrameMyServer/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"goFrameMyServer/internal/controller"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 设置进程全局时区
			err = gtime.SetTimeZone("Asia/Shanghai")
			if err != nil {
				panic(err)
			}
			s := g.Server()
			// 启动管理后台gtoken
			gfAdminToken, err := StartBackendGToken()
			if err != nil {
				return err
			}
			//启动定时任务
			cron.InitCron()
			//管理后台路由组
			s.Group("/v1/admin", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
					//service.Middleware().OpTestCtx,
				)

				//不需要登录的路由组绑定
				group.Bind()
				//需要登录的路由组绑定
				group.Group("/", func(group *ghttp.RouterGroup) {
					err := gfAdminToken.Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					group.Bind(
						controller.Book.AdminList,
						controller.Book.AdminAdd,
						controller.Book.AdminUpd,
						controller.Book.AdminDel,
						controller.Book.AdminDetail,
						controller.User.AdminList,
						controller.User.AdminAdd,
						controller.User.AdminUpd,
						controller.User.AdminDel,
						controller.User.AdminUserInfo,
						controller.User.AdminUpdPassword,
						controller.Tag.AdminList,
						controller.Tag.AdminAdd,
						controller.Tag.AdminUpd,
						controller.Tag.AdminDel,
						controller.Sort.AdminList,
						controller.Sort.AdminAdd,
						controller.Sort.AdminUpd,
						controller.Sort.AdminDel,
						controller.Language.AdminList,
						controller.Order.AdminList,
						controller.Order.AdminPostpone,
						controller.Order.AdminReturn,
						controller.Order.AdminDeliver,
						controller.Notify.AdminAdd,
						controller.Notify.AdminUpd,
						controller.Notify.AdminList,
						controller.Notify.AdminDel,
						controller.Log.AdminLoginList,
						controller.Log.AdminOpList,
						controller.SystemInfo.Info,
						controller.Comment.AdminList,
						controller.Comment.AdminDel,
						controller.Index.IndexData,
					)
				})
			})
			// 启动前台项目gtoken
			frontendToken, err := StartFrontendGToken()
			if err != nil {
				return err
			}
			//前台项目的路由
			s.Group("/v1/app", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				//不需要登录的路由组绑定
				group.Bind(
					controller.User.Register, //用户注册
					controller.Sort.List,
					controller.Tag.List,
					controller.Book.List,
					controller.Book.Details,
					controller.Comment.List,
					controller.Comment.Like,
					controller.Book.Recommended,
					controller.Notify.List,
				)
				//需要登录的路由组绑定
				group.Group("/", func(group *ghttp.RouterGroup) {
					//验证token
					err := frontendToken.Middleware(ctx, group)
					if err != nil {
						return
					}
					group.Bind(
						controller.User.Info,
						controller.QiNiu.GetToken,
						controller.User.UpdInfo,
						controller.User.UpdPassword,
						controller.User.UpdPushSwitch,
						controller.Address,
						controller.Like.Like,
						controller.Collect.Collect,
						controller.Comment.Add,
						controller.Order.Preload,
						controller.Order.Add,
						controller.User.Points,
						controller.Order.Details,
						controller.Order.Pay,
						controller.Collect.List,
						controller.Collect.Del,
						controller.Like.List,
						controller.Like.Del,
						controller.Comment.MyList,
						controller.Order.List,
						controller.Order.Return,
					)
				})
			})

			s.Run()
			return nil
		},
	}
)
