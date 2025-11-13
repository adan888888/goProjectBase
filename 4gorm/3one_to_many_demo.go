package main

import (
	"4gorm/global"
	"4gorm/models"
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	// 初始化数据库连接
	global.InitDB()

	// 自动迁移数据库表
	global.MigrateDB()

	fmt.Println("\n=== GORM 部门员工一对多关系演示 ===")

	// 1. 创建部门
	//demonstrateCreateDepartments()

	// 2. 创建员工
	//demonstrateCreateEmployees()

	// 3. 查询所有员工及其部门信息--员工表为主
	getAllEmployeesWithDepartments()

	// 4. 查询所有部门及其员工（预加载）
	//demonstrateGetAllDepartmentsWithEmployees()

	// 5. 查询特定部门的员工
	//demonstrateGetTechDepartmentEmployees()

	//// 5. 查询员工及其部门信息
	//demonstrateGetEmployeesWithDepartments()
	//
	//// 6package_manager. 更新员工薪资
	//demonstrateUpdateEmployeeSalary()
	//
	//// 7. 添加新员工到现有部门
	//demonstrateAddNewEmployeeToDepartment()
	//
	//// 8. 再次查询技术部员工
	//demonstrateGetTechDepartmentEmployeesAgain()
	//
	//// 9. 统计各部门员工数量
	//demonstrateCountEmployeesByDepartment()
	//
	//// 10. 删除员工
	//demonstrateDeleteEmployee()
	//
	//// 11. 最终统计
	//demonstrateFinalStatistics()

	fmt.Println("\n=== 部门员工一对多关系演示完成 ===")
}

// 演示创建部门
func demonstrateCreateDepartments() {
	fmt.Println("\n1. 创建部门:")
	departments := []models.DepartmentModel{
		{
			DeptNo:   "DEPT001",
			Name:     "技术部",
			Code:     "TECH001",
			Manager:  "张技术",
			Location: "北京总部3楼",
		},
		{
			DeptNo:   "DEPT002",
			Name:     "销售部",
			Code:     "SALES001",
			Manager:  "李销售",
			Location: "北京总部2楼",
		},
		{
			DeptNo:   "DEPT003",
			Name:     "人事部",
			Code:     "HR001",
			Manager:  "王人事",
			Location: "北京总部1楼",
		},
	}

	for _, dept := range departments {
		result := global.DB.Create(&dept)
		if result.Error != nil {
			fmt.Printf("创建部门失败: %v\n", result.Error)
			continue
		}
		fmt.Printf("创建部门成功! ID: %d, 编号: %s, 名称: %s, 编码: %s, 经理: %s\n",
			dept.ID, dept.DeptNo, dept.Name, dept.Code, dept.Manager)
	}
}

// 演示创建员工
func demonstrateCreateEmployees() {
	fmt.Println("\n2. 创建员工:")
	employees := []models.EmployeeModel{
		{
			Name:       "张三",
			EmployeeNo: "E001",
			Position:   "高级工程师",
			Salary:     15000.00,
			Email:      "zhangsan@company.com",
			Phone:      "13800138001",
			DeptNoX:    "DEPT001", // 技术部
		},
		{
			Name:       "李四",
			EmployeeNo: "E002",
			Position:   "初级工程师",
			Salary:     8000.00,
			Email:      "lisi@company.com",
			Phone:      "13800138002",
			DeptNoX:    "DEPT001", // 技术部
		},
		{
			Name:       "王五",
			EmployeeNo: "E003",
			Position:   "销售经理",
			Salary:     12000.00,
			Email:      "wangwu@company.com",
			Phone:      "13800138003",
			DeptNoX:    "DEPT002", // 销售部
		},
		{
			Name:       "赵六",
			EmployeeNo: "E004",
			Position:   "销售专员",
			Salary:     6000.00,
			Email:      "zhaoliu@company.com",
			Phone:      "13800138004",
			DeptNoX:    "DEPT002", // 销售部
		},
		{
			Name:       "钱七",
			EmployeeNo: "E005",
			Position:   "HR专员",
			Salary:     7000.00,
			Email:      "qianqi@company.com",
			Phone:      "13800138005",
			DeptNoX:    "DEPT003", // 人事部
		},
	}
	for _, emp := range employees {
		result := global.DB.Create(&emp)
		if result.Error != nil {
			fmt.Printf("创建员工失败: %v\n", result.Error)
			continue
		}
		fmt.Printf("创建员工成功! ID: %d, 姓名: %s, 工号: %s, 职位: %s, 薪资: %.2f\n",
			emp.ID, emp.Name, emp.EmployeeNo, emp.Position, emp.Salary)
	}
}

