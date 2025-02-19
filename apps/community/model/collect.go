package model

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

// 收藏帖子表
type Collect struct {
	gorm.Model
	PostID uint `json:"post_id"` // 帖子ID，JSON序列化时的字段名为"post_id"
	UserID uint `json:"user_id"` // 用户ID，JSON序列化时的字段名为"user_id"
}

// FindCollect 查看收藏
func (*Collect) FindCollect(DB *gorm.DB, userId uint) (*[]Collect, error) {
	var collect []Collect
	err := DB.Where("user_id = ?", userId).Find(&collect).Error
	if err != nil {
		return nil, err
	}
	return &collect, nil
}

// CollectPost 在数据库中创建一条收藏记录并原子更新帖子的收藏数量
func (*Collect) CollectPost(DB *gorm.DB, postID, userID uint) (*Collect, error) {
	// 检查是否已经收藏过
	var existingCollect Collect
	err := DB.Where("post_id = ? AND user_id = ?", postID, userID).First(&existingCollect).Error
	if err == nil {
		// 已经收藏过，返回错误信息
		return nil, fmt.Errorf("post already collected by user")
	}
	// 开始事务
	tx := DB.Begin()

	// 创建收藏记录
	collect := &Collect{
		PostID: postID,
		UserID: userID,
	}
	if err := tx.Create(collect).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// 原子操作更新帖子的收藏数量
	if err := tx.Model(&Post{}).Where("id = ?", postID).UpdateColumn("collection_count", gorm.Expr("collection_count + 1")).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return collect, nil
}

// CancelCollectPost 在数据库中删除收藏记录并原子更新帖子的收藏数量
func (*Collect) CancelCollectPost(DB *gorm.DB, userID, postID uint) error {
	// 获取收藏记录
	var collect Collect
	if err := DB.Where("user_id = ? AND post_id = ?", userID, postID).First(&collect).Error; err != nil {
		return err
	}

	// 原子操作删除收藏记录并更新帖子的收藏数量
	err := DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&collect).Error; err != nil {
			return err
		}

		// 查询帖子并更新收藏数量
		var post Post
		if err := tx.First(&post, postID).Error; err != nil {
			return err
		}
		if post.CollectionCount > 0 {
			post.CollectionCount--
			if err := tx.Save(&post).Error; err != nil {
				return err
			}
		}
		return nil
	})

	return err
}

// WhetherCollectPost 检查用户是否收藏了某个帖子
func (*Collect) WhetherCollectPost(DB *gorm.DB, postID, userID uint) error {
	var collect Collect
	err := DB.Where("post_id = ? AND user_id = ?", postID, userID).First(&collect).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("用户未收藏该帖子")
		}
		return err
	}
	return nil
}
