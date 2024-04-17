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

// 用户关注
func (s *UserServer) UserFollow(ctx context.Context, in *user.UserFollowRequest) (*user.UserFollowResponse, error) {
	l := logic.NewUserFollowLogic(ctx, s.svcCtx)
	return l.UserFollow(in)
}

// 用户取消关注
func (s *UserServer) UserCancelFollow(ctx context.Context, in *user.UserCancelFollowRequest) (*user.UserCancelFollowResponse, error) {
	l := logic.NewUserCancelFollowLogic(ctx, s.svcCtx)
	return l.UserCancelFollow(in)
}

// 用户查看自己的粉丝
func (s *UserServer) LookAllFans(ctx context.Context, in *user.LookAllFansRequest) (*user.LookAllFansResponse, error) {
	l := logic.NewLookAllFansLogic(ctx, s.svcCtx)
	return l.LookAllFans(in)
}

// 用户查看自己的关注
func (s *UserServer) LookAllFollow(ctx context.Context, in *user.LookAllFollowRequest) (*user.LookAllFollowResponse, error) {
	l := logic.NewLookAllFollowLogic(ctx, s.svcCtx)
	return l.LookAllFollow(in)
}

// 用户是否关注其他人
func (s *UserServer) WhetherFollowUser(ctx context.Context, in *user.WhetherFollowUserRequest) (*user.WhetherFollowUserResponse, error) {
	l := logic.NewWhetherFollowUserLogic(ctx, s.svcCtx)
	return l.WhetherFollowUser(in)
}
