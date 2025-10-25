package main

import (
	"4gorm/global"
	"4gorm/models"
	"encoding/json"
	"fmt"
)

func main() {
	// 初始化数据库连接
	global.InitDB()

	// 自动迁移数据库表
	global.MigrateDB()

	fmt.Println("\n=== GORM 学生课程多对多关系演示 ===")

	// 演示两种不同的数据插入方式
	//fmt.Println("\n=== 方式1: 使用自定义中间表（手动关联）===")
	//createStudentsWithCourses()

	//fmt.Println("\n=== 方式2: 使用 GORM 自动关联（标准中间表）===")
	//demonstrateGORMAutoAssociation()

	//同时创建与关联
	GormCreateAutoAssociation()

	//// 3. 查询所有学生及其课程
	//getAllStudentsWithCourses()
	//
	//// 4. 查询所有课程及其学生
	//getAllCoursesWithStudents()

	fmt.Println("\n=== 学生课程多对多关系演示完成 ===")
}

func GormCreateAutoAssociation() {
	//创建学习 连代 创建学生的选修的课程
	global.DB.SetupJoinTable(&models.StudentModel{}, "Courses", &models.StudentCourse{}) ///‼️️只有加了这个方法才会走第三表的创建钩子
	studentModels := models.StudentModel{
		Name: "李四",
		Courses: []models.CourseModel{
			{
				Name: "gin框架",
			},
			{
				Name: "zero",
			},
		}, // 选修两门课程
	}
	err := global.DB.Create(&studentModels).Error
	fmt.Println(err)
}

// 创建学生并同时选课
func createStudentsWithCourses() {
	fmt.Println("\n1. 创建学生并同时选课:")
	global.DB.Exec("DELETE FROM student_courses")
	global.DB.Exec("DELETE FROM student_models")
	global.DB.Exec("DELETE FROM course_models")
	// 先创建所有课程
	courses := []models.CourseModel{
		{
			Name: "数据结构与算法",
		},
		{
			Name: "数据库原理",
		},
		{
			Name: "软件工程",
		},
	}

	// 创建课程
	for _, course := range courses {
		result := global.DB.Create(&course)
		if result.Error != nil {
			fmt.Printf("创建课程失败: %v\n", result.Error)
			continue
		}
		fmt.Printf("创建课程成功! ID: %d, 课程名称: %s\n",
			course.ID, course.Name)
	}

	// 重新查询课程以获取正确的ID
	var allCourses []models.CourseModel
	global.DB.Find(&allCourses)

	// 创建课程映射，用于查找课程ID
	courseMap := make(map[uint]models.CourseModel)
	for _, course := range allCourses {
		courseMap[course.ID] = course
	}

	// 定义学生和他们的选课信息
	studentData := []struct {
		Name      string
		CourseIDs []uint // 要选修的课程ID列表
	}{
		{
			Name:      "张三",
			CourseIDs: []uint{1, 2}, // 选修前两门课程
		},
		{
			Name:      "李四",
			CourseIDs: []uint{2, 3}, // 选修后两门课程
		},
		{
			Name:      "王五",
			CourseIDs: []uint{1, 2, 3}, // 选修所有课程
		},
	}

	// 创建学生和选课记录
	for _, data := range studentData {
		// 创建学生
		student := models.StudentModel{
			Name: data.Name,
		}

		result := global.DB.Create(&student)
		if result.Error != nil {
			fmt.Printf("创建学生失败: %v\n", result.Error)
			continue
		}

		// 手动创建选课记录
		for _, courseID := range data.CourseIDs {
			if course, exists := courseMap[courseID]; exists {
				studentCourse := models.StudentCourse{
					StudentID: student.ID,
					CourseID:  course.ID,
				}
				global.DB.Create(&studentCourse)
			}
		}

		fmt.Printf("创建学生成功! ID: %d, 姓名: %s, 选修课程: %d门\n",
			student.ID, student.Name, len(data.CourseIDs))

		// 显示选修的课程详情
		for _, courseID := range data.CourseIDs {
			if course, exists := courseMap[courseID]; exists {
				fmt.Printf("  - %s\n", course.Name)
			}
		}
	}
}

