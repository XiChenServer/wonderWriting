package community

import (
	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"
	"calligraphy/apps/community/rpc/types/community"
	"context"
	"fmt"

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

	//调用rpc进行查找数据
	res, err := l.svcCtx.CommunityRpc.CommunityLookAllPosts(l.ctx, &community.CommunityLookAllPostsRequest{})
	if err != nil {
		return nil, err
	}
	//进行转换数据
	var postData []*types.PostInfo
	for _, v := range res.PostData {
		newPostData := &types.PostInfo{
			Id:           uint(v.Id),
			UserId:       uint(v.UserId),
			LikeCount:    uint(v.LikeCount),
			Content:      v.Content,
			ImageUrls:    v.ImageUrls,
			CreateTime:   int32(v.CreateTime),
			CollectCount: uint(v.CollectCount),
			ContentCount: uint(v.ContentCount),
		}
		fmt.Println(v.LikeCount)
		fmt.Println(newPostData.LikeCount)
		fmt.Println("123", newPostData)
		postData = append(postData, newPostData)
	}
	return &types.LookAllPostsResponse{PostData: postData}, nil
}
