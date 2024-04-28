package logic

import (
	"calligraphy/apps/user/model"
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

	res, err := (&model.User{}).FindOne(l.svcCtx.DB, uint(in.Id))

	if err != nil {
		l.Error("rpc 用户修改头像的时候出现了问题，或许没有查到，获取数据库操作出现了问题， err", err.Error())
		return nil, err
	}
	fmt.Println(res.AvatarBackground)
	res.AvatarBackground = in.Url

	err = (&model.User{}).UpdateUser(l.svcCtx.DB, uint(in.Id), res)
	if err != nil {
		l.Error("rpc 用户修改头像的时候数据库操作出现了问题， err", err.Error())
		return nil, err
	}
	return &user.UserModAvatarResponse{}, nil
}
