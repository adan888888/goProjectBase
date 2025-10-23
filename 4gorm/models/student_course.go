package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

// 简化的多对多关系模型，专注于学习核心概念

// StudentModel 学生模型
// 一个学生可以选修多门课程，一门课程可以被多个学生选修
// joinForeignKey:StudentID 当前的;JoinReferences:CourseID 对方的
type StudentModel struct {
	ID      uint          `gorm:"primaryKey"`
	Name    string        `gorm:"size:50;not null"`                                                           // 姓名
	Courses []CourseModel `gorm:"many2many:student_courses;joinForeignKey:StudentID;joinReferences:CourseID"` // 多对多关联：学生课程 ✅✅✅会自动创建外键关联
}

// CourseModel 课程模型
// 一门课程可以被多个学生选修，一个学生可以选修多门课程
type CourseModel struct {
	ID       uint           `gorm:"primaryKey"`
	Name     string         `gorm:"size:100;not null"`                                                          // 课程名称
	Students []StudentModel `gorm:"many2many:student_courses;joinForeignKey:CourseID;joinReferences:StudentID"` // 多对多关联：课程学生 ✅✅✅会自动创建外键关联
}

// 重要说明：
// 1. 使用自定义中间表时，需要确保在数据库迁移时包含 StudentCourse 模型
// 2. 在 AutoMigrate 中需要添加 &StudentCourse{}
// 3. 多对多关系会自动使用自定义的中间表结构

// StudentCourse 自定义中间表（更加灵活）
// ✅✅✅会自动创建外键关联
type StudentCourse struct {
	ID           uint           `gorm:"primaryKey"` // 主键ID
	StudentID    uint           `gorm:"not null"`   // 学生ID
	StudentModel StudentModel   `gorm:"foreignKey:StudentID"`
	CourseID     uint           `gorm:"not null"` // 课程ID
	CourseModel  CourseModel    `gorm:"foreignKey:CourseID"`
	DeletedAt    gorm.DeletedAt `gorm:"index"` // 软删除时间字段
	CreatedAt    time.Time      `gorm:"created_at"`
	Title        string         `gorm:"size:32" json:"title"` //假如自己想再加个字段
}

func (u *StudentCourse) BeforeCreate(tx *gorm.DB) error {
	var title string
	tx.Model(&CourseModel{}).Where("id = ?", u.CourseID).Select("name").Scan(&title) //‼️这里比较复杂需要查了数据库自己赋值过去
	u.Title = title
	return nil
}

// TableName 自定义表名
func (*StudentCourse) TableName() string {
	fmt.Println("自定义表名")
	return "student_courses"
}
