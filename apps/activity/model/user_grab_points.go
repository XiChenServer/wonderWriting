package model

import "github.com/jinzhu/gorm"

type UserPoints struct {
	gorm.Model
	UID     int64 `gorm:"unique" json:"uid"`
	Claimed bool  `json:"claimed"`
}

func (s *UserPoints) CheckUserClaimed(db *gorm.DB, userID uint) (bool, error) {
	var user UserPoints
	if err := db.Where("uid = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// 用户不存在，未领取过积分
			return false, nil
		}
		return false, err
	}
	// 查询到用户，判断是否已领取
	return user.Claimed, nil
}
