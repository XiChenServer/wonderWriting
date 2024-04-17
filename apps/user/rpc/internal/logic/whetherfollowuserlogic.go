package logic

import (
	"calligraphy/apps/user/model"
	"context"
	"errors"
	"fmt"

	"calligraphy/apps/user/rpc/internal/svc"
	"calligraphy/apps/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type WhetherFollowUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWhetherFollowUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WhetherFollowUserLogic {
	return &WhetherFollowUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

var ErrNotFollow = errors.New("用户未关注该用户")

// 用户是否关注其他人
func (l *WhetherFollowUserLogic) WhetherFollowUser(in *user.WhetherFollowUserRequest) (*user.WhetherFollowUserResponse, error) {
	// todo: add your logic here and delete this line
	// todo: add your logic here and delete this line
	err := (&model.Follow{}).WhetherFollowPeople(l.svcCtx.DB, uint(in.OtherId), uint(in.UserId))
	if err != nil {
		// 如果错误是由于用户未关注该帖子引起的，返回特定的错误信息
		if errors.Is(err, ErrNotFollow) {
			return nil, fmt.Errorf("用户未关注该用户")
		}
		return nil, err // 其他错误则直接返回
	}
	return &user.WhetherFollowUserResponse{}, nil
}
