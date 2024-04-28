package logic

import (
	"calligraphy/apps/user/model"
	"calligraphy/common/cryptx"
	"context"
	"fmt"

	"calligraphy/apps/user/rpc/internal/svc"
	"calligraphy/apps/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserModPwdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserModPwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserModPwdLogic {
	return &UserModPwdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserModPwdLogic) UserModPwd(in *user.UserModPwdRequest) (*user.UserModPwdResponse, error) {
	// todo: add your logic here and delete this line
	//获取信息
	res, err := (&model.User{}).FindOne(l.svcCtx.DB, uint(in.Id))
	if err != nil {
		fmt.Println("rpc UserModPwd 用户修改密码数据库出现问题，或者没有该用户，err", err.Error())
		return nil, err
	}
	res.Password = cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	//进行更新
	err = (&model.User{}).UpdateUser(l.svcCtx.DB, uint(in.Id), res)
	if err != nil {
		fmt.Println("rpc UserModPwd 用户更新密码数据库出现问题 err", err.Error())
		return nil, err
	}
	fmt.Println("rpc UserModPwd 用户修改密码成功")
	return &user.UserModPwdResponse{}, nil
}
