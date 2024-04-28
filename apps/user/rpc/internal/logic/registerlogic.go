package logic

import (
	"calligraphy/apps/user/model"
	"calligraphy/common/cryptx"
	"calligraphy/pkg/app_math"
	"context"
	"database/sql"
	"google.golang.org/grpc/status"
	"time"

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
	// 查询用户是否存在
	email := sql.NullString{
		String: in.Email,
		Valid:  true,
	}

	// 判断手机号或邮箱是否已经注册
	res, err := (&model.User{}).FindOneByEmail(l.svcCtx.DB, email.String)
	if err == nil {
		if res != nil {
			l.Error("rpc 用户在登录的时候发现，用户不存在，err", err, "email", email.String)
			return nil, status.Error(100, "该用户已存在")
		}

	}

	// 生成随机账号
	var account string
	for {
		account = app_math.GenerateRandomNumber(11)
		res, err = (&model.User{}).FindOneByAccount(l.svcCtx.DB, account)
		if res == nil {
			break
		}
	}

	// 生成随机昵称
	nickName := app_math.GenerateNickname(8)
	// 插入新用户信息
	nowTime := time.Now()
	newUser := model.User{
		Nickname:         nickName,
		Account:          account,
		Email:            email.String,
		Password:         cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
		RegistrationTime: nowTime.Unix(),
		LastLoginTime:    nowTime.Unix(),
		Status:           "Active",
		Role:             "User",
		AvatarBackground: "AvatarBackground/97e4cf398c1c453f98f8135b202479d6.jpeg",
		BackgroundImage:  "BackgroundImage/kpmg3R46Q2.jpg",
	}

	err = (&model.User{}).InsertUser(l.svcCtx.DB, &newUser)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &user.UserRegisterResponse{
		Code:    200,
		Message: "Success",
	}, nil
}
