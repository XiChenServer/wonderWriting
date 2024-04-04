package logic

import (
	"calligraphy/apps/user/rpc/types/user"
	"calligraphy/pkg/qiniu"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserModBackgroundLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserModBackgroundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserModBackgroundLogic {
	return &UserModBackgroundLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserModBackgroundLogic) UserModBackground(r *http.Request) (resp *types.UserModBackgroundResponse, err error) {
	// todo: add your logic here and delete this line
	//从jwt获取id
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	res, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: uid,
	})
	if err != nil {
		return nil, err
	}

	// 图片上传到七牛云
	file, handler, err := r.FormFile("image")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()
	url, err := qiniu.UploadToQiNiu(file, handler.Size, "BackgroundImage/", handler.Filename)
	if err != nil {
		return nil, err
	}
	//调用grpc中的服务
	_, err = l.svcCtx.UserRpc.UserModBackground(l.ctx, &user.UserModBackgroundRequest{
		Id:  res.Id,
		Url: url,
	})
	return &types.UserModBackgroundResponse{}, nil

}
