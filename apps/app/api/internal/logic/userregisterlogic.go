package logic

import (
	"calligraphy/apps/user/rpc/types/user"
	"calligraphy/common/app_redis"
	"context"
	"fmt"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterResponse, err error) {
	// todo: add your logic here and delete this line
	v, err := app_redis.Redis.GetCtx(l.ctx, req.Email)
	if err != nil {
		return nil, err
	}
	fmt.Println("3")
	if v != req.EmailCode {
		return nil, err
	}
	fmt.Println("1")
	_, err = l.svcCtx.UserRpc.Register(l.ctx, &user.UserRegisterRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	fmt.Println("2")
	return nil, nil
}
