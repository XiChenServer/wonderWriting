package logic

import (
	"calligraphy/apps/user/rpc/types/user"
	"context"
	"encoding/json"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	//从jwt中获取id
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	res, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: uid,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserInfoResponse{
		Id:               res.Id,
		NickName:         res.NickName,
		Account:          res.Account,
		Email:            res.Email,
		AvatarBackground: res.AvatarBackground,
		BackgroundImage:  res.BackgroundImage,
		Phone:            res.Phone,
	}, nil
}
