package user

import (
	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"
	"calligraphy/apps/user/rpc/types/user"
	"calligraphy/common/app_redis"
	"context"
	"fmt"

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

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (*types.UserRegisterResponse, error) {
	// 进行验证码验证
	v, err := app_redis.Redis.GetCtx(l.ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if v != req.EmailCode {
		return nil, fmt.Errorf("验证码不匹配")
	}

	// 删除 Redis 中的验证码信息
	_, err = app_redis.Redis.DelCtx(l.ctx, req.Email)
	if err != nil {
		return nil, err
	}

	// 调用 RPC 注册用户
	_, err = l.svcCtx.UserRpc.Register(l.ctx, &user.UserRegisterRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	// 注册成功，返回空响应
	return &types.UserRegisterResponse{}, nil
}
