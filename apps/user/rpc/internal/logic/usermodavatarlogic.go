package logic

import (
	"context"
	"fmt"

	"calligraphy/apps/user/rpc/internal/svc"
	"calligraphy/apps/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserModAvatarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserModAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserModAvatarLogic {
	return &UserModAvatarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserModAvatarLogic) UserModAvatar(in *user.UserModAvatarRequest) (*user.UserModAvatarResponse, error) {
	// todo: add your logic here and delete this line

	res, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)

	if err != nil {
		return nil, err
	}
	fmt.Println(res.AvatarBackground.String)
	res.AvatarBackground.String = in.Url
	res.AvatarBackground.Valid = true
	fmt.Println(res.AvatarBackground.String)
	err = l.svcCtx.UserModel.Update(l.ctx, res)
	if err != nil {
		return nil, err
	}
	return &user.UserModAvatarResponse{}, nil
}
