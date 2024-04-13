package model

import (
	"errors"
	"github.com/jinzhu/gorm"
)

// 用户表
type User struct {
	UserID           uint   `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Nickname         string `gorm:"not null" json:"nickname"`
	Account          string `gorm:"unique;not null" json:"account"`
	Email            string `gorm:"unique;default:null" json:"email"`
	Phone            string `gorm:"not null" json:"phone"`
	Password         string `gorm:"not null" json:"-"`
	RegistrationTime int64  `gorm:"autoCreateTime" json:"registration_time"`
	LastLoginTime    int64  `gorm:"autoUpdateTime" json:"last_login_time"`
	Status           string `gorm:"type:enum('Active', 'Inactive');default:'Active'" json:"status"`
	Role             string `gorm:"type:enum('User', 'Admin');default:'User'" json:"role"`
	BackgroundImage  string `gorm:"default:null" json:"background_image"`
	AvatarBackground string `gorm:"default:null" json:"avatar_background"`
	PostCount        int    `gorm:"default:0" json:"post_count"`
	FollowCount      int    `gorm:"default:0" json:"follow_count"`
	FansCount        int    `gorm:"default:0" json:"fans_count"`
	LikeCount        int    `gorm:"default:0" json:"like_count"`
	PointCount       int    `gorm:"default:0" json:"point_count"`
}

// 关注表
type Follow struct {
	FollowID       uint `gorm:"primaryKey;autoIncrement" json:"follow_id"`
	FollowerUserID uint `gorm:"index" json:"follower_user_id"`
	FollowedUserID uint `gorm:"index" json:"followed_user_id"`
}

// FindOneByEmail 根据邮箱查询用户
func (m *User) FindOneByEmail(db *gorm.DB, email string) (*User, error) {
	var user User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 返回 nil 表示找不到用户
		}
		return nil, err // 其他数据库错误
	}
	return &user, nil
}

// InsertUser 插入用户记录
func (m *User) InsertUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

// FindOneByAccount 根据账号查询用户
func (m *User) FindOneByAccount(db *gorm.DB, account string) (*User, error) {
	var user User
	if err := db.Where("account = ?", account).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 返回 nil 表示找不到用户
		}
		return nil, err // 其他数据库错误
	}
	return &user, nil
}

// FindOneByAccount 根据逐渐查询用户
func (m *User) FindOne(db *gorm.DB, id uint) (*User, error) {
	var user User
	if err := db.Where("user_id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 返回 nil 表示找不到用户
		}
		return nil, err // 其他数据库错误
	}
	return &user, nil
}

// UpdateUser 更新用户信息
func (m *User) UpdateUser(db *gorm.DB, id uint, updateData *User) error {
	if err := db.Model(&User{}).Where("user_id = ?", id).Updates(updateData).Error; err != nil {
		return err
	}
	return nil
}

// GetTopLikedUsers 获取获赞数前一千名的用户
func (m *User) GetTopLikedUsers(db *gorm.DB) ([]User, error) {
	var users []User
	result := db.Order("like_count DESC").Limit(1000).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// UpdatePointsGrab 每日抢积分没人的积分增加100
func (m *User) UpdatePointsGrab(db *gorm.DB, userID uint) error {
	// 更新指定用户的积分，增加100积分
	result := db.Model(&User{}).Where("user_id = ?", userID).Update("point_count", gorm.Expr("point_count + 100"))
	if result.Error != nil {
		return result.Error
	}
	return nil
}
