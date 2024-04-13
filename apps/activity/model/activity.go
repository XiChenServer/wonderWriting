package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

// Activity 活动表结构
type Activity struct {
	gorm.Model
	Name         string               `json:"name" gorm:"column:name"`                              // 活动名称
	ActivityInfo string               `json:"activity_info" gorm:"column:activity_info;type:TEXT"`  // 活动信息
	Location     string               `json:"location" gorm:"column:location"`                      // 活动地点
	DateTime     string               `json:"dateTime" gorm:"column:date_time"`                     // 活动日期和时间
	Organizer    string               `json:"organizer" gorm:"column:organizer"`                    // 组织者
	EndDateTime  string               `json:"endDateTime" gorm:"column:end_date_time"`              // 活动结束时间
	Duration     string               `json:"duration" gorm:"column:duration"`                      // 活动时长
	RewardsInfo  string               `json:"rewardsInfo" gorm:"column:rewards_info"`               // 奖励信息记录为文本
	Participants []UserSignUpActivity `json:"participants" gorm:"many2many:activity_participants;"` // 参与者
}

// User 用户表结构
type UserSignUpActivity struct {
	gorm.Model
	UserId      uint   `json:"userId" gorm:"column:user_id"`
	Name        string `json:"name" gorm:"column:name"`
	ContactInfo string `json:"contactInfo" gorm:"column:contact_info"`
	Experience  bool   `json:"experience" gorm:"column:experience"`
	ActivityID  uint   `json:"activityID" gorm:"column:activity_id"`
}

func (*Activity) GetAllActivities(db *gorm.DB, start, pageSize int) ([]Activity, error) {
	var activities []Activity
	// 查询数据库，获取当前页的活动数据，并按照创建时间降序排序
	err := db.Order("created_at desc").Offset(start).Limit(pageSize).Find(&activities).Error
	if err != nil {
		return nil, err
	}
	return activities, nil
}

// GetActivityInfo 根据活动ID获取活动信息
func (*Activity) GetActivityInfo(db *gorm.DB, id uint) (*Activity, error) {
	var activity Activity
	fmt.Println("123", id)
	// 根据活动ID查询数据库中的活动信息
	err := db.Where("id = ?", id).First(&activity).Error
	if err != nil {
		fmt.Println("123", id, err.Error())
		return nil, err
	}
	return &activity, nil
}

// CheckUserSignUp 检查用户是否已经报名活动
func (*UserSignUpActivity) CheckUserSignUp(db *gorm.DB, userID, activityID uint) (bool, error) {
	var count int
	err := db.Model(&UserSignUpActivity{}).Where("user_id = ? AND activity_id = ?", userID, activityID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// CreateSignUpRecord 创建报名记录
func (*UserSignUpActivity) CreateSignUpRecord(db *gorm.DB, userID, activityID uint) error {
	// 创建报名记录，将用户ID和活动ID关联起来
	signUpRecord := &UserSignUpActivity{
		UserId:     userID,
		ActivityID: activityID,
	}
	err := db.Create(signUpRecord).Error
	if err != nil {
		return err
	}
	return nil
}

// GetUserActivities 获取用户参加的活动ID列表
func (*UserSignUpActivity) GetUserActivities(db *gorm.DB, userID uint, start, pageSize int) ([]uint, error) {
	// 查询用户参加的活动ID列表，假设有一个 UserActivity 结构体表示用户活动关联关系
	var userActivities []UserSignUpActivity
	err := db.Where("user_id = ?", userID).Offset(start).Limit(pageSize).Find(&userActivities).Error
	if err != nil {
		return nil, err
	}

	// 提取用户参加的活动ID
	activityIDs := make([]uint, len(userActivities))
	for i, ua := range userActivities {
		activityIDs[i] = ua.ActivityID
	}

	return activityIDs, nil
}

// GetUserActivitiesCount 获取用户参加的活动总数
func (*UserSignUpActivity) GetUserActivitiesCount(db *gorm.DB, userID uint) (int, error) {
	var count int
	// 查询用户参加的活动总数
	err := db.Model(&UserSignUpActivity{}).Where("user_id = ?", userID).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
