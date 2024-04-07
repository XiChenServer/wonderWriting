package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type CheckIn struct {
	gorm.Model
	UserID          uint      `gorm:"not null;index" json:"user_id"`
	ContinuousDays  int32     `gorm:"default:0" json:"continuous_days"`
	LastCheckInTime time.Time `gorm:"type:timestamp" json:"last_checkin_time"`
}

type RecordContent struct {
	gorm.Model
	UserID  uint    `gorm:"not null;index" json:"user_id"`
	Content string  `gorm:"type:text" json:"content"`
	Image   string  `gorm:"not null" json:"image"`
	Score   float64 `gorm:"type:decimal(10,2);default:0" json:"score"`
}

// 开启打卡记录
func (*CheckIn) CreateCheckIn(db *gorm.DB, userId uint) (*CheckIn, error) {
	checkIn := &CheckIn{
		UserID:          userId,
		ContinuousDays:  0,          // 设置默认值为0
		LastCheckInTime: time.Now(), // 设置为当前时间
	}
	err := db.Create(checkIn).Error
	if err != nil {
		return nil, err
	}
	return checkIn, nil
}

// 上传书法记录
func (*RecordContent) CreateRecordContent(db *gorm.DB, userId uint, content, image string, score float64) (*RecordContent, error) {
	recordContent := &RecordContent{
		UserID:  userId,
		Content: content,
		Image:   image,
		Score:   score,
	}
	if err := db.Create(recordContent).Error; err != nil {
		return nil, err
	}
	return recordContent, nil
}

// UpdateCheckInInfo 更新打卡信息
func (*CheckIn) UpdateCheckInInfo(db *gorm.DB, userId uint) error {
	// 查询用户的打卡记录
	checkInTime := time.Now()
	var checkIn CheckIn
	if err := db.Where("user_id = ?", userId).First(&checkIn).Error; err != nil {
		return err
	}

	// 判断是否中断连续签到
	if checkIn.LastCheckInTime.Day() != checkInTime.Day() {
		// 连续签到中断，重置连续签到天数为0
		checkIn.ContinuousDays = 0
	} else {
		// 未中断连续签到，连续签到天数加1
		checkIn.ContinuousDays++
	}
	checkIn.LastCheckInTime = checkInTime // 更新最后签到时间

	// 执行更新操作
	if err := db.Save(&checkIn).Error; err != nil {
		return err
	}

	return nil
}

// 查看某人的所有打卡记录
func (*RecordContent) LookAllRecordByOwn(db *gorm.DB, userId uint) ([]RecordContent, error) {
	recordContent := []RecordContent{}
	if err := db.Where("user_id = ?", userId).Find(&recordContent).Error; err != nil {
		return nil, err
	}
	return recordContent, nil
}
