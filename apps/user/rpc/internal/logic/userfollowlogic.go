package logic

import (
	"calligraphy/apps/user/model"
	"context"
	"fmt"

	"calligraphy/apps/user/rpc/internal/svc"
	"calligraphy/apps/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFollowLogic {
	return &UserFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户关注
func (l *UserFollowLogic) UserFollow(in *user.UserFollowRequest) (*user.UserFollowResponse, error) {
	// 开启数据库事务
	tx := l.svcCtx.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() // 发生异常时回滚事务
			panic(r)      // 继续抛出异常
		} else {
			tx.Commit() // 没有错误时提交事务
		}
	}()

	// 检查要关注的用户是否存在
	followedUser, err := (&model.User{}).FindOne(tx, uint(in.OtherId))
	if err != nil {
		return nil, err
	}
	if followedUser == nil {
		return nil, fmt.Errorf("user with ID %d not found", in.OtherId)
	}

	// 检查是否已经关注了该用户
	existingFollow, err := (&model.Follow{}).FindOneByFollowerAndFollowed(tx, uint(in.OtherId), uint(in.UserId))
	if err != nil {
		return nil, err
	}
	if existingFollow != nil {
		return nil, fmt.Errorf("user is already followed")
	}

	// 创建关注记录
	newFollow := &model.Follow{
		FollowerUserID: uint(in.UserId),
		FollowedUserID: uint(in.OtherId),
	}
	if err := newFollow.CreateUserFollow(tx, newFollow); err != nil {
		return nil, err
	}

	// 更新用户的关注数和粉丝数
	if err := newFollow.IncrementFollowCount(tx, uint(in.UserId)); err != nil {
		return nil, err
	}
	if err := newFollow.IncrementFansCount(tx, followedUser.UserID); err != nil {
		return nil, err
	}

	// 返回响应
	return &user.UserFollowResponse{}, nil
}
