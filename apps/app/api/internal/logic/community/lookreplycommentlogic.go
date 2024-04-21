package community

import (
	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"
	"calligraphy/apps/community/rpc/types/community"
	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
)

type LookReplyCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLookReplyCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LookReplyCommentLogic {
	return &LookReplyCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LookReplyCommentLogic) LookReplyComment(req *types.LookReplyCommentRequest) (resp *types.LookReplyCommentResponse, err error) {
	// todo: add your logic here and delete this line
	var pageSize uint32 = 20
	if req.PageSize > 0 {
		pageSize = req.PageSize
	}
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	//调用rpc进行查找数据
	res, err := l.svcCtx.CommunityRpc.LookReplyComment(l.ctx, &community.LookReplyCommentRequest{
		Page:      uint32(req.Page),
		PageSize:  pageSize,
		CommentId: uint32(req.CommentId),
		UserId:    uint32(uid),
	})
	if err != nil {
		return &types.LookReplyCommentResponse{}, err
	}

	//进行转换数据
	var replyCommen []*types.ReplyCommentInfo
	for _, v := range res.ReplyCommentData {
		var newReplyComment = types.ReplyCommentInfo{
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
		}

		replyCommen = append(replyCommen, &newReplyComment)
	}
	return &types.LookReplyCommentResponse{
		ReplyCommentData: replyCommen,
		CurrentPage:      res.CurrentPage,
		PageSize:         res.PageSize,
		Offset:           res.Offset,
		Overflow:         res.Overflow,
		TotalPage:        res.TotalPages,
		TotalCount:       res.TotalCount,
	}, nil
}
