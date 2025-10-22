package models

import (
	"time"

	"gorm.io/gorm"
)

// DepartmentModel 演示一对多关系
// 一个部门可以有多个员工，一个员工只能属于一个部门。    员工(多)---部门(一)
type DepartmentModel struct {
	ID           uint            `gorm:"primaryKey"`
	DeptNo       string          `gorm:"size:20;unique;not null"`              // 部门编号 部门编号通常比ID更稳定，不会因为数据迁移而改变
	Name         string          `gorm:"size:50;not null"`                     // 部门名称
	Code         string          `gorm:"size:20;unique;not null"`              // 部门编码
	Manager      string          `gorm:"size:30"`                              // 部门经理
	Location     string          `gorm:"size:100"`                             // 部门位置
	EmployeeList []EmployeeModel `gorm:"foreignKey:DeptNoX;references:DeptNo"` // 一对多关联，使用字符串外键。foreignKey:他人;references:自己
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"` // grom框架只要有DeletedAt这就是软删除字段
}

type EmployeeModel struct {
	ID         uint            `gorm:"primaryKey"`
	Name       string          `gorm:"size:30;not null"`                     // 员工姓名
	EmployeeNo string          `gorm:"size:20;unique;not null"`              // 工号
	Position   string          `gorm:"size:30"`                              // 职位
	Salary     float64         `gorm:"type:decimal(10,2)"`                   // 薪资
	Email      string          `gorm:"size:50"`                              // 邮箱
	Phone      string          `gorm:"size:20"`                              // 电话
	DeptNoX    string          `gorm:"size:20;not null"`                     // 部门编号（外键） ❌❌❌一定不能和一中的字体名字一样，要不然死活不能创建成功，要不然就删除Department DepartmentModel `gorm:"foreignKey:DeptNoX;references:DeptNo"`
	Department DepartmentModel `gorm:"foreignKey:DeptNoX;references:DeptNo"` // 多对一关联，使用字符串外键 foreignKey:自己;references:他人
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
