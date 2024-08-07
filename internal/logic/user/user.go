package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/duke-git/lancet/v2/validator"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"goFrameMyServer/api/v1"
	"goFrameMyServer/api/v1/admin"
	v1 "goFrameMyServer/api/v1/app"
	"goFrameMyServer/internal/consts"
	"goFrameMyServer/internal/dao"
	"goFrameMyServer/internal/model"
	"goFrameMyServer/internal/model/entity"
	"goFrameMyServer/internal/service"
	"goFrameMyServer/utility"
)

type sUser struct {
}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

func (u *sUser) Register(ctx context.Context, req *model.RegisterInput) (out *model.RegisterOutput, err error) {
	//验证是否是有效的邮箱地址
	ok := validator.IsEmail(req.UserName)
	if !ok {
		return nil, errors.New("用户名非邮件,不可注册")
	}
	//验证数据库是否已存在用户
	count, err := dao.User.Ctx(ctx).Where("user_name = ?", req.UserName).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("用户名已被注册")
	}
	//处理加密盐和密码的逻辑
	UserSalt := grand.S(10)
	//加密密码 MD5加密
	password := utility.EncryptPassword(req.Password, UserSalt)
	user := &entity.User{
		UserName: req.UserName,
		Role:     "user",
		Password: password,
		Email:    req.UserName,
		UserSalt: UserSalt,
		Nickname: req.UserName,
	}
	//插入数据库
	userId, err := dao.User.Ctx(ctx).Data(user).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &model.RegisterOutput{Id: uint(userId)}, nil
}

func (u *sUser) Info(ctx context.Context) (out *v1.UserInfoRes, err error) {
	userId := gconv.Uint(ctx.Value(consts.CtxUserId))
	err = dao.User.Ctx(ctx).Where("id = ? and status = 0", userId).Scan(&out)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	return out, nil
}
func (u *sUser) UpdInfo(ctx context.Context, req *v1.UserUpdInfoReq) error {
	userId := gconv.Uint(ctx.Value(consts.CtxUserId))
	_, err := dao.User.Ctx(ctx).WherePri(userId).Data(g.Map{"avatar": req.Avatar, "nickname": req.Nickname,
		"description": req.Description, "mobile": req.Mobile, "email": req.Email}).Update()
	if err != nil {
		return errors.New("头像更换失败")
	}
	return nil
}

func (u *sUser) UpdPassword(ctx context.Context, req *v1.UserUpdPasswordReq) (err error) {
	var user *entity.User
	//从缓存中国获取userId
	userId := gconv.Uint(ctx.Value(consts.CtxUserId))
	//查询用户信息
	err = dao.User.Ctx(ctx).WherePri(userId).Scan(&user)
	if err != nil {
		return errors.New("查询用户失败")
	}
	//加密密码 MD5加密 进行比对是否一致
	password := utility.EncryptPassword(req.Password, user.UserSalt)
	if user.Password != password {
		return errors.New("密码错误，无法修改")
	}
	//加密新密码
	newPassword := utility.EncryptPassword(req.NewPassword, user.UserSalt)
	//修改新密码到数据库
	_, err = dao.User.Ctx(ctx).WherePri(userId).Data(g.Map{"password": newPassword}).Update()
	if err != nil {
		return errors.New("修改密码失败")
	}
	//删除token 使token过期
	_, err = g.Redis().Del(ctx, fmt.Sprintf("GToken:User:%d", userId))
	if err != nil {
		return errors.New("删除Token失败")
	}
	return nil
}

func (u *sUser) UpdUserPushStatus(ctx context.Context) (err error) {
	var user *entity.User
	var status int
	//从缓存中国获取userId
	userId := gconv.Uint(ctx.Value(consts.CtxUserId))
	//查询用户信息
	err = dao.User.Ctx(ctx).WherePri(userId).Scan(&user)
	if err != nil {
		return errors.New("查询用户失败")
	}
	if user.PushSwitch == 0 {
		status = 1
	} else {
		status = 0
	}
	_, err = dao.User.Ctx(ctx).WherePri(userId).Data(g.Map{"push_switch": status}).Update()
	if err != nil {
		return errors.New("修改失败")
	}
	return nil
}

func (u *sUser) Points(ctx context.Context, req *v1.UserPointReq) (res *v1.UserPointRes, err error) {
	//从缓存中国获取userId
	userId := gconv.Uint(ctx.Value(consts.CtxUserId))
	err = dao.User.Ctx(ctx).WherePri(userId).Fields("points").Scan(&res)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	return
}

