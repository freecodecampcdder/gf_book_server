package controller

import (
	"context"
	"goFrameMyServer/api/v1/admin"
	v1 "goFrameMyServer/api/v1/app"
	"goFrameMyServer/internal/model"
	"goFrameMyServer/internal/service"
)

var User = cUser{}

type cUser struct {
}

func (u *cUser) Register(ctx context.Context, req *v1.UserRegisterReq) (res *v1.UserRegisterRes, err error) {
	out, err := service.User().Register(ctx, &model.RegisterInput{UserName: req.UserName, Password: req.Password})
	if err != nil {
		return nil, err
	}
	res = &v1.UserRegisterRes{
		UserId: out.Id,
	}
	return res, nil
}

func (u *cUser) Info(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	res, err = service.User().Info(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *cUser) UpdInfo(ctx context.Context, req *v1.UserUpdInfoReq) (res *v1.UserUpdAvatarRes, err error) {
	err = service.User().UpdInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	d := new(v1.UserUpdAvatarRes)
	d.Successful = true
	return d, nil
}

func (u *cUser) UpdPassword(ctx context.Context, req *v1.UserUpdPasswordReq) (res *v1.UserUpdPasswordRes, err error) {
	err = service.User().UpdPassword(ctx, req)
	if err != nil {
		return nil, err
	}
	d := new(v1.UserUpdPasswordRes)
	d.Successful = true
	return d, nil
}

func (u *cUser) UpdPushSwitch(ctx context.Context, req *v1.UserPushSwitchReq) (res *v1.UserPushSwitchRes, err error) {
	err = service.User().UpdUserPushStatus(ctx)
	if err != nil {
		return nil, err
	}
	d := new(v1.UserPushSwitchRes)
	d.Successful = true
	return d, nil
}

func (u *cUser) Points(ctx context.Context, req *v1.UserPointReq) (res *v1.UserPointRes, err error) {
	res, err = service.User().Points(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}

func (u *cUser) AdminList(ctx context.Context, req *admin.UserListReq) (res *admin.UserListRes, err error) {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Size = 20
	}
	res, err = service.User().AdminUserList(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}
func (u *cUser) AdminAdd(ctx context.Context, req *admin.UserAddReq) (res *admin.UserAddRes, err error) {
	res, err = service.User().AdminUserAdd(ctx, req)
	if err != nil {
		return nil, err
	}
	d := new(admin.UserAddRes)
	d.Successful = true
	return
}

func (u *cUser) AdminUpd(ctx context.Context, req *admin.UserUpdReq) (res *admin.UserUpdRes, err error) {
	res, err = service.User().AdminUserUpd(ctx, req)
	if err != nil {
		return nil, err
	}
	d := new(admin.UserUpdRes)
	d.Successful = true
	return
}

func (u *cUser) AdminDel(ctx context.Context, req *admin.UserDelReq) (res *admin.UserDelRes, err error) {
	res, err = service.User().AdminUserDel(ctx, req)
	if err != nil {
		return nil, err
	}
	d := new(admin.UserDelRes)
	d.Successful = true
	return
}

func (u *cUser) AdminUpdPassword(ctx context.Context, req *admin.UserUpdPasswordReq) (res *admin.UserUpdPasswordRes, err error) {
	res, err = service.User().AdminUserUpdPassword(ctx, req)
	if err != nil {
		return nil, err
	}
	d := new(admin.UserUpdPasswordRes)
	d.Successful = true
	return
}

func (u *cUser) AdminUserInfo(ctx context.Context, req *admin.UserGetInfoReq) (res *admin.UserGetInfoRes, err error) {
	res, err = service.User().AdminUserGetInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}
