package community

import (
	"calligraphy/apps/community/rpc/types/community"
	"context"
	"encoding/json"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ViewUnreadCommentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewViewUnreadCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewUnreadCommentsLogic {
	return &ViewUnreadCommentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ViewUnreadCommentsLogic) ViewUnreadComments(req *types.ViewUnreadCommentsRequest) (resp *types.ViewUnreadCommentsResponse, err error) {
	// todo: add your logic here and delete this line
	var pageSize uint32 = 20
	if req.PageSize > 0 {
		pageSize = req.PageSize
	}
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	res, err := l.svcCtx.CommunityRpc.ViewUnreadComments(l.ctx, &community.ViewUnreadCommentsRequest{
		UserId:   uint32(uid),
		Page:     uint32(req.Page),
		PageSize: pageSize,
	})
	if err != nil {
		return &types.ViewUnreadCommentsResponse{}, err
	}
	var commentData []*types.CommentInfo
	for _, v := range res.CommentsData {
		newUserInfo := types.UserSimpleInfo{
			Id:          uint(v.UserInfo.Id),
			NickName:    v.UserNickname,
			Account:     v.UserInfo.Account,
			AvatarImage: v.UserAvatar,
		}
		newCommentData := &types.CommentInfo{
			Id:         uint(v.Id),
			CreateTime: v.CreateTime,
			PostId:     uint(v.PostId),
			Comment:    v.Comment,
			UserInfo:   newUserInfo,
		}
		commentData = append(commentData, newCommentData)
	}

	var replyCommentData []*types.ReplyCommentInfo
	for _, v := range res.ReplyCommentData {
		newReplyCommentData := &types.ReplyCommentInfo{
			Id:            v.Id,
			CommentId:     v.CommentId,
			UserId:        v.UserId,
			UserNickName:  v.UserNickName,
			UserAvatar:    v.UserAvatar,
			Content:       v.Content,
			ReplyNickName: v.ReplyNickName,
			ReplyUserId:   v.ReplyUserId,
			LikeCount:     v.LikeCount,
			CreateTime:    v.CreateTime,
			PostId:        v.PostId,
		}
		replyCommentData = append(replyCommentData, newReplyCommentData)
	}

	return &types.ViewUnreadCommentsResponse{
		CommentData:      commentData,
		ReplyCommentData: replyCommentData,
		CurrentPage:      res.CurrentPage,
		PageSize:         res.PageSize,
		Offset:           res.Offset,
		Overflow:         res.Overflow,
		TotalPage:        res.TotalPages,
		TotalCount:       res.TotalCount,
	}, nil
}
