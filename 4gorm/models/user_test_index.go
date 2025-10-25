package models

import (
	"time"

	"gorm.io/gorm"
)

// UserTestIndex 用于测试索引性能的用户表
type UserTestIndex struct {
	ID        uint           `gorm:"primaryKey"`
	Username  string         `gorm:"size:50;not null"`                 // 用户名
	Email     string         `gorm:"size:100;not null"`                // 邮箱
	Age       int            `gorm:"not null"`                         // 年龄
	City      string         `gorm:"size:50;not null"`                 // 城市
	Phone     string         `gorm:"size:20;not null"`                 // 电话
	Address   string         `gorm:"size:200"`                          // 地址
	Status    int            `gorm:"not null;default:1"`               // 状态
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// TableName 自定义表名
func (UserTestIndex) TableName() string {
	return "user_test_index"
}
