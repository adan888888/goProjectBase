package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Record struct {
	ID    uint64 `gorm:"primary_key"`
	Name  string
	Money float64
}

type SingleTable struct {
	ID         uint64 `gorm:"primary_key"`
	Name       string `gorm:"unique"`
	Age        uint
	Adres      string
	CreateTime time.Time `gorm:"autoCreateTime"`
}

func (s *SingleTable) BeforeCreate(tx *gorm.DB) error {
	fmt.Println("创建的钩子函数")
	return nil
}

// BeforeUpdate 钩子函数
func (u *SingleTable) BeforeUpdate(tx *gorm.DB) error {
	fmt.Println("更新的钩子函数")
	u.CreateTime = time.Now()
	return nil
}
