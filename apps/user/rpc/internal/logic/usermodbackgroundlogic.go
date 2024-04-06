package logic

import (
	"calligraphy/apps/user/model"
	"context"

	"calligraphy/apps/user/rpc/internal/svc"
	"calligraphy/apps/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserModBackgroundLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserModBackgroundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserModBackgroundLogic {
	return &UserModBackgroundLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserModBackgroundLogic) UserModBackground(in *user.UserModBackgroundRequest) (*user.UserModBackgroundResponse, error) {
	// todo: add your logic here and delete this line
	res, err := (&model.User{}).FindOne(l.svcCtx.DB, uint(in.Id))
	if err != nil {
		return nil, err
	}
	res.BackgroundImage = in.Url
	err = (&model.User{}).UpdateUser(l.svcCtx.DB, uint(in.Id), res)
	if err != nil {
		return nil, err
	}
	return &user.UserModBackgroundResponse{}, nil
}
