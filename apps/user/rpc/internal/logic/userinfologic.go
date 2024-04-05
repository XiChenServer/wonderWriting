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
	res, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	return &user.UserInfoResponse{
		Id:               res.UserID,
		NickName:         res.Nickname,
		Account:          res.Account,
		Email:            res.Email.String,
		AvatarBackground: qiniu.ImgUrl + res.AvatarBackground.String,
		BackgroundImage:  qiniu.ImgUrl + res.BackgroundImage.String,
		Phone:            res.Phone,
		PointCount:       res.PointCount,
		PostCount:        res.PointCount,
		FansCount:        res.FansCount,
		FollowCount:      res.FollowCount,
		LikeCount:        res.LikeCount,
	}, nil
}
