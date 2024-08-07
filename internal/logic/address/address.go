package address

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"goFrameMyServer/api/v1"
	v1 "goFrameMyServer/api/v1/app"
	"goFrameMyServer/internal/consts"
	"goFrameMyServer/internal/dao"
	"goFrameMyServer/internal/model"
	"goFrameMyServer/internal/model/entity"
	"goFrameMyServer/internal/service"
)

type sAddress struct {
}

func init() {
	service.RegisterAddress(New())
}

func New() *sAddress {
	return &sAddress{}
}

func (a *sAddress) Add(ctx context.Context, req *v1.AddressAddReq) (id int64, err error) {
	userId := gconv.Uint(ctx.Value(consts.CtxUserId))
	address := &entity.Address{
		UserId:         userId,
		ReceiverName:   req.ReceiverName,
		ReceiverPhone:  req.ReceiverPhone,
		AddressContent: req.AddressContent,
		Status:         uint(req.Status),
	}
	addressId, err := dao.Address.Ctx(ctx).Data(address).InsertAndGetId()
	if err != nil {
		return 0, errors.New("添加地址失败")
	}
	return addressId, nil
}
func (a *sAddress) Upd(ctx context.Context, req *v1.AddressUpdReq) (err error) {
	userId := gconv.Uint(ctx.Value(consts.CtxUserId))
	address := &entity.Address{
		ReceiverName:   req.ReceiverName,
		ReceiverPhone:  req.ReceiverPhone,
		AddressContent: req.AddressContent,
		Status:         uint(req.Status),
	}
	if address.Status == 2 {
		_, err = dao.Address.Ctx(ctx).Where("user_id = ?", userId).Data(g.Map{"status": 1}).Update()
		if err != nil {
			return errors.New("修改默认地址失败")
		}
	}
	_, err = dao.Address.Ctx(ctx).WherePri(req.Id).OmitEmptyData().Data(address).Update()
	if err != nil {
		return errors.New("修改地址失败")
	}
	return nil
}

func (a *sAddress) Del(ctx context.Context, id int64) (err error) {
	userId := gconv.Uint(ctx.Value(consts.CtxUserId))
	_, err = dao.Address.Ctx(ctx).Delete("user_id = ? and id = ?", userId, id)
	if err != nil {
		return errors.New("地址删除失败")
	}
	return nil
}

func (a *sAddress) List(ctx context.Context, req *v1.AddressListReq) (res *v1.AddressListRes, err error) {
	var address []*model.Address

	userId := gconv.Uint(ctx.Value(consts.CtxUserId))
	var m = dao.Address.Ctx(ctx).Where("user_id = ?", userId)
	res = &v1.AddressListRes{
		CommonPaginationRes: api.CommonPaginationRes{
			//Page: req.Page,
			//Size: req.Size,
			List: []*model.Address{},
		},
	}
	//listModel := m.Page(req.Page, req.Size)
	//if res.Total, err = listModel.Count(); err != nil {
	//	return res, err
	//}
	//if res.Total == 0 {
	//	return res, nil
	//}
	if err = m.Order("status desc,id desc").Scan(&address); err != nil {
		return res, err
	}
	if len(address) == 0 {
		return res, err
	}
	res.List = address
	return res, nil
}
