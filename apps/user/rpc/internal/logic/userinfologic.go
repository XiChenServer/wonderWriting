package logic

import (
	"calligraphy/apps/user/model"
	"calligraphy/pkg/qiniu"
	"context"
	"google.golang.org/grpc/status"

	"calligraphy/apps/user/rpc/internal/svc"
	"calligraphy/apps/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	// todo: add your logic here and delete this line
	// 查询用户是否存在
	res, err := (&model.User{}).FindOne(l.svcCtx.DB, uint(in.Id))
	if err != nil {
		if res == nil {
			l.Error("rpc 用户查看信息失败，该用户不存在")
			return nil, status.Error(100, "用户不存在")
		}
		l.Error("rpc 用户查看信息时，数据库操作出现问题")
		return nil, status.Error(500, err.Error())
	}
	l.Info("rpc 用户查看信息成功， user_id", in.Id)
	return &user.UserInfoResponse{
		Id:               int64(res.UserID),
		NickName:         res.Nickname,
		Account:          res.Account,
		Email:            res.Email,
		AvatarBackground: qiniu.ImgUrl + res.AvatarBackground,
		BackgroundImage:  qiniu.ImgUrl + res.BackgroundImage,
		Phone:            res.Phone,
		PointCount:       int64(res.PointCount),
		PostCount:        int64(res.PostCount),
		FansCount:        int64(res.FansCount),
		FollowCount:      int64(res.FollowCount),
		LikeCount:        int64(res.LikeCount),
	}, nil
}
