package pkg

import (
	oldgorm "github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

// ConvertDB 将旧版 GORM 的 DB 转换为新版 GORM 的 DB
func ConvertDB(oldDB *oldgorm.DB) (*gorm.DB, error) {
	newDB, err := gorm.Open(oldDB.Dialect(), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return newDB, nil
}
