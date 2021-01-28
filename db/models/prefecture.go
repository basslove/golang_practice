package models

import "time"

type Prefecture struct {
	ID   int64 `gorm:"primaryKey"`
	Name string `gorm:"column:name"`
	IsActive bool `gorm:"column:is_active"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Prefecture) TableName() string {
	return "prefectures"
}


