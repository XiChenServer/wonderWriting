package model

import "github.com/jinzhu/gorm"

// Post 帖子表模型
type Post struct {
	gorm.Model
	UserID    uint        `json:"user_id"`                         // 用户ID，JSON序列化时的字段名为"user_id"
	Content   string      `gorm:"type:text" json:"content"`        // 帖子内容，JSON序列化时的字段名为"content"
	Images    []PostImage `gorm:"foreignKey:PostID" json:"images"` // 帖子图片，外键关联到PostImage表的PostID字段，JSON序列化时的字段名为"images"
	Likes     []Like      `gorm:"foreignKey:PostID" json:"likes"`  // 帖子点赞，外键关联到Like表的PostID字段，JSON序列化时的字段名为"likes"
	LikeCount int64       `json:"like_count"`
}
