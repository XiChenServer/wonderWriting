package model

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

// User 用户表
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

// Follow 关注表
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

// FindOne 根据逐渐查询用户
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

// FindOneByFollowerAndFollowed 根据关注者和被关注者查询关注记录
func (m *Follow) FindOneByFollowerAndFollowed(db *gorm.DB, followerUserID, followedUserID uint) (*Follow, error) {
	var follow Follow
	if err := db.Where("follower_user_id = ? AND followed_user_id = ?", followerUserID, followedUserID).First(&follow).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 没有找到关注记录
		}
		return nil, err // 其他错误
	}
	return &follow, nil
}

// CreateUserFollow 创建用户关注记录
func (m *Follow) CreateUserFollow(db *gorm.DB, newFollow *Follow) error {
	return db.Create(newFollow).Error
}

// IncrementFollowCount 增加用户关注数
func (m *Follow) IncrementFollowCount(db *gorm.DB, userID uint) error {
	return db.Model(&User{}).Where("user_id = ?", userID).UpdateColumn("follow_count", gorm.Expr("follow_count + 1")).Error
}

// IncrementFansCount 增加用户粉丝数
func (m *Follow) IncrementFansCount(db *gorm.DB, userID uint) error {
	return db.Model(&User{}).Where("user_id = ?", userID).UpdateColumn("fans_count", gorm.Expr("fans_count + 1")).Error
}

// DeleteFollow 删除关注记录
func (m *Follow) DeleteFollow(db *gorm.DB) error {
	return db.Delete(m).Error
}

// DecrementFollowCount 减少用户关注数
func (m *Follow) DecrementFollowCount(db *gorm.DB, userID uint) error {
	return db.Model(&User{}).Where("user_id = ?", userID).UpdateColumn("follow_count", gorm.Expr("follow_count - 1")).Error
}

// DecrementFansCount 减少用户粉丝数
func (m *Follow) DecrementFansCount(db *gorm.DB, userID uint) error {
	return db.Model(&User{}).Where("user_id = ?", userID).UpdateColumn("fans_count", gorm.Expr("fans_count - 1")).Error
}

// LookAllFollow 查找自己的关注人
func (m *Follow) LookAllFollow(db *gorm.DB, userID uint) (*[]Follow, error) {
	var follows []Follow
	err := db.Where("follower_user_id = ?", userID).Find(&follows).Error
	if err != nil {
		return nil, err
	}
	return &follows, nil
}

// LookAllFans 查找自己的粉丝
func (m *Follow) LookAllFans(db *gorm.DB, userID uint) (*[]Follow, error) {
	var fans []Follow
	err := db.Where("followed_user_id = ?", userID).Find(&fans).Error
	if err != nil {
		return nil, err
	}
	return &fans, nil
}

// WhetherLikedPost 用户是否关注一个人
func (*Follow) WhetherLikedPost(DB *gorm.DB, otherId, userID uint) error {
	var like Follow
	err := DB.Where("followed_user_id = ? AND follower_user_id = ?", otherId, userID).First(&like).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("用户未关注该用户")
		}
		return err
	}
	return nil
}
