package user

import (
	"calligraphy/apps/user/rpc/types/user"
	"calligraphy/common/app_redis"
	"calligraphy/common/jwtx"
	"context"
	"fmt"

	"time"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserForgetPwdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserForgetPwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserForgetPwdLogic {
	return &UserForgetPwdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserForgetPwdLogic) UserForgetPwd(req *types.UserForgetPwdRequest) (*types.UserForgetPwdResponse, error) {
	// 调用rpc进行获取信息
	res, err := l.svcCtx.UserRpc.UserForgetPwd(l.ctx, &user.UserForgetPwdRequest{Email: req.Email})
	if err != nil {
		return nil, err
	}

	// 检查验证码是否匹配
	exists, err := app_redis.Redis.ExistsCtx(l.ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("验证码不存在或已过期")
	}

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

	// 生成 JWT
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	accessToken, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, res.Id)
	if err != nil {
		return nil, err
	}

	return &types.UserForgetPwdResponse{
		AccessToken:  accessToken,
		AccessExpire: accessExpire,
	}, nil
}