// 演示 GORM 自动关联（使用标准中间表）
func demonstrateGORMAutoAssociation() {
	fmt.Println("\n2. 使用 GORM 自动关联创建学生和课程:")

	// 先清空现有数据，避免冲突
	global.DB.Exec("DELETE FROM student_courses")
	global.DB.Exec("DELETE FROM student_models")
	global.DB.Exec("DELETE FROM course_models")

	// 创建课程
	courses := []models.CourseModel{
		{
			Name: "计算机网络",
		},
		{
			Name: "操作系统",
		},
	}

	// 创建课程
	for _, course := range courses {
		result := global.DB.Create(&course)
		if result.Error != nil {
			fmt.Printf("创建课程失败: %v\n", result.Error)
			continue
		}
		fmt.Printf("创建课程成功! ID: %d, 课程名称: %s\n",
			course.ID, course.Name)
	}

	// 创建学生（不包含 Courses 字段，避免 GORM 自动关联）
	students := []models.StudentModel{
		{
			Name: "小明",
		},
		{
			Name: "小红",
		},
	}

	// 创建学生
	for _, student := range students {
		result := global.DB.Create(&student)
		if result.Error != nil {
			fmt.Printf("创建学生失败: %v\n", result.Error)
			continue
		}
		fmt.Printf("创建学生成功! ID: %d, 姓名: %s\n",
			student.ID, student.Name)
	}

	// 使用 GORM 的 Association 功能进行关联
	// 获取第一个学生
	var firstStudent models.StudentModel
	global.DB.First(&firstStudent)

	// 获取所有课程
	var allCourses []models.CourseModel
	global.DB.Find(&allCourses)

	// 使用 GORM 的 Association 功能
	// 这会自动创建标准中间表记录（不包含自定义字段）
	global.DB.Model(&firstStudent /*一个学生*/).Association("Courses" /*学生模型中的Courses字段*/).Append(allCourses /*多个课程*/)

	fmt.Printf("使用 GORM 自动关联成功! 学生 '%s' 选修了 %d 门课程\n",
		firstStudent.Name, len(allCourses))

	// 注意：这种方式创建的中间表记录不包含 student_name 和 course_name 字段
	// 因为 GORM 的自动关联只处理标准的多对多关系
}

// 创建学生
func createStudents() {
	fmt.Println("\n1. 创建学生:")
	students := []models.StudentModel{
		{
			Name: "张三",
		},
		{
			Name: "李四",
		},
		{
			Name: "王五",
		},
	}

	for _, student := range students {
		result := global.DB.Create(&student)
		if result.Error != nil {
			fmt.Printf("创建学生失败: %v\n", result.Error)
			continue
		}
		fmt.Printf("创建学生成功! ID: %d, 姓名: %s\n",
			student.ID, student.Name)
	}
}

// 创建课程
func createCourses() {
	fmt.Println("\n2. 创建课程:")
	courses := []models.CourseModel{
		{
			Name: "数据结构与算法",
		},
		{
			Name: "数据库原理",
		},
		{
			Name: "软件工程",
		},
	}

	for _, course := range courses {
		result := global.DB.Create(&course)
		if result.Error != nil {
			fmt.Printf("创建课程失败: %v\n", result.Error)
			continue
		}
		fmt.Printf("创建课程成功! ID: %d, 课程名称: %s\n",
			course.ID, course.Name)
	}
}

// 查询所有学生及其课程
func getAllStudentsWithCourses() {
	fmt.Println("\n3. 查询所有学生及其课程:")
	var students []models.StudentModel
	result := global.DB.Preload("Courses").Find(&students)
	if result.Error != nil {
		fmt.Printf("查询学生失败: %v\n", result.Error)
		return
	}

	fmt.Printf("查询到 %d 个学生:\n", len(students))
	for i, student := range students {
		fmt.Printf("\n学生 %d:\n", i+1)
		fmt.Printf("  - 基本信息: ID=%d, 姓名=%s\n",
			student.ID, student.Name)
		fmt.Printf("  - 选修课程 (%d门):\n", len(student.Courses))
		for j, course := range student.Courses {
			fmt.Printf("    %d. %s\n",
				j+1, course.Name)
		}
		fmt.Println("  " + "--------------------------------------------------")
	}
	// 将对象转换为JSON字符串
	jsonData, _ := json.MarshalIndent(students, "", "  ")
	fmt.Printf("学生信息:\n%s\n", string(jsonData))
}

