package model

import "gorm.io/gorm"

// 迁移函数中添加评论表
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Post{}, &PostImage{}, &Like{}, &Comment{})
}
