package order

import (
	"context"
	"errors"
	"github.com/duke-git/lancet/v2/datetime"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"goFrameMyServer/api/v1"
	"goFrameMyServer/api/v1/admin"
	v1 "goFrameMyServer/api/v1/app"
	"goFrameMyServer/internal/consts"
	"goFrameMyServer/internal/dao"
	"goFrameMyServer/internal/model"
	"goFrameMyServer/internal/model/entity"
	"goFrameMyServer/internal/service"
	"time"
)

type sOrder struct {
}

func init() {
	service.RegisterOrder(New())
}

func New() *sOrder {
	return &sOrder{}
}

func (s *sOrder) Preload(ctx context.Context, req *v1.OrderPreloadReq) (res *v1.OrderPreloadRes, err error) {
	sql := dao.Book.Ctx(ctx).WherePri(req.BookId)
	if req.Way == 1 {
		sql = sql.Fields("id as book_id,cover,title,price")
	} else {
		sql = sql.Fields("id as book_id,cover,title,buy_price as price")
	}
	err = sql.Scan(&res)
	if err != nil {
		return nil, errors.New("查询错误")
	}
	//res.Price = res.Price / 100
	return
}

func (s *sOrder) Add(ctx context.Context, req *v1.OrderAddReq) (res *v1.OrderAddRes, err error) {
	var address *entity.Address
	var book *entity.Book
	var order *entity.Order
	userId := gconv.Uint(ctx.Value(consts.CtxUserId))

	err = dao.Address.Ctx(ctx).WherePri(req.AddressId).Scan(&address)
	if err != nil {
		return nil, errors.New("创建订单失败")
	}
	err = dao.Order.Ctx(ctx).
		Where("user_id = ? and book_id = ? and way = ? and address_name = ? and address_phone = ? and address = ? and status = 0", userId, req.BookId, req.Way, address.ReceiverName, address.ReceiverPhone, address.AddressContent).
		Scan(&order)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	if order != nil {
		res = new(v1.OrderAddRes)
		res.Id = int64(order.Id)
		return
	}
	err = dao.Book.Ctx(ctx).WherePri(req.BookId).Scan(&book)
	if err != nil {
		return nil, errors.New("创建订单失败")
	}

	data := &entity.Order{
		UserId:       userId,
		AddressName:  address.ReceiverName,
		AddressPhone: address.ReceiverPhone,
		Address:      address.AddressContent,
		BookId:       req.BookId,
		Way:          req.Way,
		Status:       0,
	}

	if req.Way == 1 {
		data.Price = uint(book.Price)
		data.LendAt = gtime.New(datetime.GetNowDate())
		//为什么是+37天 因为7天无理由退货(/滑稽)！
		data.ReturnAt = gtime.New(datetime.AddDay(time.Now(), 37))

	} else {
		data.Price = uint(book.BuyPrice)
	}
	id, err := dao.Order.Ctx(ctx).InsertAndGetId(data)
	if err != nil {
		return nil, errors.New("创建订单失败")
	}
	res = new(v1.OrderAddRes)
	res.Id = id
	return
}

func (s *sOrder) Details(ctx context.Context, req *v1.OrderDetailsReq) (res *v1.OrderDetailsRes, err error) {
	userId := gconv.Uint(ctx.Value(consts.CtxUserId))
	err = dao.Order.Ctx(ctx).Where("user_id = ? and id = ?", userId, req.OrderId).Scan(&res)
	if err != nil {
		return nil, errors.New("查询错误")
	}
	return
}

