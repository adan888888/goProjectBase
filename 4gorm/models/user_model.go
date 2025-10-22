package models

import (
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	ID              uint             `gorm:"primaryKey"`
	Name            string           `gorm:"size:100;not null"`
	Age             int              `gorm:"not null;default:18"`
	UserDetailModel *UserDetailModel `gorm:"foreignKey:UserID"` //反向引用
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
type UserDetailModel struct {
	ID        int
	UserID    int       `gorm:"unique"` //一对一的情况下，需要加上唯一约束
	Email     string    `gorm:"size:32"`
	UserModel UserModel `gorm:"foreignKey:UserID"`
}
