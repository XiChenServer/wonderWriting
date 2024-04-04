package logic

import (
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
	res, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	res.BackgroundImage.String = in.Url
	err = l.svcCtx.UserModel.Update(l.ctx, res)
	if err != nil {
		return nil, err
	}
	return &user.UserModBackgroundResponse{}, nil
}
