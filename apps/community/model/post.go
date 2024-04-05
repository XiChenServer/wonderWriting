package model

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	UserID          uint        `json:"user_id"`                              // 用户ID，JSON序列化时的字段名为"user_id"
	Content         string      `gorm:"type:text" json:"content"`             // 帖子内容，JSON序列化时的字段名为"content"
	Images          []PostImage `gorm:"foreignKey:PostID" json:"images"`      // 帖子图片，外键关联到PostImage表的PostID字段，JSON序列化时的字段名为"images"
	Likes           []Like      `gorm:"foreignKey:PostID" json:"likes"`       // 帖子点赞，外键关联到Like表的PostID字段，JSON序列化时的字段名为"likes"
	LikeCount       uint        `json:"like_count"`                           // 帖子点赞数量
	CollectionCount uint        `json:"collection_count"`                     // 帖子收藏数量
	Collections     []Collect   `gorm:"foreignKey:PostID" json:"collections"` // 帖子收藏，多对多关联，使用中间表post_collections
	CommentCount    uint        `gorm:"-" json:"comment_count"`               // 帖子评论数量
	Comments        []Comment   `gorm:"foreignKey:PostID" json:"comment"`     // 帖子评论，外键关联到Comment表的PostID字段，JSON序列化时的字段名为"comment"
}

// CreatePost 创建帖子
func (*Post) CreatePost(dao *gorm.DB, userId uint, content string, urls []string) (*Post, error) {
	// 首先创建帖子
	newPost := &Post{
		UserID:  userId,
		Content: content,
	}
	if err := dao.Create(newPost).Error; err != nil {
		return nil, err
	}

	// 然后循环创建帖子图片
	for _, v := range urls {
		newImage := &PostImage{
			PostID:   newPost.Model.ID,
			ImageUrl: v,
		}
		if err := dao.Create(newImage).Error; err != nil {
			return nil, err
		}
	}

	return newPost, nil
}

// DeletePost 删除帖子
func (*Post) DeletePost(dao *gorm.DB, post_id uint32) (*Post, error) {
	err := dao.Where("post_id = ?", post_id).Delete(&PostImage{}).Error
	if err != nil {
		return nil, err
	}
	err = dao.Where("id = ?", post_id).Delete(&Post{}).Error
	if err != nil {
		return nil, err
	}
	err = dao.Where("post_id = ?", post_id).Delete(&Comment{}).Error
	if err != nil {
		return nil, err
	}
	err = dao.Where("post_id = ?", post_id).Delete(&Like{}).Error
	if err != nil {
		return nil, err
	}
	err = dao.Where("post_id = ?", post_id).Delete(&Collect{}).Error
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// LookAllPosts 查找多有的帖子
func (*Post) LookAllPosts(dao *gorm.DB) ([]*Post, error) {
	var posts []*Post
	err := dao.Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

// LookPostByOwn 查看自己的帖子
func (*Post) LookPostByOwn(dao *gorm.DB, userId uint) ([]*Post, error) {
	var posts []*Post
	err := dao.Where("user_id = ?", userId).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}
