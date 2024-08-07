package index

import (
	"context"
	"errors"
	"github.com/duke-git/lancet/v2/datetime"
	"goFrameMyServer/api/v1/admin"
	"goFrameMyServer/internal/dao"
	"goFrameMyServer/internal/service"
	"sort"
	"strings"
	"time"
)

type sIndex struct {
}

func init() {
	service.RegisterIndex(New())
}

func New() *sIndex {
	return &sIndex{}
}

func (a *sIndex) IndexData(ctx context.Context, req *admin.IndexReq) (res *admin.IndexRes, err error) {
	res = new(admin.IndexRes)
	res.BookTotal, err = dao.Book.Ctx(ctx).Count()
	if err != nil {
		return nil, errors.New("查询错误")
	}
	res.BookWeeks, err = dao.Book.Ctx(ctx).Where("created_at >= ? and created_at <= ?", datetime.FormatTimeToStr(datetime.BeginOfWeek(time.Now()), "yyyy-mm-dd hh:mm:ss"), datetime.FormatTimeToStr(datetime.EndOfWeek(time.Now()), "yyyy-mm-dd hh:mm:ss")).Count()
	if err != nil {
		return nil, errors.New("查询错误")
	}
	res.BorrowBook, err = dao.Order.Ctx(ctx).Where("status = ?", 6).Count()
	if err != nil {
		return nil, errors.New("查询错误")
	}
	res.BorrowPeople, err = dao.Order.Ctx(ctx).Where("status = ?", 6).Group("user_id").Count()
	if err != nil {
		return nil, errors.New("查询错误")
	}
	res.AlsoBook, err = dao.Order.Ctx(ctx).Where("status = ?", 3).Count()
	if err != nil {
		return nil, errors.New("查询错误")
	}
	res.AlsoPeople, err = dao.Order.Ctx(ctx).Where("status = ?", 3).Group("user_id").Count()
	if err != nil {
		return nil, errors.New("查询错误")
	}
	res.OverdueBook, err = dao.Order.Ctx(ctx).Where("status = ? and return_at < ?", 3, datetime.GetNowDateTime()).Count()
	if err != nil {
		return nil, errors.New("查询错误")
	}
	err = dao.Interview.Ctx(ctx).Limit(7).Order("id desc").Scan(&res.LatelyLine)
	if err != nil {
		return nil, errors.New("查询错误")
	}
	sort.Slice(res.LatelyLine, func(i, j int) bool {
		return i > j
	})
	for i := 0; i < len(res.LatelyLine); i++ {
		res.LatelyLine[i].T = strings.Split(res.LatelyLine[i].T, " ")[0]
	}
	err = dao.Book.Ctx(ctx).Limit(15).Order("borrow_num desc").Fields("id,title,borrow_num").Scan(&res.BookData)
	if err != nil {
		return nil, errors.New("查询错误")
	}
	err = dao.SortData.Ctx(ctx).As("sd").LeftJoin("sort as s", "sd.sort_id = s.id").Fields("sd.id,sd.sort_id,s.title,sd.num").
		Limit(10).Order("sd.num desc").Scan(&res.SortData)
	if err != nil {
		return nil, errors.New("查询错误")
	}
	return
}
