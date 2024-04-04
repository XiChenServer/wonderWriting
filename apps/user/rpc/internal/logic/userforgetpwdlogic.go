package logic

import (
	"context"
	"database/sql"

	"calligraphy/apps/user/rpc/internal/svc"
	"calligraphy/apps/user/rpc/types/user"

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
	email := sql.NullString{
		String: in.Email,
		Valid:  true,
	}
	res, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, email)
	if err != nil {
		return nil, err
	}
	return &user.UserForgetPwdResponse{
		Id: res.UserID,
	}, nil
}
