package logic

import (
	"calligraphy/apps/user/model"
	"context"
	"fmt"

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
		fmt.Println("rpc UserModInfo 用户修改信息的时候没有找到该用户，或者数据库操作出现了问题，err", err.Error())
		return nil, err
	}
	res.Nickname = in.NickName
	res.Phone = in.Phone

	//进行修改
	err = (&model.User{}).UpdateUser(l.svcCtx.DB, uint(in.Id), res)
	if err != nil {
		fmt.Println("rpc UserModInfo 用户修改信息的时候数据库操作出现了问题，err", err.Error())
		return nil, err
	}
	return &user.UserModInfoResponse{}, nil
}
