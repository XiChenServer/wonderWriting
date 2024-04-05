package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Model struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// Comment 评论表模型
type Comment struct {
	gorm.Model
	PostID  uint   `json:"post_id"`           // 帖子ID，JSON序列化时的字段名为"post_id"
	Post    Post   `gorm:"foreignKey:PostID"` // 关联的帖子，使用PostID作为外键
	UserID  uint   `json:"user_id"`           // 用户ID，JSON序列化时的字段名为"user_id"
	Content string `json:"content"`           // 评论内容，JSON序列化时的字段名为"content"
}
