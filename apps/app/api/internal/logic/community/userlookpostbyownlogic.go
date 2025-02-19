package community

import (
	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"
	"calligraphy/apps/community/rpc/types/community"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserLookPostByOwnLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLookPostByOwnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLookPostByOwnLogic {
	return &UserLookPostByOwnLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLookPostByOwnLogic) UserLookPostByOwn(req *types.LookPostByOwnRequest) (resp *types.LookPostByOwnResponses, err error) {
	// todo: add your logic here and delete this line
	//从jwt获取id
	var pageSize uint32 = 20
	if req.PageSize > 0 {
		pageSize = req.PageSize
	}
	//调用rpc进行查找数据
	res, err := l.svcCtx.CommunityRpc.CommunityLookPostByOwn(l.ctx, &community.CommunityLookPostByOwnRequest{
		UserId:   req.UserId,
		Page:     req.Page,
		PageSize: pageSize,
	})
	if err != nil {
		return &types.LookPostByOwnResponses{}, err
	}

	//进行转换数据
	var postData []*types.PostInfo
	for _, v := range res.PostData {
		var userInfo = types.UserSimpleInfo{
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
		fmt.Println(v.LikeCount)
		fmt.Println(newPostData.LikeCount)
		postData = append(postData, newPostData)
	}
	return &types.LookPostByOwnResponses{
		PostData:    postData,
		CurrentPage: res.CurrentPage,
		PageSize:    res.PageSize,
		Offset:      res.Offset,
		Overflow:    res.Overflow,
		TotalPage:   res.TotalPages,
		TotalCount:  res.TotalCount,
	}, err
}