func (s *sOrder) Pay(ctx context.Context, req *v1.OrderPayReq) (res *v1.OrderPayRes, err error) {
	var user *entity.User
	var order *entity.Order
	userId := gconv.Uint(ctx.Value(consts.CtxUserId))
	err = dao.User.Ctx(ctx).WherePri(userId).Fields("id,points").Scan(&user)
	if err != nil {
		return nil, errors.New("支付失败")
	}
	err = dao.Order.Ctx(ctx).WherePri(req.OrderId).Where("status = 0").Scan(&order)
	if err != nil {
		return nil, errors.New("支付失败")
	}
	if order == nil {
		return nil, errors.New("未查找到订单,可能已经支付。")
	}
	if user.Points < order.Price {
		return nil, errors.New("积分余额不足,请联系管理员充值积分")
	}
	err = dao.User.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.User.Ctx(ctx).TX(tx).WherePri(userId).Decrement("points", order.Price)
		if err != nil {
			return errors.New("扣除积分失败")
		}
		_, err = dao.Order.Ctx(ctx).TX(tx).WherePri(req.OrderId).Data(g.Map{"status": 1}).Update()
		if err != nil {
			return errors.New("扣除积分失败")
		}
		_, err = dao.Book.Ctx(ctx).TX(tx).WherePri(order.BookId).Increment("borrow_num", 1)
		if err != nil {
			return errors.New("扣除积分失败")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return
}

func (s *sOrder) AdminList(ctx context.Context, req *admin.OrderListReq) (res *admin.OrderListRes, err error) {
	var list []*model.OrderAdminList

	var m = dao.Order.Ctx(ctx).As("o")
	if req.Status != -1 {
		m = m.Where("o.status = ?", req.Status)
	}
	res = &admin.OrderListRes{
		CommonPaginationRes: api.CommonPaginationRes{
			Page: req.Page,
			Size: req.Size,
			List: []*model.OrderAdminList{},
		},
	}

	if res.Total, err = m.Count(); err != nil {
		return res, err
	}
	if res.Total == 0 {
		return res, nil
	}
	m = m.Page(req.Page, req.Size)
	err = m.LeftJoin("user as u", "o.user_id = u.id").LeftJoin("book as b", "b.id = o.book_id").
		Fields("o.id as id,o.user_id,u.nickname,o.address_name,o.address_phone,o.book_id," +
			"b.title as book_title,o.price,o.way,o.status,o.lend_at,o.return_at").Order("o.id desc").Scan(&list)
	res.List = list
	return
}

func (s *sOrder) AdminUpdStatus(ctx context.Context, orderId, status int) (err error) {

	sql := dao.Order.Ctx(ctx).WherePri(orderId).Data(g.Map{"status": status})
	if status == 3 {
		sql = sql.Where("status = 6")
	}
	if status == 2 {
		sql = sql.Where("status = 1")
	}
	_, err = sql.Update()
	if err != nil {
		return errors.New("修改状态失败")
	}
	return
}

func (s *sOrder) AdminPostpone(ctx context.Context, req *admin.OrderPostponeReq) (res *admin.OrderPostponeRes, err error) {
	var order *entity.Order
	err = dao.Order.Ctx(ctx).WherePri(req.OrderId).Where("status = 6").Scan(&order)
	if err != nil {
		return nil, errors.New("订单查询错误")
	}
	if order == nil {
		return nil, errors.New("查询错误")
	}
	_, err = dao.Order.Ctx(ctx).WherePri(req.OrderId).Data(g.Map{"return_at": datetime.AddDay(order.ReturnAt.Time, 30)}).Update()
	if err != nil {
		return nil, errors.New("添加时间错误")
	}
	return
}

func (s *sOrder) List(ctx context.Context, req *v1.OrderListReq) (res *v1.OrderListRes, err error) {
	var list []*model.OrderList

	var m = dao.Order.Ctx(ctx).As("o")
	if req.Status != -1 {
		m = m.Where("o.status = ?", req.Status)
	}
	res = &v1.OrderListRes{
		CommonPaginationRes: api.CommonPaginationRes{
			Page: req.Page,
			Size: req.Size,
			List: []*model.OrderList{},
		},
	}

	if res.Total, err = m.Count(); err != nil {
		return res, err
	}
	if res.Total == 0 {
		return res, nil
	}
	m = m.Page(req.Page, req.Size)
	err = m.LeftJoin("book as b", "b.id = o.book_id").
		Fields("o.id as id,o.book_                                          id," +
			"b.title as book_title,b.cover,o.price,o.way,o.status,o.lend_at,o.return_at").Order("o.id desc").Scan(&list)
	res.List = list
	return
}
