package logic

import (
	"calligraphy/apps/community/model"
	"calligraphy/pkg/qiniu"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"

	"calligraphy/apps/community/rpc/internal/svc"
	"calligraphy/apps/community/rpc/types/community"

	"github.com/zeromicro/go-zero/core/logx"
)

type LookReplyCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLookReplyCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LookReplyCommentLogic {
	return &LookReplyCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// LookReplyComment 查看回复评论（带缓存和分页）
func (l *LookReplyCommentLogic) LookReplyComment(in *community.LookReplyCommentRequest) (*community.LookReplyCommentResponse, error) {
	// 构建缓存键
	cacheKey := fmt.Sprintf("reply_comments:%d:%d:%d", in.CommentId, in.UserId, in.Page)

	// 尝试从缓存中获取数据
	cachedData, err := getFromCacheForReplyComments(l, cacheKey)
	if err == nil {
		return cachedData, nil // 缓存命中，直接返回缓存数据
	}

	// 缓存未命中，查询数据库获取数据
	replyCommentData, totalCount, err := getReplyComments(l.svcCtx.DB, in.CommentId, in.UserId, int(in.Page), int(in.PageSize))
	if err != nil {
		return nil, err
	}

	// 计算总页数
	totalPages := totalCount / int64(in.PageSize)
	if totalCount%int64(in.PageSize) != 0 {
		totalPages++
	}

	// 构建并返回回复评论信息响应
	resp := &community.LookReplyCommentResponse{
		ReplyCommentData: replyCommentData,
		CurrentPage:      in.Page,
		PageSize:         in.PageSize,
		Offset:           uint32((in.Page - 1) * in.PageSize),
		Overflow:         in.Page > uint32(totalPages),
		TotalPages:       uint32(totalPages),
		TotalCount:       uint64(totalCount),
	}

	cacheTime := 60 * 5 // 缓存时间为5分钟
	// 将查询结果存入缓存
	err = setToCacheForReplyComments(l, cacheKey, resp, cacheTime)
	if err != nil {
		fmt.Println("Failed to set cache:", err)
	}

	return resp, nil
}

// 查询回复评论信息并进行分页
func getReplyComments(db *gorm.DB, commentID, userID uint32, page, pageSize int) ([]*community.ReplyCommentInfo, int64, error) {
	res, err := (&model.ReplyComment{}).FindReplyCommentsByPage(db, uint(commentID), uint(userID), uint(page), uint(pageSize))
	if err != nil {
		return nil, 0, err
	}
	// 获取总记录数
	totalCount, err := (&model.ReplyComment{}).FindReplyCommentCount(db, uint(commentID), uint(userID))
	if err != nil {
		return nil, 0, err
	}

	// 构建用于返回的回复评论信息切片
	var replyCommentInfo []*community.ReplyCommentInfo

	// 遍历查询到的回复评论信息
	for _, v := range *res {
		// 将 time.Time 转换为 Unix 时间戳 (int64)
		unixTime := v.CreatedAt.Unix()

		// 将 int64 类型的 Unix 时间戳转换为 int32 类型
		int32Time := int32(unixTime)
		newReplyComment := &community.ReplyCommentInfo{
			Id:            uint32(v.ID),
			CommentId:     uint32(v.CommentID),
			UserId:        uint32(v.UserID),
			UserNickName:  v.UserNickName,
			UserAvatar:    qiniu.ImgUrl + v.UserAvatar,
			Content:       v.Content,
			ReplyNickName: v.ReplyNickName,
			ReplyUserId:   uint32(v.ReplyUserId),
			LikeCount:     uint32(v.LikeCount),
			CreateTime:    int32Time,
		}

		replyCommentInfo = append(replyCommentInfo, newReplyComment)
	}

	return replyCommentInfo, totalCount, nil
}

// 从缓存中获取回复评论数据
func getFromCacheForReplyComments(l *LookReplyCommentLogic, key string) (*community.LookReplyCommentResponse, error) {
	val, err := l.svcCtx.RDB.GetCtx(l.ctx, key)
	if err != nil {
		return nil, err
	}

	var data community.LookReplyCommentResponse
	err = json.Unmarshal([]byte(val), &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// 将查询结果存入缓存
func setToCacheForReplyComments(l *LookReplyCommentLogic, key string, data *community.LookReplyCommentResponse, expiration int) error {
	// 序列化数据
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// 存入缓存
	err = l.svcCtx.RDB.SetexCtx(l.ctx, key, string(jsonData), expiration)
	if err != nil {
		return err
	}

	return nil
}
