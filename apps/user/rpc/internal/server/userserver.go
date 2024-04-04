// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"calligraphy/apps/user/rpc/internal/logic"
	"calligraphy/apps/user/rpc/internal/svc"
	"calligraphy/apps/user/rpc/types/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Login(ctx context.Context, in *user.UserLoginRequest) (*user.UserLoginResponse, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) Register(ctx context.Context, in *user.UserRegisterRequest) (*user.UserRegisterResponse, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserServer) UserInfo(ctx context.Context, in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	l := logic.NewUserInfoLogic(ctx, s.svcCtx)
	return l.UserInfo(in)
}

func (s *UserServer) UserForgetPwd(ctx context.Context, in *user.UserForgetPwdRequest) (*user.UserForgetPwdResponse, error) {
	l := logic.NewUserForgetPwdLogic(ctx, s.svcCtx)
	return l.UserForgetPwd(in)
}

func (s *UserServer) UserModPwd(ctx context.Context, in *user.UserModPwdRequest) (*user.UserModPwdResponse, error) {
	l := logic.NewUserModPwdLogic(ctx, s.svcCtx)
	return l.UserModPwd(in)
}

func (s *UserServer) UserModAvatar(ctx context.Context, in *user.UserModAvatarRequest) (*user.UserModAvatarResponse, error) {
	l := logic.NewUserModAvatarLogic(ctx, s.svcCtx)
	return l.UserModAvatar(in)
}

func (s *UserServer) UserModBackground(ctx context.Context, in *user.UserModBackgroundRequest) (*user.UserModBackgroundResponse, error) {
	l := logic.NewUserModBackgroundLogic(ctx, s.svcCtx)
	return l.UserModBackground(in)
}

func (s *UserServer) UserModInfo(ctx context.Context, in *user.UserModInfoRequest) (*user.UserModInfoResponse, error) {
	l := logic.NewUserModInfoLogic(ctx, s.svcCtx)
	return l.UserModInfo(in)
}
