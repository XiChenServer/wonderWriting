package model // Like 点赞表模型
import "github.com/jinzhu/gorm"

type Like struct {
	gorm.Model
	PostID uint `json:"post_id"` // 帖子ID，JSON序列化时的字段名为"post_id"
	UserID uint `json:"user_id"` // 用户ID，JSON序列化时的字段名为"user_id"
}
