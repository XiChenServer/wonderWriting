package model

import (
	"calligraphy/apps/user/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

// Comment 评论表模型
type Comment struct {
	gorm.Model
	PostID       uint   `json:"post_id"`                  // 帖子ID，JSON序列化时的字段名为"post_id"
	Post         Post   `gorm:"foreignKey:PostID"`        // 关联的帖子，使用PostID作为外键
	UserID       uint   `json:"user_id"`                  // 用户ID，JSON序列化时的字段名为"user_id"
	Content      string `gorm:"type:text" json:"content"` // 评论内容，JSON序列化时的字段名为"content"，类型为TEXT
	UserAvatar   string `json:"user_avatar"`              // 回复者的头像
	UserNickName string `json:"user_nickname"`            // 回复者的昵称
	LikeCount    uint   `json:"like_count"`               // 点赞数量
}

// ReplyComment 回复评论
type ReplyComment struct {
	gorm.Model
	CommentID     uint   `json:"comment_id"`               // 回复的评论的ID
	UserID        uint   `json:"user_id"`                  // 回复者的ID
	UserNickName  string `json:"user_nickname"`            // 回复者的昵称
	UserAvatar    string `json:"user_avatar"`              // 回复者的头像
	Content       string `gorm:"type:text" json:"content"` // 回复的内容，类型为TEXT
	ReplyNickName string `json:"reply_nickname"`           // 给谁回复，那个人的昵称
	ReplyUserId   uint   `json:"reply_user_id"`            // 给谁回复，那个人的ID
	LikeCount     uint   `json:"like_count"`               // 点赞数量
}

// LikeComment 对评论点赞
type LikeComment struct {
	gorm.Model
	CommentID uint `json:"comment_id"` // 评论的ID
	UserID    uint `json:"user_id"`    // 点赞用户的ID
}

// ReplyComment 回复评论
func (*ReplyComment) ReplyComment(DB *gorm.DB, commentID, userID, replyUserID, postID uint, replyNickName, content string) (*ReplyComment, error) {
	// 开启事务
	tx := DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() // 回滚事务
			log.Printf("recovered from panic: %v", r)
		} else {
			tx.Commit() // 提交事务
		}
	}()

	// 创建回复评论记录
	var user model.User
	if err := tx.Where("user_id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	replyComment := &ReplyComment{
		CommentID:     commentID,
		UserID:        userID,
		UserNickName:  user.Nickname,
		UserAvatar:    user.AvatarBackground,
		Content:       content,
		ReplyNickName: replyNickName,
		ReplyUserId:   replyUserID,
		LikeCount:     0,
	}
	if err := tx.Create(replyComment).Error; err != nil {
		return nil, err
	}

	// 原子操作更新帖子的评论数量
	if err := tx.Model(&Post{}).Where("id = ?", postID).UpdateColumn("comment_count", gorm.Expr("comment_count + 1")).Error; err != nil {
		return nil, err
	}

	return replyComment, nil
}

// LikeComment 点赞评论
func (lc *LikeComment) LikeComment(DB *gorm.DB, commentID, userID uint) error {
	// 开启事务
	tx := DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() // 回滚事务
			log.Printf("recovered from panic: %v", r)
		} else {
			tx.Commit() // 提交事务
		}
	}()

	// 检查用户是否已经点赞过该评论
	var existingLike LikeComment
	if err := tx.Where("comment_id = ? AND user_id = ?", commentID, userID).First(&existingLike).Error; err != nil {
		// 如果未点赞过，则创建点赞记录
		newLike := LikeComment{
			CommentID: commentID,
			UserID:    userID,
		}
		if err := tx.Create(&newLike).Error; err != nil {
			return err
		}
		// 更新评论的点赞数量
		if err := IncrementCommentLikeCount(tx, commentID); err != nil {
			return err
		}
	} else {
		// 如果已经点赞过，则取消点赞并更新点赞数量
		if err := tx.Delete(&existingLike).Error; err != nil {
			return err
		}
		if err := DecrementCommentLikeCount(tx, commentID); err != nil {
			return err
		}
	}

	return nil
}

// IncrementCommentLikeCount 增加评论的点赞数量，使用事务保证原子性
func IncrementCommentLikeCount(DB *gorm.DB, commentID uint) error {
	// 查询评论
	var comment Comment
	if err := DB.Where("id = ?", commentID).First(&comment).Error; err != nil {
		return err
	}

	// 更新点赞数量
	if err := DB.Model(&Comment{}).Where("id = ?", commentID).Update("like_count", gorm.Expr("like_count + 1")).Error; err != nil {
		return err
	}

	return nil
}

// DecrementCommentLikeCount 减少评论的点赞数量，使用事务保证原子性
func DecrementCommentLikeCount(DB *gorm.DB, commentID uint) error {
	// 查询评论
	var comment Comment
	if err := DB.Where("id = ?", commentID).First(&comment).Error; err != nil {
		return err
	}

	// 更新点赞数量
	if err := DB.Model(&Comment{}).Where("id = ?", commentID).Update("like_count", gorm.Expr("like_count - 1")).Error; err != nil {
		return err
	}

	return nil
}

// FindComment 查看评论
func (*Comment) FindComment(DB *gorm.DB, postId uint) (*[]Comment, error) {
	var comments []Comment
	if err := DB.Where("post_id = ?", postId).Find(&comments).Error; err != nil {
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

// FindReplyCommentCount 查询回复评论的总记录数
func (m *ReplyComment) FindReplyCommentCount(db *gorm.DB, commentID, userID uint) (int64, error) {
	var count int64
	err := db.Model(&ReplyComment{}).Where("comment_id = ? AND user_id = ?", commentID, userID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// FindReplyCommentsByPage 分页查询回复评论信息
func (m *ReplyComment) FindReplyCommentsByPage(db *gorm.DB, commentID, userID, page, pageSize uint) (*[]ReplyComment, error) {
	var res []ReplyComment
	err := db.Where("comment_id = ? AND user_id = ?", commentID, userID).Offset((page - 1) * pageSize).Limit(pageSize).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return &res, nil
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
