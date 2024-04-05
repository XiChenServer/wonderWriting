package model

import (
	"github.com/jinzhu/gorm"
)

// Comment 评论表模型
type Comment struct {
	gorm.Model
	PostID  uint   `json:"post_id"`           // 帖子ID，JSON序列化时的字段名为"post_id"
	Post    Post   `gorm:"foreignKey:PostID"` // 关联的帖子，使用PostID作为外键
	UserID  uint   `json:"user_id"`           // 用户ID，JSON序列化时的字段名为"user_id"
	Content string `json:"content"`           // 评论内容，JSON序列化时的字段名为"content"

}

// CommentPost 在数据库中创建一条评论记录并原子更新帖子的评论数量
func (*Comment) CommentPost(DB *gorm.DB, postID, userID uint, content string) (*Comment, error) {
	// 开始事务
	tx := DB.Begin()

	// 创建评论记录
	comment := &Comment{
		PostID:  postID,
		UserID:  userID,
		Content: content,
	}
	if err := tx.Create(comment).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// 原子操作更新帖子的评论数量
	if err := tx.Model(&Post{}).Where("id = ?", postID).UpdateColumn("comment_count", gorm.Expr("comment_count + 1")).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return comment, nil
}

// CancelCommentPost 在数据库中删除评论记录并原子更新帖子的评论数量
func (*Comment) CancelCommentPost(DB *gorm.DB, commentID, postID uint) error {
	// 获取评论记录
	var comment Comment
	if err := DB.Where("id = ? AND post_id = ?", commentID, postID).First(&comment).Error; err != nil {
		return err
	}

	// 原子操作删除评论记录并更新帖子的评论数量
	err := DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&comment).Error; err != nil {
			return err
		}

		// 查询帖子并更新评论数量
		var post Post
		if err := tx.First(&post, postID).Error; err != nil {
			return err
		}
		if post.CommentCount > 0 {
			post.CommentCount--
			if err := tx.Save(&post).Error; err != nil {
				return err
			}
		}
		return nil
	})

	return err
}
