package model

import "github.com/jinzhu/gorm"

// Comment 评论表模型
type Comment struct {
	gorm.Model
	PostID  uint   `json:"post_id"` // 帖子ID，JSON序列化时的字段名为"post_id"
	UserID  uint   `json:"user_id"` // 用户ID，JSON序列化时的字段名为"user_id"
	Content string `json:"content"` // 评论内容，JSON序列化时的字段名为"content"
}
