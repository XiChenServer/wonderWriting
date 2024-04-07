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
