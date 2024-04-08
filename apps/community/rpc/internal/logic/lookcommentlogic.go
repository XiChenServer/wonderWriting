package logic

import (
	"calligraphy/apps/community/model"
	"calligraphy/apps/community/rpc/internal/svc"
	"calligraphy/apps/community/rpc/types/community"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type LookCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLookCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LookCommentLogic {
	return &LookCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// LookComment 查看帖子的评论（带缓存和分页）
func (l *LookCommentLogic) LookComment(in *community.LookCommentRequest) (*community.LookCommentResponse, error) {
	// 构建缓存键
	cacheKey := fmt.Sprintf("post_comments:%d:%d:%d", in.PostId, in.Page, in.PageSize)

	// 尝试从缓存中获取数据
	cachedData, err := getFromCacheForComments(l, cacheKey)
	if err == nil {
		return cachedData, nil // 缓存命中，直接返回缓存数据
	}

	// 缓存未命中，查询数据库获取数据
	commentData, totalCount, err := getComments(l.svcCtx.DB, in.PostId, int(in.Page), int(in.PageSize))
	if err != nil {
		return nil, err
	}

	// 计算总页数
	totalPages := totalCount / int64(in.PageSize)
	if totalCount%int64(in.PageSize) != 0 {
		totalPages++
	}

	// 构建并返回评论信息响应
	resp := &community.LookCommentResponse{
		CommentData: commentData,
		CurrentPage: in.Page,
		PageSize:    in.PageSize,
		Offset:      uint32((in.Page - 1) * in.PageSize),
		Overflow:    in.Page > uint32(totalPages),
		TotalPages:  uint32(totalPages),
		TotalCount:  uint64(totalCount),
	}

	cacheTime := 60 * 5 // 缓存时间为5分钟
	// 将查询结果存入缓存
	err = l.svcCtx.RDB.SetexCtx(l.ctx, cacheKey, toJson(resp), cacheTime)
	if err != nil {
		fmt.Println("Failed to set cache:", err)
	}

	return resp, nil
}

// 查询帖子的评论信息并进行分页
func getComments(db *gorm.DB, postId uint32, page, pageSize int) ([]*community.CommentInfo, int64, error) {
	var res []model.Comment

	offset := (page - 1) * pageSize
	err := db.Where("post_id = ?", postId).Offset(offset).Limit(pageSize).Find(&res).Error
	if err != nil {
		return nil, 0, err
	}

	// 获取总记录数
	var totalCount int64
	err = db.Model(&model.Comment{}).Where("post_id = ?", postId).Count(&totalCount).Error
	if err != nil {
		return nil, 0, err
	}

	// 构建用于返回的评论信息切片
	var commentInfo []*community.CommentInfo

	// 遍历查询到的评论信息
	for _, v := range res {
		// 将 time.Time 转换为 Unix 时间戳 (int64)
		unixTime := v.CreatedAt.Unix()

		// 将 int64 类型的 Unix 时间戳转换为 int32 类型
		int32Time := int32(unixTime)
		newComment := &community.CommentInfo{
			Id:         uint32(v.ID),
			CreateTime: int32Time,
			PostId:     postId,
			Comment:    v.Content,
		}

		// 查询用户信息
		userInfo, err := getUserInfo(db, int(v.UserID))
		if err != nil {
			return nil, 0, err
		}
		newComment.UserInfo = userInfo

		commentInfo = append(commentInfo, newComment)
	}

	return commentInfo, totalCount, nil
}

// 从缓存中获取数据
func getFromCacheForComments(l *LookCommentLogic, key string) (*community.LookCommentResponse, error) {
	val, err := l.svcCtx.RDB.GetCtx(l.ctx, key)
	if err != nil {
		return nil, err
	}

	var data community.LookCommentResponse
	err = json.Unmarshal([]byte(val), &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
