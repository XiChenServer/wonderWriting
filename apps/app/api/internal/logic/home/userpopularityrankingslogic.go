package home

import (
	"calligraphy/apps/home/rpc/types/home"
	"context"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserPopularityRankingsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserPopularityRankingsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPopularityRankingsLogic {
	return &UserPopularityRankingsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserPopularityRankingsLogic) UserPopularityRankings(req *types.UserPopularityRankingsRequest) (resp *types.UserPopularityRankingsResponse, err error) {
	// todo: add your logic here and delete this line
	var pageSize uint32 = 20
	if req.PageSize > 0 {
		pageSize = req.PageSize
	}
	res, err := l.svcCtx.HomeRpc.UserPopularityRankings(l.ctx, &home.UserPopularityRankingsRequest{
		Page:     req.Page,
		PageSize: pageSize,
	})
	if err != nil {
		return nil, err
	}
	var userPopularData []*types.UserPopularInfo
	for _, v := range res.UserPopularData {
		newUserPopularData := &types.UserPopularInfo{
			UserId:    v.UserId,
			NickName:  v.NickName,
			Account:   v.Account,
			LikeCount: v.LikeCount,
			Avatar:    v.Avatar,
		}
		userPopularData = append(userPopularData, newUserPopularData)
	}
	return &types.UserPopularityRankingsResponse{
		UserPopularData: userPopularData,
		CurrentPage:     res.CurrentPage,
		PageSize:        res.PageSize,
		Offset:          res.Offset,
		Overflow:        res.Overflow,
		TotalPage:       res.TotalPages,
		TotalCount:      res.TotalCount,
	}, nil
}
