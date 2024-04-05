package model

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Like struct {
	gorm.Model
	PostID uint `json:"post_id"` // 帖子ID，JSON序列化时的字段名为"post_id"
	UserID uint `json:"user_id"` // 用户ID，JSON序列化时的字段名为"user_id"
}

// LikePost 在数据库中创建一条点赞记录并原子更新帖子的点赞数量
func (*Like) LikePost(DB *gorm.DB, postID, userID uint) (*Like, error) {
	// 查询帖子是否存在
	post := &Post{}
	err := DB.Where("id = ?", postID).First(post).Error
	if err != nil {
		// 帖子不存在，返回错误信息
		return nil, errors.New("post not found")
	}

	// 检查用户是否已经点赞过该帖子
	existingLike := &Like{}
	err = DB.Where("post_id = ? AND user_id = ?", postID, userID).First(existingLike).Error
	if err == nil {
		// 如果已经点赞过，则检查是否已删除点赞记录
		if existingLike.DeletedAt != nil {
			return nil, errors.New("user already liked this post and unliked it")
		}
		// 如果没有删除，则返回错误信息
		return nil, errors.New("user already liked this post")
	}

	// 创建点赞记录
	like := &Like{
		PostID: postID,
		UserID: userID,
	}

	// 使用原子操作创建点赞记录并更新帖子的点赞数量
	err = DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(like).Error; err != nil {
			return err
		}

		// 原子操作更新帖子的点赞数量
		if err := tx.Model(&Post{}).Where("id = ?", postID).Update("like_count", gorm.Expr("like_count + 1")).Error; err != nil {
			return err
		}

		return nil
	})

	return like, err
}

// CancelLikePost 在数据库中删除点赞记录并原子更新帖子的点赞数量
func (*Like) CancelLikePost(DB *gorm.DB, likeID, postID uint) error {
	// 获取点赞记录
	var like Like
	if err := DB.Where("id = ? AND post_id = ?", likeID, postID).First(&like).Error; err != nil {
		return err
	}

	// 使用原子操作删除点赞记录并更新帖子的点赞数量
	err := DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&like).Error; err != nil {
			return err
		}

		// 原子操作更新帖子的点赞数量
		if err := tx.Model(&Post{}).Where("id = ?", postID).Update("like_count", gorm.Expr("like_count - 1")).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}
