package logic

import (
	"calligraphy/apps/community/model"
	"calligraphy/apps/community/rpc/internal/svc"
	"calligraphy/apps/community/rpc/types/community"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCollectPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectPostLogic {
	return &CollectPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CollectPostLogic) CollectPost(in *community.CommunityCollectPostRequest) (*community.CommunityCollectPostResponse, error) {
	// todo: add your logic here and delete this line
	Operations := &model.Collect{}
	res, err := Operations.CollectPost(l.svcCtx.DB, uint(in.PostId), uint(in.UserId))
	if err != nil {
		l.Logger.Error("用户收藏帖子失败， err: ", err.Error())
		return nil, err
	}
	l.Logger.Infof("用户收藏帖子成功，user_id:", in.UserId)
	return &community.CommunityCollectPostResponse{CollectId: uint32(res.ID)}, nil
}
