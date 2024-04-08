package home

import (
	"calligraphy/apps/home/rpc/types/home"
	"context"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostPopularityRankingsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostPopularityRankingsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostPopularityRankingsLogic {
	return &PostPopularityRankingsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostPopularityRankingsLogic) PostPopularityRankings(req *types.PostPopularityRankingsRequest) (resp *types.PostPopularityRankingsResponse, err error) {
	// todo: add your logic here and delete this line
	//调用rpc获取数据
	var pageSize uint32 = 20
	if req.PageSize > 0 {
		pageSize = req.PageSize
	}
	res, err := l.svcCtx.HomeRpc.PostPopularityRankings(l.ctx, &home.PostPopularityRankingsRequest{
		Page:     req.Page,
		PageSize: pageSize,
	})
	if err != nil {
		return nil, err
	}

	//对于数据进行转换
	var postInfo []*types.PostPopularityInfo
	for _, v := range res.PostPopularData {
		newPostInfo := &types.PostPopularityInfo{
			PostId:          v.PostId,
			Content:         v.Content,
			LikeCont:        v.LikeCont,
			CollectionCount: v.CollectionCount,
			CommentCount:    v.CommentCount,
			PostImage:       v.PostImage,
		}
		newUser := &types.UserPopularInfo{
			UserId:    v.PopularInfo.UserId,
			NickName:  v.PopularInfo.NickName,
			Account:   v.PopularInfo.Account,
			LikeCount: v.PopularInfo.LikeCount,
			Avatar:    v.PopularInfo.Avatar,
		}
		newPostInfo.PopularInfo = newUser
		postInfo = append(postInfo, newPostInfo)
	}
	return &types.PostPopularityRankingsResponse{
		PostPopularData: postInfo,
		CurrentPage:     res.CurrentPage,
		PageSize:        res.PageSize,
		Offset:          res.Offset,
		Overflow:        res.Overflow,
		TotalPage:       res.TotalPages,
		TotalCount:      res.TotalCount,
	}, nil
}
