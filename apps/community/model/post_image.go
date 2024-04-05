package model

import "github.com/jinzhu/gorm"

// PostImage 帖子图片表模型
type PostImage struct {
	gorm.Model
	PostID   uint   `json:"post_id"`   // 帖子ID，JSON序列化时的字段名为"post_id"
	ImageUrl string `json:"image_url"` // 图片地址，JSON序列化时的字段名为"image_url"
}
