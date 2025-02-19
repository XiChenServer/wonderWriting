package community

import (
	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"
	"calligraphy/apps/community/rpc/types/community"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type LookCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLookCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LookCommentLogic {
	return &LookCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LookCommentLogic) LookComment(req *types.LookCommentRequest) (resp *types.LookCommentResponse, err error) {
	// todo: add your logic here and delete this line
	var pageSize uint32 = 20
	if req.PageSize > 0 {
		pageSize = req.PageSize
	}

	//用rpc获取数据
	res, err := l.svcCtx.CommunityRpc.LookComment(l.ctx, &community.LookCommentRequest{
		PostId:   uint32(req.PostId),
		Page:     uint32(req.Page),
		PageSize: pageSize,
	})
	if err != nil {
		return &types.LookCommentResponse{}, err
	}

	//进行数据的转换
	var commentData []*types.CommentInfo
	for _, v := range res.CommentData {
		newCommentData := &types.CommentInfo{
			Id:         uint(v.Id),
			CreateTime: v.CreateTime,
			PostId:     uint(v.PostId),
			Comment:    v.Comment,
		}
		newUserInfo := types.UserSimpleInfo{
			Id:          uint(v.UserInfo.Id),
			NickName:    v.UserInfo.NickName,
			Account:     v.UserInfo.Account,
			AvatarImage: v.UserInfo.AvatarImage,
		}
		newCommentData.UserInfo = newUserInfo
		commentData = append(commentData, newCommentData)
	}

	return &types.LookCommentResponse{
		CommentData: commentData,
		CurrentPage: res.CurrentPage,
		PageSize:    res.PageSize,
		Offset:      res.Offset,
		Overflow:    res.Overflow,
		TotalPage:   res.TotalPages,
		TotalCount:  res.TotalCount,
	}, nil

}
