package logic

import (
	"calligraphy/apps/user/model"
	"context"

	"calligraphy/apps/user/rpc/internal/svc"
	"calligraphy/apps/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserModInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserModInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserModInfoLogic {
	return &UserModInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserModInfoLogic) UserModInfo(in *user.UserModInfoRequest) (*user.UserModInfoResponse, error) {
	// todo: add your logic here and delete this line

	//查询
	res, err := (&model.User{}).FindOne(l.svcCtx.DB, uint(in.Id))
	if err != nil {
		return nil, err
	}
	res.Nickname = in.NickName
	res.Phone = in.Phone

	//进行修改
	err = (&model.User{}).UpdateUser(l.svcCtx.DB, uint(in.Id), res)
	if err != nil {
		return nil, err
	}
	return &user.UserModInfoResponse{}, nil
}