// 演示查询所有部门及其员工（预加载）
func demonstrateGetAllDepartmentsWithEmployees() {
	fmt.Println("\n3. 查询所有部门及其员工:")
	var deptList []models.DepartmentModel
	result := global.DB.Unscoped().Preload("EmployeeList").Find(&deptList)
	if result.Error != nil {
		fmt.Printf("查询部门失败: %v\n", result.Error)
		return
	}
	// 将对象转换为JSON字符串
	jsonData, err := json.MarshalIndent(deptList, "", "  ")
	if err != nil {
		fmt.Printf("JSON转换失败: %v\n", err)
	} else {
		fmt.Printf("部门信息:\n%s\n", string(jsonData))
	}
}

// 查询所有员工及其部门信息（新方法）
func getAllEmployeesWithDepartments() {
	fmt.Println("\n=== 查询所有员工及其部门信息 ===")
	var empList []models.EmployeeModel
	result := global.DB. /*Unscoped().会把删除的（软删除）也查询出来*/ Preload("Department").Find(&empList)
	if result.Error != nil {
		fmt.Printf("查询员工失败: %v\n", result.Error)
		return
	}

	fmt.Printf("查询到 %d 个员工:\n", len(empList))
	for i, emp := range empList {
		fmt.Printf("\n员工 %d:\n", i+1)
		fmt.Printf("  - 基本信息: ID=%d, 姓名=%s, 工号=%s\n", emp.ID, emp.Name, emp.EmployeeNo)
		fmt.Printf("  - 职位信息: 职位=%s, 薪资=%.2f\n", emp.Position, emp.Salary)
		fmt.Printf("  - 联系方式: 邮箱=%s, 电话=%s\n", emp.Email, emp.Phone)
		fmt.Printf("  - 所属部门: 部门编号=%s, 部门名称=%s, 部门编码=%s\n",
			emp.Department.DeptNo, emp.Department.Name, emp.Department.Code)
		fmt.Printf("  - 部门信息: 经理=%s, 位置=%s\n", emp.Department.Manager, emp.Department.Location)
		fmt.Println("  " + strings.Repeat("-", 50))
	}
	// 将对象转换为JSON字符串
	jsonData, err := json.MarshalIndent(empList, "", "  ")
	if err != nil {
		fmt.Printf("JSON转换失败: %v\n", err)
	} else {
		fmt.Printf("员工信息:\n%s\n", string(jsonData))
	}
}

// 演示查询技术部的员工
func demonstrateGetTechDepartmentEmployees() {
	fmt.Println("\n4. 查询技术部的员工:")
	var techDept models.DepartmentModel
	//SELECT * FROM `employee_models` WHERE `employee_models`.`dept_no_x` = 'DEPT001' AND `employee_models`.`deleted_at` IS NULL
	//SELECT * FROM `department_models` WHERE dept_no = 'DEPT001' AND `department_models`.`deleted_at` IS NULL ORDER BY `department_models`.`id` LIMIT 1
	result := global.DB.Preload("EmployeeList").Where("dept_no = ?", "DEPT001").First(&techDept)
	if result.Error != nil {
		fmt.Printf("查询技术部失败: %v\n", result.Error)
	} else {
		fmt.Printf("技术部: %s, 经理: %s\n", techDept.Name, techDept.Manager)
		fmt.Printf("员工列表 (%d人):\n", len(techDept.EmployeeList))
		for _, emp := range techDept.EmployeeList {
			fmt.Printf("  - 姓名: %s, 工号: %s, 职位: %s, 薪资: %.2f\n",
				emp.Name, emp.EmployeeNo, emp.Position, emp.Salary)
		}
	}
}

// 演示查询员工及其部门信息
func demonstrateGetEmployeesWithDepartments() {
	fmt.Println("\n5. 查询员工及其部门信息:")
	var empList []models.EmployeeModel
	result := global.DB.Preload("Department").Find(&empList)
	if result.Error != nil {
		fmt.Printf("查询员工失败: %v\n", result.Error)
		return
	}

	for _, emp := range empList {
		fmt.Printf("员工: %s (%s) - 职位: %s, 薪资: %.2f\n",
			emp.Name, emp.EmployeeNo, emp.Position, emp.Salary)
		fmt.Printf("  所属部门: %s (%s), 经理: %s\n",
			emp.Department.Name, emp.Department.Code, emp.Department.Manager)
	}
}

