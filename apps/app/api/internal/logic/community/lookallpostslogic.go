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

func (l *LookAllPostsLogic) LookAllPosts(req *types.LookAllPostsRequest) (*types.LookAllPostsResponse, error) {
	// 获取页码和每页大小参数

	pageSize := req.PageSize
	pageSizeNum := 20 // 默认每页大小为20
	if pageSize != 0 {
		pageSizeNum = int(pageSize)
	}

	// 调用RPC进行查找数据
	res, err := l.svcCtx.CommunityRpc.CommunityLookAllPosts(l.ctx, &community.CommunityLookAllPostsRequest{Page: req.Page, PageSize: uint32(pageSizeNum)})

	if err != nil {
		fmt.Println("321", err.Error())
		return nil, err
	}

	// 处理RPC返回结果
	if res == nil {
		return nil, fmt.Errorf("RPC response is nil")
	}

	// 转换数据为本地结构
	var postData []*types.PostInfo
	for _, v := range res.PostData {
		userInfo := types.UserSimpleInfo{
			Id:          uint(v.UserInfo.Id),
			NickName:    v.UserInfo.NickName,
			Account:     v.UserInfo.Account,
			AvatarImage: v.UserInfo.AvatarImage,
		}
		newPostData := &types.PostInfo{
			Id:           uint(v.Id),
			UserId:       uint(v.UserId),
			LikeCount:    uint(v.LikeCount),
			Content:      v.Content,
			ImageUrls:    v.ImageUrls,
			CreateTime:   int32(v.CreateTime),
			CollectCount: uint(v.CollectCount),
			ContentCount: uint(v.ContentCount),
			UserInfo:     userInfo,
		}
		postData = append(postData, newPostData)
	}

	// 构建响应对象并返回
	return &types.LookAllPostsResponse{
		PostData:    postData,
		CurrentPage: res.CurrentPage,
		PageSize:    res.PageSize,
		Offset:      res.Offset,
		Overflow:    res.Overflow,
		TotalPage:   res.TotalPages,
		TotalCount:  res.TotalCount,
	}, nil
}
