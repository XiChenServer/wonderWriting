package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type CheckIn struct {
	gorm.Model
	UserID          uint            `gorm:"not null" json:"user_id"`
	ContinuousDays  int             `gorm:"default:0" json:"continuous_days"`
	Record          []RecordContent `gorm:"foreignKey:CheckInID" json:"record"`
	LastCheckInTime int64           `gorm:"autoUpdateTime" json:"last_checkin_time"`
	CheckInDate     time.Time       `gorm:"not null" json:"checkin_date"`
}

type RecordContent struct {
	gorm.Model
	CheckInID uint   `gorm:"not null;index" json:"check_in_id"`
	Content   string `gorm:"not null" json:"content"`
	Image     string `gorm:"not null" json:"image"`
}