// 演示更新员工薪资
func demonstrateUpdateEmployeeSalary() {
	fmt.Println("\n6. 更新员工薪资:")
	var empToUpdate models.EmployeeModel
	result := global.DB.Where("employee_no = ?", "E001").First(&empToUpdate)
	if result.Error != nil {
		fmt.Printf("查询员工失败: %v\n", result.Error)
	} else {
		oldSalary := empToUpdate.Salary
		empToUpdate.Salary = 18000.00
		result = global.DB.Save(&empToUpdate)
		if result.Error != nil {
			fmt.Printf("更新员工薪资失败: %v\n", result.Error)
		} else {
			fmt.Printf("更新员工薪资成功! %s: %.2f -> %.2f\n",
				empToUpdate.Name, oldSalary, empToUpdate.Salary)
		}
	}
}

// 演示添加新员工到现有部门
func demonstrateAddNewEmployeeToDepartment() {
	fmt.Println("\n7. 添加新员工到技术部:")
	newEmp := models.EmployeeModel{
		Name:       "孙八",
		EmployeeNo: "E006",
		Position:   "架构师",
		Salary:     25000.00,
		Email:      "sunba@company.com",
		Phone:      "13800138006",
		DeptNoX:    "DEPT001", // 技术部
	}
	result := global.DB.Create(&newEmp)
	if result.Error != nil {
		fmt.Printf("创建新员工失败: %v\n", result.Error)
	} else {
		fmt.Printf("创建新员工成功! 姓名: %s, 工号: %s, 职位: %s, 薪资: %.2f\n",
			newEmp.Name, newEmp.EmployeeNo, newEmp.Position, newEmp.Salary)
	}
}

// 演示再次查询技术部员工
func demonstrateGetTechDepartmentEmployeesAgain() {
	fmt.Println("\n8. 再次查询技术部员工:")
	var updatedTechDept models.DepartmentModel
	result := global.DB.Preload("EmployeeList").Where("dept_no = ?", "DEPT001").First(&updatedTechDept)
	if result.Error != nil {
		fmt.Printf("查询技术部失败: %v\n", result.Error)
	} else {
		fmt.Printf("技术部: %s, 经理: %s\n", updatedTechDept.Name, updatedTechDept.Manager)
		fmt.Printf("员工列表 (%d人):\n", len(updatedTechDept.EmployeeList))
		for _, emp := range updatedTechDept.EmployeeList {
			fmt.Printf("  - 姓名: %s, 工号: %s, 职位: %s, 薪资: %.2f\n",
				emp.Name, emp.EmployeeNo, emp.Position, emp.Salary)
		}
	}
}

// 演示统计各部门员工数量
func demonstrateCountEmployeesByDepartment() {
	fmt.Println("\n9. 统计各部门员工数量:")
	var deptStats []models.DepartmentModel
	result := global.DB.Preload("EmployeeList").Find(&deptStats)
	if result.Error != nil {
		fmt.Printf("查询部门统计失败: %v\n", result.Error)
		return
	}

	for _, dept := range deptStats {
		fmt.Printf("%s: %d人\n", dept.Name, len(dept.EmployeeList))
	}
}

// 演示删除员工
func demonstrateDeleteEmployee() {
	fmt.Println("\n10. 删除员工:")
	var empToDelete models.EmployeeModel
	//SELECT * FROM `employee_models` WHERE employee_no = 'E006' AND `employee_models`.`deleted_at` IS NULL ORDER BY `employee_models`.`id` LIMIT 1
	//UPDATE `employee_models` SET `deleted_at`='2025-10-22 15:38:58.253' WHERE `employee_models`.`id` = 6package_manager AND `employee_models`.`deleted_at` IS NULL
	result := global.DB.Where("employee_no = ?", "E006").First(&empToDelete)
	if result.Error != nil {
		fmt.Printf("查询要删除的员工失败: %v\n", result.Error)
	} else {
		result = global.DB.Delete(&empToDelete)
		if result.Error != nil {
			fmt.Printf("删除员工失败: %v\n", result.Error)
		} else {
			fmt.Printf("删除员工成功! 删除了 %d 条记录\n", result.RowsAffected)
		}
	}
}

// 演示最终统计
func demonstrateFinalStatistics() {
	fmt.Println("\n11. 最终统计:")
	var finalStats []models.DepartmentModel
	result := global.DB.Preload("EmployeeList").Find(&finalStats)
	if result.Error != nil {
		fmt.Printf("查询最终统计失败: %v\n", result.Error)
		return
	}

	totalEmployees := 0
	for _, dept := range finalStats {
		empCount := len(dept.EmployeeList)
		totalEmployees += empCount
		fmt.Printf("%s: %d人\n", dept.Name, empCount)
	}
	fmt.Printf("总员工数: %d人\n", totalEmployees)
}
