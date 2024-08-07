package book

import (
	"context"
	"errors"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/gogf/gf/v2/database/gdb"
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

type sBook struct {
}

func init() {
	service.RegisterBook(New())
}

func New() *sBook {
	return &sBook{}
}

func (a *sBook) List(ctx context.Context, req *v1.BookListReq) (res *v1.BookListRes, err error) {
	var book []*model.BookList

	var m = dao.Book.Ctx(ctx).As("b")
	if req.Tag > 0 {
		m = m.LeftJoin("book_tag as bt", "b.id = bt.book_id")
		m = m.Where("bt.tag_id = ?", req.Tag)
	}
	if req.Sort > 0 {
		m = m.LeftJoin("book_sort as bs", "b.id = bs.book_id")
		m = m.Where("bs.sort_id = ?", req.Sort)
	}
	if req.Name != "" {
		m = m.WhereLike("b.title", "%"+req.Name+"%")
	}

	res = &v1.BookListRes{
		CommonPaginationRes: api.CommonPaginationRes{
			Page: req.Page,
			Size: req.Size,
			List: []*model.BookList{},
		},
	}
	listModel := m.Page(req.Page, req.Size)
	if res.Total, err = listModel.Count(); err != nil {
		return res, err
	}
	if res.Total == 0 {
		return res, nil
	}
	listModel = listModel.LeftJoin("language as l", "b.language_id = l.id")
	listModel = listModel.Fields("b.id as id,l.language as language,b.title as title,b.author,b.cover,b.translator,b.price,b.created_at")
	//1 最新 2 最热 3 推荐
	if req.Order == 1 || req.Order == 0 {
		listModel = listModel.Order("b.created_at desc")
	} else if req.Order == 2 {
		listModel = listModel.Order("b.wish_num desc")
	} else if req.Order == 3 {
		listModel = listModel.Order("b.recommended desc,id desc ")
	}

	if err = listModel.Scan(&book); err != nil {
		return res, err
	}
	if len(book) == 0 {
		return res, err
	}
	//for i := 0; i < len(book); i++ {
	//	price := book[i].Price / float64(100)
	//	book[i].Price = mathutil.RoundToFloat(price, 2)
	//}
	res.List = book
	return res, nil
}

func (a *sBook) Recommended(ctx context.Context, req *v1.BookRecommendedListReq) (res *v1.BookRecommendedListRes, err error) {
	var book []*model.BookRecommendedList

	var m = dao.Book.Ctx(ctx)
	res = &v1.BookRecommendedListRes{
		CommonPaginationRes: api.CommonPaginationRes{
			Page: req.Page,
			Size: req.Size,
			List: []*model.BookRecommendedList{},
		},
	}
	listModel := m.Page(req.Page, req.Size)
	if res.Total, err = listModel.Count(); err != nil {
		return res, err
	}
	if res.Total == 0 {
		return res, nil
	}

	listModel = listModel.Order("recommended desc,id desc ")
	if err = listModel.Scan(&book); err != nil {
		return res, err
	}
	if len(book) == 0 {
		return res, err
	}
	res.List = book
	return res, nil
}

func (a *sBook) Details(ctx context.Context, req *v1.BookDetailsReq) (res *v1.BookDetailsRes, err error) {
	res = new(v1.BookDetailsRes)
	err = dao.Book.Ctx(ctx).As("b").LeftJoin("language as l", "b.language_id = l.id").
		Fields("b.*,l.language").
		Where("b.id = ?", req.Id).Scan(&res)
	if err != nil {
		return nil, errors.New("查询详情失败")
	}
	err = dao.BookSort.Ctx(ctx).As("bs").
		LeftJoin("sort as s", "bs.sort_id = s.id").
		Fields("s.id,s.title,s.pid").
		Where("bs.book_id = ?", req.Id).
		Scan(&res.Sort)
	if err != nil {
		return nil, errors.New("查询分类失败")
	}
	err = dao.BookTag.Ctx(ctx).As("bt").
		LeftJoin("tag as t", "bt.tag_id = t.id").
		Fields("t.id,t.title").
		Where("bt.book_id = ?", req.Id).
		Scan(&res.Tag)
	if err != nil {
		return nil, errors.New("查询标签失败")
	}
	userId := req.UserId
	if userId > 0 {
		likeCount, err := dao.Like.Ctx(ctx).Where("user_id = ? and book_id = ?", userId, res.Id).Count()
		if err != nil {
			return nil, errors.New("查询错误")
		}
		if likeCount > 0 {
			res.IsWish = true
		}
		collectCount, err := dao.Collect.Ctx(ctx).Where("user_id = ? and book_id = ?", userId, res.Id).Count()
		if err != nil {
			return nil, errors.New("查询错误")
		}
		if collectCount > 0 {
			res.IsCollect = true
		}
	}
	//res.Price = mathutil.RoundToFloat(res.Price/float64(100), 2)
	//res.BuyPrice = mathutil.RoundToFloat(res.BuyPrice/float64(100), 2)
	return
}

func (a *sBook) AdminList(ctx context.Context, req *admin.BookListReq) (res *admin.BookListRes, err error) {
	var book []*model.AdminBookList

	var m = dao.Book.Ctx(ctx).As("b")
	if req.Name != "" {
		m = m.WhereLike("title", "%"+req.Name+"%")
	}
	res = &admin.BookListRes{
		CommonPaginationRes: api.CommonPaginationRes{
			Page: req.Page,
			Size: req.Size,
			List: []*model.AdminBookList{},
		},
	}
	listModel := m.Page(req.Page, req.Size)
	if res.Total, err = listModel.Count(); err != nil {
		return res, err
	}
	if res.Total == 0 {
		return res, nil
	}
	if err = listModel.Scan(&book); err != nil {
		return res, err
	}
	if len(book) == 0 {
		return res, err
	}
	for i := 0; i < len(book); i++ {
		//price := book[i].Price / float64(100)
		//book[i].Price = mathutil.RoundToFloat(price, 2)
		//BuyPrice := book[i].BuyPrice / float64(100)
		//book[i].BuyPrice = mathutil.RoundToFloat(BuyPrice, 2)
		in, _ := convertor.ToInt(book[i].PressTime)
		book[i].PressTime = time.Unix(in, 0).Format("2006-01")
	}
	res.List = book
	return
}

func (a *sBook) AdminAdd(ctx context.Context, req *admin.BookAddReq) (res *admin.BookAddRes, err error) {
	count, err := dao.Book.Ctx(ctx).Where("title = ?", req.Title).Count()
	if count > 0 {
		return nil, errors.New("请勿重复添加。")
	}
	userId := gconv.Uint(ctx.Value(consts.CtxUserId))
	err = dao.Book.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		bookId, err := dao.Book.Ctx(ctx).TX(tx).Data(entity.Book{
			LanguageId:   req.LanguageId,
			Title:        req.Title,
			Author:       req.Author,
			Cover:        req.Cover,
			Translator:   req.Translator,
			Description:  req.Description,
			Status:       req.Status,
			Isbn:         req.Isbn,
			Press:        req.Press,
			PressTime:    req.PressTime,
			PageNum:      req.PageNum,
			Price:        req.Price,
			BuyPrice:     req.BuyPrice,
			InventoryNum: req.InventoryNum,
			UserId:       int(userId),
			Recommended:  uint(req.Recommended),
		}).InsertAndGetId()
		if err != nil {
			return errors.New("插入错误")
		}
		var bookSortData []*entity.BookSort
		for _, v := range req.Sort {
			data := &entity.BookSort{
				BookId: uint(bookId),
				SortId: v,
			}
			bookSortData = append(bookSortData, data)
		}
		_, err = dao.BookSort.Ctx(ctx).TX(tx).Insert(bookSortData)
		if err != nil {
			return errors.New("插入分类失败")
		}
		var bookTagData []*entity.BookTag
		for _, v := range req.Tag {
			data := &entity.BookTag{
				BookId: uint(bookId),
				TagId:  v,
			}
			bookTagData = append(bookTagData, data)
		}
		if len(bookTagData) > 0 {
			_, err = dao.BookTag.Ctx(ctx).TX(tx).Insert(bookTagData)
			if err != nil {
				return errors.New("插入标签失败")
			}
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return
}

func (a *sBook) AdminUpd(ctx context.Context, req *admin.BookUpdReq) (res *admin.BookUpdRes, err error) {
	count, err := dao.Book.Ctx(ctx).WherePri(req.Id).Count()
	if count <= 0 {
		return nil, errors.New("未查询到图书,请输入正确的ID")
	}
	err = dao.Book.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.Book.Ctx(ctx).TX(tx).WherePri(req.Id).Data(&req).Update()
		if err != nil {
			return errors.New("修改图书失败")
		}
		_, err = dao.BookSort.Ctx(ctx).TX(tx).Where("book_id = ?", req.Id).Delete()
		if err != nil {
			return errors.New("删除分类失败")
		}
		var bookSortData []*entity.BookSort
		for _, v := range req.Sort {
			data := &entity.BookSort{
				BookId: uint(req.Id),
				SortId: v,
			}
			bookSortData = append(bookSortData, data)
		}
		_, err = dao.BookSort.Ctx(ctx).TX(tx).Insert(bookSortData)
		if err != nil {
			return errors.New("插入分类失败")
		}
		_, err = dao.BookTag.Ctx(ctx).TX(tx).Where("book_id = ?", req.Id).Delete()
		if err != nil {
			return errors.New("删除标签失败")
		}
		var bookTagData []*entity.BookTag
		for _, v := range req.Tag {
			data := &entity.BookTag{
				BookId: uint(req.Id),
				TagId:  v,
			}
			bookTagData = append(bookTagData, data)
		}
		if len(bookTagData) > 0 {
			_, err = dao.BookTag.Ctx(ctx).TX(tx).Insert(bookTagData)
			if err != nil {
				return errors.New("插入标签失败")
			}
		}
		return err
	})
	return
}

