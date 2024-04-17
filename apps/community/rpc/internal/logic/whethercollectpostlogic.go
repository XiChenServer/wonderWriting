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

type WhetherCollectPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWhetherCollectPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WhetherCollectPostLogic {
	return &WhetherCollectPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

var ErrNotCollect = errors.New("用户未收藏该帖子")

// WhetherCollectPost 用户是否收藏帖子
func (l *WhetherCollectPostLogic) WhetherCollectPost(in *community.WhetherCollectPostRequest) (*community.WhetherCollectPostResponse, error) {
	// todo: add your logic here and delete this line
	err := (&model.Collect{}).WhetherCollectPost(l.svcCtx.DB, uint(in.PostId), uint(in.UserId))
	if err != nil {
		// 如果错误是由于用户未关注该帖子引起的，返回特定的错误信息
		if errors.Is(err, ErrNotCollect) {
			return nil, fmt.Errorf("用户未收藏该帖子")
		}
		return nil, err // 其他错误则直接返回
	}
	return &community.WhetherCollectPostResponse{}, nil
}
