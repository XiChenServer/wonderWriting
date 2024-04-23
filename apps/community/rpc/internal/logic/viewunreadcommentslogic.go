package logic

import (
	"calligraphy/apps/community/model"
	"calligraphy/apps/community/rpc/internal/svc"
	"calligraphy/apps/community/rpc/types/community"
	"context"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type ViewUnreadCommentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewViewUnreadCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewUnreadCommentsLogic {
	return &ViewUnreadCommentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ViewUnreadComments 查看未读的评论（带缓存和分页）
func (l *ViewUnreadCommentsLogic) ViewUnreadComments(in *community.ViewUnreadCommentsRequest) (*community.ViewUnreadCommentsResponse, error) {
	// 构建缓存键
	cacheKey := fmt.Sprintf("unread_comments:%d:%d:%d", in.UserId, in.Page, in.PageSize)

	// 尝试从缓存中获取数据
	cachedData, err := getFromCacheForUnreadComments(l, cacheKey)
	if err == nil {
		return cachedData, nil // 缓存命中，直接返回缓存数据
	}

	// 缓存未命中，查询数据库获取数据
	userUnreadMessages, err := (&model.UserUnreadMessages{}).UserUnreadMessages(l.svcCtx.DB, in.UserId)
	if err != nil {
		return nil, err
	}

	// 从未读消息记录中获取未读评论列表
	var unreadComments []*community.CommentInfo
	for _, v := range userUnreadMessages.Comments {
		newUnreadComments := community.CommentInfo{
			Id:           uint32(v.ID),
			CreateTime:   int32(v.CreatedAt.Unix()),
			PostId:       uint32(v.PostID),
			Comment:      v.Content,
			UserAvatar:   v.UserAvatar,
			UserNickname: v.UserNickName,
			LikeCount:    uint32(v.LikeCount),
			UserId:       uint32(v.UserID),
		}
		unreadComments = append(unreadComments, &newUnreadComments)
	}

	var unreadReplies []*community.ReplyCommentInfo
	for _, v := range userUnreadMessages.Replies {
		newUnreadReplies := community.ReplyCommentInfo{
			Id:            uint32(v.ID),
			CommentId:     uint32(v.CommentID),
			UserId:        uint32(v.UserID),
			UserNickName:  v.UserNickName,
			UserAvatar:    v.UserAvatar,
			Content:       v.Content,
			LikeCount:     uint32(v.LikeCount),
			CreateTime:    int32(v.CreatedAt.Unix()),
			PostId:        int32(v.PostId),
			ReplyNickName: v.ReplyNickName,
			ReplyUserId:   uint32(v.ReplyUserId),
		}
		unreadReplies = append(unreadReplies, &newUnreadReplies)
	}
	// 计算总页数
	totalComments := len(unreadComments)
	totalReplies := len(unreadReplies)
	totalRecords := totalComments + totalReplies
	totalPages := totalRecords / int(in.PageSize)
	if totalRecords%int(in.PageSize) != 0 {
		totalPages++
	}

	// 分页处理未读评论和回复评论
	startIndex := int(in.Page-1) * int(in.PageSize)
	endIndex := startIndex + int(in.PageSize)
	if endIndex > totalRecords {
		endIndex = totalRecords
	}
	var pagedComments []*community.CommentInfo
	var pagedReplies []*community.ReplyCommentInfo

	if startIndex < totalComments {

		pagedComments = unreadComments[startIndex:endIndex]
	} else {
		startIndex -= totalComments
		pagedReplies = unreadReplies[startIndex:endIndex]
	}

	// 构建 ViewUnreadCommentsResponse 对象
	response := &community.ViewUnreadCommentsResponse{
		CommentsData:     pagedComments,
		ReplyCommentData: pagedReplies,
		CurrentPage:      in.Page,
		PageSize:         in.PageSize,
		Offset:           uint32(startIndex),
		Overflow:         in.Page > uint32(totalPages),
		TotalPages:       uint32(totalPages),
		TotalCount:       uint64(totalRecords),
	}

	// 将查询结果存入缓存
	cacheTime := 60 * 5 // 缓存时间为5分钟
	err = l.svcCtx.RDB.SetexCtx(l.ctx, cacheKey, toJson(response), cacheTime)
	if err != nil {
		fmt.Println("Failed to set cache:", err)
	}

	return nil, nil
}

// 从缓存中获取未读评论数据
func getFromCacheForUnreadComments(l *ViewUnreadCommentsLogic, key string) (*community.ViewUnreadCommentsResponse, error) {
	val, err := l.svcCtx.RDB.GetCtx(l.ctx, key)
	if err != nil {
		return nil, err
	}

	var data community.ViewUnreadCommentsResponse
	err = json.Unmarshal([]byte(val), &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
