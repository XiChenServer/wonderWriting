package model

import (
	"calligraphy/pkg/qiniu"
	"github.com/jinzhu/gorm"
)

// PostImage 帖子图片表模型
type PostImage struct {
	gorm.Model
	PostID   uint   `json:"post_id"`   // 帖子ID，JSON序列化时的字段名为"post_id"
	ImageUrl string `json:"image_url"` // 图片地址，JSON序列化时的字段名为"image_url"
}

func (*PostImage) FindImageByPostId(DB *gorm.DB, post_id uint) ([]string, error) {
	var image []PostImage

	err := DB.Where("post_id = ?", post_id).Find(&image).Error
	if err != nil {
		return nil, err
	}

	var urls []string
	for _, v := range image {
		urls = append(urls, qiniu.ImgUrl+v.ImageUrl)
	}
	return urls, err

}
