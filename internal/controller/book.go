package controller

import (
	"context"
	"goFrameMyServer/api/v1/admin"
	v1 "goFrameMyServer/api/v1/app"
	"goFrameMyServer/internal/service"
)

var Book = cBook{}

type cBook struct {
}

func (a *cBook) List(ctx context.Context, req *v1.BookListReq) (res *v1.BookListRes, err error) {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Size = 20
	}
	res, err = service.Book().List(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *cBook) Details(ctx context.Context, req *v1.BookDetailsReq) (res *v1.BookDetailsRes, err error) {
	res, err = service.Book().Details(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}

func (a *cBook) Recommended(ctx context.Context, req *v1.BookRecommendedListReq) (res *v1.BookRecommendedListRes, err error) {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Size = 5
	}
	res, err = service.Book().Recommended(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *cBook) AdminList(ctx context.Context, req *admin.BookListReq) (res *admin.BookListRes, err error) {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Size = 20
	}
	res, err = service.Book().AdminList(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *cBook) AdminAdd(ctx context.Context, req *admin.BookAddReq) (res *admin.BookAddRes, err error) {
	res, err = service.Book().AdminAdd(ctx, req)
	if err != nil {
		return nil, err
	}
	d := new(admin.BookAddRes)
	d.Successful = true
	return
}

func (a *cBook) AdminUpd(ctx context.Context, req *admin.BookUpdReq) (res *admin.BookUpdRes, err error) {
	res, err = service.Book().AdminUpd(ctx, req)
	if err != nil {
		return nil, err
	}
	d := new(admin.BookUpdRes)
	d.Successful = true
	return
}

func (a *cBook) AdminDel(ctx context.Context, req *admin.BookDelReq) (res *admin.BookDelRes, err error) {
	res, err = service.Book().AdminDel(ctx, req)
	if err != nil {
		return nil, err
	}
	d := new(admin.BookDelRes)
	d.Successful = true
	return
}

func (a *cBook) AdminDetail(ctx context.Context, req *admin.BookDetailReq) (res *admin.BookDetailRes, err error) {
	res, err = service.Book().AdminDetail(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}
