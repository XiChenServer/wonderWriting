package logic

import (
	"calligraphy/apps/community/model"
	"context"
	"errors"
	"fmt"

	"calligraphy/apps/community/rpc/internal/svc"
	"calligraphy/apps/community/rpc/types/community"

	"github.com/zeromicro/go-zero/core/logx"
)

type WhetherLikePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWhetherLikePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WhetherLikePostLogic {
	return &WhetherLikePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

var ErrNotLiked = errors.New("用户未点赞该帖子")

// WhetherLikePost 用户是否点赞帖子
func (l *WhetherLikePostLogic) WhetherLikePost(in *community.WhetherLikePostRequest) (*community.WhetherLikePostResponse, error) {
	err := (&model.Like{}).WhetherLikedPost(l.svcCtx.DB, uint(in.PostId), uint(in.UserId))
	if err != nil {
		// 如果错误是由于用户未关注该帖子引起的，返回特定的错误信息
		if errors.Is(err, ErrNotLiked) {
			return nil, fmt.Errorf("用户未点赞该帖子")
		}
		return nil, err // 其他错误则直接返回
	}
	return &community.WhetherLikePostResponse{}, nil
}
