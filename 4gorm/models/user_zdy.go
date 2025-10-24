package models

import (
	"encoding/json"
)

// UserZdy 自定义用户模型，演示 JSON 序列化标签
type UserZdy struct {
	ID       uint     `json:"id" gorm:"primaryKey"`
	Username string   `json:"username" gorm:"size:50;not null"`
	Tags     []string `json:"tags" gorm:"serializer:json"`     // 使用 serializer:json 标签
	Settings Settings `json:"settings" gorm:"serializer:json"` // 复杂对象使用 serializer:json
}

// Settings 用户设置信息
type Settings struct {
	Theme    string `json:"theme"`
	Language string `json:"language"`
}

// ToJSON 将用户对象转换为 JSON 字符串
func (u *UserZdy) ToJSON() (string, error) {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

// FromJSON 从 JSON 字符串创建用户对象
func (u *UserZdy) FromJSON(jsonStr string) error {
	return json.Unmarshal([]byte(jsonStr), u)
}

