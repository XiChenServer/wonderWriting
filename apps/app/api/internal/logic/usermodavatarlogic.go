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

type UserModAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserModAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserModAvatarLogic {
	return &UserModAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserModAvatarLogic) UserModAvatar(r *http.Request) (resp *types.UserModAvatarResponse, err error) {
	// todo: add your logic here and delete this line
	//从jwt获取id
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()

	//图片上传到七牛云
	file, handler, err := r.FormFile("image")
	if err != nil {
		fmt.Println(err)
		return &types.UserModAvatarResponse{}, err
	}
	defer file.Close()
	url, err := qiniu.UploadToQiNiu(file, handler.Size, "AvatarBackground/", handler.Filename)
	if err != nil {
		fmt.Println(err.Error())
		return &types.UserModAvatarResponse{}, err
	}
	//调用grpc中的服务
	_, err = l.svcCtx.UserRpc.UserModAvatar(l.ctx, &user.UserModAvatarRequest{
		Id:  uid,
		Url: url,
	})
	if err != nil {
		return &types.UserModAvatarResponse{}, err
	}
	return &types.UserModAvatarResponse{}, nil
}
