package qiniu

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
	"goFrameMyServer/internal/consts"
	"goFrameMyServer/internal/service"
)

type sQiniu struct {
}

func init() {
	service.RegisterQiniu(New())
}

func New() *sQiniu {
	return &sQiniu{}
}

func (s *sQiniu) GetToken(ctx context.Context) (string, error) {

	data, err := g.Redis().Get(ctx, consts.QiNiuToken)

	if data.String() == "" {

		//获得七牛云配置文件参数
		bucket := g.Cfg().MustGet(ctx, "qiniu.bucket").String()
		accessKey := g.Cfg().MustGet(ctx, "qiniu.accessKey").String()
		secretKey := g.Cfg().MustGet(ctx, "qiniu.secretKey").String()
		// 简单上传凭证
		putPolicy := storage.PutPolicy{
			Scope: bucket,
		}

		mac := auth.New(accessKey, secretKey)
		upToken := putPolicy.UploadToken(mac)
		fmt.Println(upToken)

		// 设置上传凭证有效期
		putPolicy = storage.PutPolicy{
			Scope: bucket,
		}
		putPolicy.Expires = 7200 //示例2小时有效期

		upToken = putPolicy.UploadToken(mac)
		err = g.Redis().SetEX(ctx, consts.QiNiuToken, upToken, consts.QiNiuTokenTime)

		if err != nil {
			return "", errors.New("存redis失败")
		}
		return upToken, nil
	} else {
		return data.String(), nil
	}
}
