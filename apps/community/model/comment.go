package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

// Comment 评论表模型
type Comment struct {
	gorm.Model
	CommentReplays []*Comment `gorm:"foreignkey:ParentID"` //自引用
	PostID         uint       `json:"post_id"`             // 帖子ID，JSON序列化时的字段名为"post_id"
	Post           Post       `gorm:"foreignKey:PostID"`   // 关联的帖子，使用PostID作为外键
	UserID         uint       `json:"user_id"`             // 用户ID，JSON序列化时的字段名为"user_id"
	Content        string     `json:"content"`             // 评论内容，JSON序列化时的字段名为"content"
	Parent         *Comment   `gorm:"foreignkey:ParentID"` //自引用
	ParentID       uint       `json:"parent_id"`           //评论父级ID  0

}

func (*Comment) FindComment(DB *gorm.DB, post_id uint) (*[]Comment, error) {
	var comments []Comment
	if err := DB.Where("post_id = ?", post_id).Find(&comments).Error; err != nil {
		return nil, err
	}
	return &comments, nil
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
	fmt.Println(userID)
	fmt.Println(comment)
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

// FindCommentsByPage 查询评论信息分页
func (*Comment) FindCommentsByPage(DB *gorm.DB, page, pageSize, postID uint) (*[]Comment, error) {
	var res []Comment

	offset := (page - 1) * pageSize
	err := DB.Where("post_id = ?", postID).Offset(offset).Limit(pageSize).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// FindCommentCount 查询评论的总数
func (*Comment) FindCommentCount(DB *gorm.DB, postId uint) (int64, error) {
	var totalCount int64
	err := DB.Model(Comment{}).Where("post_id = ?", postId).Count(&totalCount).Error
	if err != nil {
		return 0, err
	}
	return totalCount, nil
}

//// CommentPost 在数据库中创建一条评论记录并原子更新帖子的评论数量
//func (*Comment) CommentPost(DB *gorm.DB, postID, userID, parentID uint, content string) (*Comment, error) {
//	// 查询帖子是否存在
//	var post Post
//	if err := DB.First(&post, postID).Error; err != nil {
//		return nil, err
//	}
//
//	// 开始事务
//	tx := DB.Begin()
//
//	// 创建评论记录
//	comment := &Comment{
//		PostID:   postID,
//		UserID:   userID,
//		ParentID: parentID, // 设置父级评论ID
//		Content:  content,
//	}
//	if err := tx.Create(comment).Error; err != nil {
//		tx.Rollback()
//		return nil, err
//	}
//
//	// 原子操作更新帖子的评论数量
//	if err := tx.Model(&Post{}).Where("id = ?", postID).Update("comment_count", post.CommentCount+1).Error; err != nil {
//		tx.Rollback()
//		return nil, err
//	}
//
//	// 提交事务
//	if err := tx.Commit().Error; err != nil {
//		tx.Rollback()
//		return nil, err
//	}
//
//	return comment, nil
//}
