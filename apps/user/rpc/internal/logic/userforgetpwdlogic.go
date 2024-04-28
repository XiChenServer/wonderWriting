package logic

import (
	"calligraphy/apps/user/model"
	"calligraphy/apps/user/rpc/internal/svc"
	"calligraphy/apps/user/rpc/types/user"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserForgetPwdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserForgetPwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserForgetPwdLogic {
	return &UserForgetPwdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserForgetPwdLogic) UserForgetPwd(in *user.UserForgetPwdRequest) (*user.UserForgetPwdResponse, error) {
	// todo: add your logic here and delete this line

	res, err := (&model.User{}).FindOneByEmail(l.svcCtx.DB, in.Email)
	if err != nil {
		l.Error("rpc 用户查找邮箱出现问题， err", err.Error())
		return nil, err
	}
	l.Info("rpc 用户忘记密码成功，有该邮箱，email", in.Email)
	return &user.UserForgetPwdResponse{
		Id: int64(res.UserID),
	}, nil
}
