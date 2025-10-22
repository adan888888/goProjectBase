package models

import (
	"time"

	"gorm.io/gorm"
)

// StudentModel 学生模型
// 一个学生可以选修多门课程，一门课程可以被多个学生选修
type StudentModel struct {
	ID        uint          `gorm:"primaryKey"`
	StudentNo string        `gorm:"size:20;unique;not null"`    // 学号
	Name      string        `gorm:"size:50;not null"`           // 姓名
	Major     string        `gorm:"size:50"`                    // 专业
	Courses   []CourseModel `gorm:"many2many:student_courses;"` // 多对多关联：学生课程
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// CourseModel 课程模型
// 一门课程可以被多个学生选修，一个学生可以选修多门课程
type CourseModel struct {
	ID        uint           `gorm:"primaryKey"`
	CourseNo  string         `gorm:"size:20;unique;not null"`    // 课程编号
	Name      string         `gorm:"size:100;not null"`          // 课程名称
	Teacher   string         `gorm:"size:50"`                    // 授课教师
	Students  []StudentModel `gorm:"many2many:student_courses;"` // 多对多关联：课程学生
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

//这两个的标签要一样 自动会创建第三张关联表（中间表）
//Students  []StudentModel  `gorm:"many2many:student_courses;"`
//Courses   []CourseModel  `gorm:"many2many:student_courses;"`