func (u *sUser) AdminUserList(ctx context.Context, req *admin.UserListReq) (res *admin.UserListRes, err error) {
	var user []*model.AdminUserList
	var m = dao.User.Ctx(ctx)
	if req.Name != "" {
		m = m.WhereLike("nickname", "%"+req.Name+"%")
	}
	res = &admin.UserListRes{
		CommonPaginationRes: api.CommonPaginationRes{
			Page: req.Page,
			Size: req.Size,
			List: []*model.AdminUserList{},
		},
	}
	listModel := m.Page(req.Page, req.Size)
	if res.Total, err = listModel.Count(); err != nil {
		return res, err
	}
	if res.Total == 0 {
		return res, nil
	}
	if err = listModel.Scan(&user); err != nil {
		return res, err
	}
	if len(user) == 0 {
		return res, err
	}
	res.List = user
	return
}

func (u *sUser) AdminUserAdd(ctx context.Context, req *admin.UserAddReq) (res *admin.UserAddRes, err error) {
	if req.Email != "" {
		//验证是否是有效的邮箱地址
		ok := validator.IsEmail(req.Email)
		if !ok {
			return nil, errors.New("邮箱验证失败")
		}
	}
	if req.Role == "user" || req.Role == "admin" {
		//验证数据库是否已存在用户
		count, err := dao.User.Ctx(ctx).Where("user_name = ?", req.UserName).Count()
		if err != nil {
			return nil, err
		}
		if count > 0 {
			return nil, errors.New("用户名已被注册")
		}
		//处理加密盐和密码的逻辑
		UserSalt := grand.S(10)
		//加密密码 MD5加密
		password := utility.EncryptPassword(req.Password, UserSalt)
		user := &entity.User{
			UserName: req.UserName,
			Role:     req.Role,
			Password: password,
			Email:    req.Email,
			UserSalt: UserSalt,
			Nickname: req.Nickname,
			Mobile:   req.Mobile,
		}
		//插入数据库
		_, err = dao.User.Ctx(ctx).Data(user).Insert()
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("用户角色错误,请认真填写")
	}
	return
}

func (u *sUser) AdminUserUpd(ctx context.Context, req *admin.UserUpdReq) (res *admin.UserUpdRes, err error) {
	count, err := dao.User.Ctx(ctx).WherePri(req.Id).Count()
	if count <= 0 {
		return nil, errors.New("未查询到此用户")
	}
	if req.Role == "user" || req.Role == "admin" {
		_, err = dao.User.Ctx(ctx).WherePri(req.Id).Data(entity.User{
			Role:     req.Role,
			Status:   uint(req.Status),
			Nickname: req.Nickname,
			Email:    req.Email,
			Mobile:   req.Mobile,
		}).OmitEmpty().Update()
		if err != nil {
			return nil, errors.New("修改用户信息失败")
		}
	} else {
		return nil, errors.New("用户角色设置错误")
	}
	return
}

func (u *sUser) AdminUserDel(ctx context.Context, req *admin.UserDelReq) (res *admin.UserDelRes, err error) {
	_, err = dao.User.Ctx(ctx).WhereIn("id", req.Id).Delete()
	if err != nil {
		return nil, errors.New("删除错误")
	}
	return
}

func (u *sUser) AdminUserUpdPassword(ctx context.Context, req *admin.UserUpdPasswordReq) (res *admin.UserUpdPasswordRes, err error) {
	var user *entity.User
	err = dao.User.Ctx(ctx).WherePri(req.Id).Scan(&user)
	if err != nil {
		return nil, errors.New("查询用户失败")
	}
	if user == nil {
		return nil, errors.New("未查询到此用户")
	}
	//加密新密码
	newPassword := utility.EncryptPassword(req.Password, user.UserSalt)
	//修改新密码到数据库
	_, err = dao.User.Ctx(ctx).WherePri(req.Id).Data(g.Map{"password": newPassword}).Update()
	if err != nil {
		return nil, errors.New("修改密码失败")
	}
	//删除token 使token过期
	_, err = g.Redis().Del(ctx, fmt.Sprintf("GToken:User:%d", req.Id))
	if err != nil {
		return nil, errors.New("删除Token失败")
	}
	return
}

func (u *sUser) AdminUserGetInfo(ctx context.Context, req *admin.UserGetInfoReq) (res *admin.UserGetInfoRes, err error) {
	err = dao.User.Ctx(ctx).WherePri(req.Id).Scan(&res)
	if err != nil {
		return nil, errors.New("获取详情失败")
	}
	if res == nil {
		return nil, errors.New("未查询到此用户")
	}
	res.Password = ""
	return
}