// 查询所有课程及其学生
func getAllCoursesWithStudents() {
	fmt.Println("\n4. 查询所有课程及其学生:")
	var courses []models.CourseModel
	result := global.DB.Preload("Students").Find(&courses)
	if result.Error != nil {
		fmt.Printf("查询课程失败: %v\n", result.Error)
		return
	}

	fmt.Printf("查询到 %d 门课程:\n", len(courses))
	for i, course := range courses {
		fmt.Printf("\n课程 %d:\n", i+1)
		fmt.Printf("  - 基本信息: ID=%d, 课程名称=%s\n",
			course.ID, course.Name)
		fmt.Printf("  - 选修学生 (%d人):\n", len(course.Students))
		for j, student := range course.Students {
			fmt.Printf("    %d. %s\n",
				j+1, student.Name)
		}
		fmt.Println("  " + "--------------------------------------------------")
	}
}

// 学生选课
func enrollStudentInCourses() {
	fmt.Println("\n5. 学生选课:")

	// 获取第一个学生
	var student models.StudentModel
	global.DB.First(&student)

	// 获取前3门课程
	var courses []models.CourseModel
	global.DB.Limit(3).Find(&courses)

	// 学生选课
	global.DB.Model(&student).Association("Courses").Append(courses)

	fmt.Printf("学生 '%s' 选修了 %d 门课程:\n", student.Name, len(courses))
	for _, course := range courses {
		fmt.Printf("  - %s\n",
			course.Name)
	}
}

// 查询特定学生的课程
func getCoursesByStudent() {
	fmt.Println("\n6. 查询特定学生的课程:")

	// 获取第一个学生
	var student models.StudentModel
	global.DB.First(&student)

	// 查询该学生的所有课程
	var courses []models.CourseModel
	global.DB.Model(&student).Association("Courses").Find(&courses)

	fmt.Printf("学生 '%s' 的课程 (%d门):\n", student.Name, len(courses))
	for i, course := range courses {
		fmt.Printf("  %d. %s\n",
			i+1, course.Name)
	}
}

// 查询特定课程的学生
func getStudentsByCourse() {
	fmt.Println("\n7. 查询特定课程的学生:")

	// 获取第一门课程
	var course models.CourseModel
	global.DB.First(&course)

	// 查询该课程的所有学生
	var students []models.StudentModel
	global.DB.Model(&course).Association("Students").Find(&students)

	fmt.Printf("课程 '%s' 的学生 (%d人):\n", course.Name, len(students))
	for i, student := range students {
		fmt.Printf("  %d. %s\n",
			i+1, student.Name)
	}
}

// 学生退课
func dropCourseFromStudent() {
	fmt.Println("\n8. 学生退课:")

	// 获取第一个学生
	var student models.StudentModel
	global.DB.First(&student)

	// 获取第一门课程
	var course models.CourseModel
	global.DB.First(&course)

	// 学生退课
	global.DB.Model(&student).Association("Courses").Delete(course)

	fmt.Printf("学生 '%s' 退选了课程 '%s'\n", student.Name, course.Name)
}

// 最终统计
func finalStatistics() {
	fmt.Println("\n9. 最终统计:")

	// 统计学生数量
	var studentCount int64
	global.DB.Model(&models.StudentModel{}).Count(&studentCount)

	// 统计课程数量
	var courseCount int64
	global.DB.Model(&models.CourseModel{}).Count(&courseCount)

	// 统计选课关系数量
	var relationCount int64
	global.DB.Table("student_courses").Count(&relationCount)

	fmt.Printf("学生总数: %d人\n", studentCount)
	fmt.Printf("课程总数: %d门\n", courseCount)
	fmt.Printf("选课关系: %d个\n", relationCount)
}
