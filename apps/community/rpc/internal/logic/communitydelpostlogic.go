package logic

import (
	"calligraphy/apps/community/model"
	"context"

	"calligraphy/apps/community/rpc/internal/svc"
	"calligraphy/apps/community/rpc/types/community"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommunityDelPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommunityDelPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommunityDelPostLogic {
	return &CommunityDelPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommunityDelPostLogic) CommunityDelPost(in *community.CommunityDelPostRequest) (*community.CommunityDelPostResponse, error) {
	// todo: add your logic here and delete this line
	postOperations := model.Post{}
	_, err := postOperations.DeletePost(l.svcCtx.DB, in.PostId)
	if err != nil {
		l.Logger.Error("rpc 用户在删除帖子的时候，数据库操作出现了问题， err", err.Error())
		return nil, err
	}
	l.Logger.Infof("rpc 用户删除帖子成功。postId:", in.PostId)
	return &community.CommunityDelPostResponse{}, nil
}
