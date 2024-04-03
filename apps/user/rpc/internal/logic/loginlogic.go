package logic

import (
	"calligraphy/apps/user/model"
	"calligraphy/common/cryptx"
	"context"
	"database/sql"
	"google.golang.org/grpc/status"

	"calligraphy/apps/user/rpc/internal/svc"
	"calligraphy/apps/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.UserLoginRequest) (*user.UserLoginResponse, error) {
	// todo: add your logic here and delete this line
	// 查询用户是否存在
	email := sql.NullString{
		String: in.Email,
		Valid:  true,
	}
	res, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, email)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	// 判断密码是否正确
	password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	if password != res.Password {
		return nil, status.Error(100, "密码错误")
	}

	return &user.UserLoginResponse{
		Code:    200,
		Message: "Success",
	}, nil
}
