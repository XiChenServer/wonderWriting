package logic

import (
	"calligraphy/apps/community/model"
	"context"
	"fmt"

	"calligraphy/apps/community/rpc/internal/svc"
	"calligraphy/apps/community/rpc/types/community"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelCollectPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelCollectPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelCollectPostLogic {
	return &CancelCollectPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CancelCollectPostLogic) CancelCollectPost(in *community.CommunityCancelCollectPostRequest) (*community.CommunityCancelCollectPostResponse, error) {
	// todo: add your logic here and delete this line
	fmt.Println(in.PostId)
	Operations := &model.Collect{}

	err := Operations.CancelCollectPost(l.svcCtx.DB, uint(in.UserId), uint(in.PostId))
	if err != nil {
		l.Logger.Error("rpc 用户取消收藏的时候数据库出现问题，err:", err.Error())
		return nil, err
	}
	l.Logger.Infof("rpc: cancelCollectPost successful")
	return &community.CommunityCancelCollectPostResponse{}, nil
}
