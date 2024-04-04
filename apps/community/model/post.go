package model

import (
	"gorm.io/gorm"
)

// Post 帖子表模型
type Post struct {
	gorm.Model
	UserID    uint        `json:"user_id"`                         // 用户ID，JSON序列化时的字段名为"user_id"
	Content   string      `gorm:"type:text" json:"content"`        // 帖子内容，JSON序列化时的字段名为"content"
	Images    []PostImage `gorm:"foreignKey:PostID" json:"images"` // 帖子图片，外键关联到PostImage表的PostID字段，JSON序列化时的字段名为"images"
	Likes     []Like      `gorm:"foreignKey:PostID" json:"likes"`  // 帖子点赞，外键关联到Like表的PostID字段，JSON序列化时的字段名为"likes"
	LikeCount int64       `json:"like_count"`
}

// PostImage 帖子图片表模型
type PostImage struct {
	gorm.Model
	PostID   uint   `json:"post_id"`   // 帖子ID，JSON序列化时的字段名为"post_id"
	ImageUrl string `json:"image_url"` // 图片地址，JSON序列化时的字段名为"image_url"
}

// Like 点赞表模型
type Like struct {
	gorm.Model
	PostID uint `json:"post_id"` // 帖子ID，JSON序列化时的字段名为"post_id"
	UserID uint `json:"user_id"` // 用户ID，JSON序列化时的字段名为"user_id"
}

// Comment 评论表模型
type Comment struct {
	gorm.Model
	PostID  uint   `json:"post_id"` // 帖子ID，JSON序列化时的字段名为"post_id"
	UserID  uint   `json:"user_id"` // 用户ID，JSON序列化时的字段名为"user_id"
	Content string `json:"content"` // 评论内容，JSON序列化时的字段名为"content"
}
