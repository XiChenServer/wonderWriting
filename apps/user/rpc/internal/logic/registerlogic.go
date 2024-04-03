package logic

import (
	"calligraphy/apps/user/model"
	"calligraphy/common/cryptx"
	"calligraphy/pkg/app_math"
	"context"
	"database/sql"
	"google.golang.org/grpc/status"

	"calligraphy/apps/user/rpc/internal/svc"
	"calligraphy/apps/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.UserRegisterRequest) (*user.UserRegisterResponse, error) {
	// todo: add your logic here and delete this line
	email := sql.NullString{
		String: in.Email,
		Valid:  true,
	}
	// 判断手机号是否已经注册
	_, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, email)
	if err == nil {
		return nil, status.Error(100, "该用户已存在")
	}
	var account string
	for {
		account = app_math.GenerateRandomNumber(11)
		_, err = l.svcCtx.UserModel.FindOneByAccount(l.ctx, account)
		if err != nil {
			break
		}
	}

	if err == model.ErrNotFound {
		newUser := model.Users{
			Nickname: in.NickName,
			Account:  account,
			Email:    email,
			Password: cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
		}

		res, err := l.svcCtx.UserModel.Insert(l.ctx, &newUser)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		newUser.UserID, err = res.LastInsertId()
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		return &user.UserRegisterResponse{
			Code:    200,
			Message: "Success",
		}, nil

	}

	return nil, status.Error(500, err.Error())
}