func (a *sBook) AdminDel(ctx context.Context, req *admin.BookDelReq) (res *admin.BookDelRes, err error) {
	_, err = dao.Book.Ctx(ctx).WhereIn("id", req.Id).Delete()
	if err != nil {
		return nil, errors.New("删除错误")
	}
	return
}

func (a *sBook) AdminDetail(ctx context.Context, req *admin.BookDetailReq) (res *admin.BookDetailRes, err error) {
	var tag []*entity.BookTag
	var sort []*entity.BookSort
	err = dao.Book.Ctx(ctx).WherePri(req.Id).Scan(&res)
	if err != nil {
		return nil, errors.New("查询错误")
	}
	if res == nil {
		return nil, errors.New("未查询到此书，请检查序号保证正确。")
	}

	err = dao.BookTag.Ctx(ctx).Where("book_id = ?", req.Id).Scan(&tag)
	if err != nil {
		return nil, errors.New("查询标签失败")
	}
	for _, v := range tag {
		res.Tag = append(res.Tag, v.TagId)
	}
	err = dao.BookSort.Ctx(ctx).Where("book_id = ?", req.Id).Scan(&sort)
	if err != nil {
		return nil, errors.New("查询分类失败")
	}
	for _, v := range sort {
		res.Sort = append(res.Sort, v.SortId)
	}
	return
}
