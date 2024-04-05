package community

import (
	"context"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LookAllPostsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLookAllPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LookAllPostsLogic {
	return &LookAllPostsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LookAllPostsLogic) LookAllPosts() (resp *types.LookAllPostsResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
