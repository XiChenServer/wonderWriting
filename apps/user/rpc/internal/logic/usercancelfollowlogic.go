package logic

import (
	"calligraphy/apps/user/model"
	"context"
	"fmt"

	"calligraphy/apps/user/rpc/internal/svc"
	"calligraphy/apps/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCancelFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCancelFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCancelFollowLogic {
	return &UserCancelFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UserCancelFollow 用户取消关注
func (l *UserCancelFollowLogic) UserCancelFollow(in *user.UserCancelFollowRequest) (*user.UserCancelFollowResponse, error) {
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

	// 检查要取消关注的用户是否存在
	followedUser, err := (&model.User{}).FindOne(tx, uint(in.OtherId))
	if err != nil {
		return nil, err
	}
	if followedUser == nil {
		return nil, fmt.Errorf("user with ID %d not found", in.OtherId)
	}

	// 检查是否已经关注了该用户
	existingFollow, err := (&model.Follow{}).FindOneByFollowerAndFollowed(tx, uint(in.OtherId), followedUser.UserID)
	if err != nil {
		return nil, err
	}
	if existingFollow == nil {
		return nil, fmt.Errorf("user with ID %d is not followed by user with ID %d", in.OtherId, in.UserId)
	}

	// 删除关注记录
	if err := existingFollow.DeleteFollow(tx); err != nil {
		return nil, err
	}

	// 更新用户的关注数和粉丝数
	if err := existingFollow.DecrementFollowCount(tx, uint(in.UserId)); err != nil {
		return nil, err
	}
	if err := existingFollow.DecrementFansCount(tx, followedUser.UserID); err != nil {
		return nil, err
	}

	// 返回响应
	return &user.UserCancelFollowResponse{}, nil
}
