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

	// 0. 创建学生并同时选课
	createStudentsWithCourses()

	//// 1.创建学生
	//createStudents()
	//// 2.创建课程
	//createCourses()
	// 3. 查询所有学生及其课程
	getAllStudentsWithCourses()

	//// 4. 查询所有课程及其学生
	//getAllCoursesWithStudents()
	//
	//// 5. 学生选课
	//enrollStudentInCourses()
	//
	//// 6. 查询特定学生的课程
	//getCoursesByStudent()
	//
	//// 7. 查询特定课程的学生
	//getStudentsByCourse()
	//
	//// 8. 学生退课
	//dropCourseFromStudent()
	//
	//// 9. 最终统计
	//finalStatistics()
	//
	//fmt.Println("\n=== 学生课程多对多关系演示完成 ===")
}

// 创建学生并同时选课
func createStudentsWithCourses() {
	fmt.Println("\n1. 创建学生并同时选课:")

	// 先创建课程
	courses := []models.CourseModel{
		{
			CourseNo: "CS301",
			Name:     "数据结构与算法",
			Teacher:  "张教授",
		},
		{
			CourseNo: "CS302",
			Name:     "数据库原理",
			Teacher:  "李教授",
		},
		{
			CourseNo: "CS303",
			Name:     "软件工程",
			Teacher:  "王教授",
		},
	}

	// 创建课程
	for _, course := range courses {
		result := global.DB.Create(&course)
		if result.Error != nil {
			fmt.Printf("创建课程失败: %v\n", result.Error)
			continue
		}
		fmt.Printf("创建课程成功! ID: %d, 课程编号: %s, 课程名称: %s, 学分: %d, 教师: %s\n",
			course.ID, course.CourseNo, course.Name, course.Teacher)
	}

	// 创建学生并同时选课
	students := []struct {
		StudentNo string
		Name      string
		Major     string
		CourseIDs []uint // 要选修的课程ID列表
	}{
		{
			StudentNo: "2025001",
			Name:      "张三",
			Major:     "计算机科学与技术",
			CourseIDs: []uint{1, 2}, // 选修前两门课程
		},
		{
			StudentNo: "2025002",
			Name:      "李四",
			Major:     "软件工程",
			CourseIDs: []uint{2, 3}, // 选修后两门课程
		},
		{
			StudentNo: "2025003",
			Name:      "王五",
			Major:     "数据科学与大数据技术",
			CourseIDs: []uint{1, 2, 3}, // 选修所有课程
		},
	}

	for _, studentData := range students {
		// 创建学生
		student := models.StudentModel{
			StudentNo: studentData.StudentNo,
			Name:      studentData.Name,
			Major:     studentData.Major,
		}

		result := global.DB.Create(&student)
		if result.Error != nil {
			fmt.Printf("创建学生失败: %v\n", result.Error)
			continue
		}

		// 查询要选修的课程
		var selectedCourses []models.CourseModel
		global.DB.Where("id IN ?", studentData.CourseIDs).Find(&selectedCourses)

		// 关联课程到学生
		if len(selectedCourses) > 0 {
			//Association ("Courses") 必须与模型中的字段名完全一致 -- 指定要操作的关联字段
			//Append() 方法的作用:1.将课程添加到学生的课程列表中  2.自动处理关联表的插入 -- 执行实际的关联操作
			global.DB.Model(&student).Association("Courses").Append(selectedCourses)
		}

		fmt.Printf("创建学生成功! ID: %d, 学号: %s, 姓名: %s, 专业: %s, 选修课程: %d门\n",
			student.ID, student.StudentNo, student.Name, student.Major, len(selectedCourses))

		// 显示选修的课程详情
		for _, course := range selectedCourses {
			fmt.Printf("  - %s (%s) - 学分:%d, 教师:%s\n",
				course.Name, course.CourseNo, course.Teacher)
		}
	}
}

// 创建学生
func createStudents() {
	fmt.Println("\n1. 创建学生:")
	students := []models.StudentModel{
		{
			StudentNo: "2024001",
			Name:      "张三",
			Major:     "计算机科学与技术",
		},
		{
			StudentNo: "2024002",
			Name:      "李四",
			Major:     "软件工程",
		},
		{
			StudentNo: "2023001",
			Name:      "王五",
			Major:     "数据科学与大数据技术",
		},
	}

	for _, student := range students {
		result := global.DB.Create(&student)
		if result.Error != nil {
			fmt.Printf("创建学生失败: %v\n", result.Error)
			continue
		}
		fmt.Printf("创建学生成功! ID: %d, 学号: %s, 姓名: %s, 专业: %s\n",
			student.ID, student.StudentNo, student.Name, student.Major)
	}
}

// 创建课程
func createCourses() {
	fmt.Println("\n2. 创建课程:")
	courses := []models.CourseModel{
		{
			CourseNo: "CS101",
			Name:     "数据结构与算法",
			Teacher:  "张教授",
		},
		{
			CourseNo: "CS102",
			Name:     "数据库原理",
			Teacher:  "李教授",
		},
		{
			CourseNo: "CS201",
			Name:     "软件工程",
			Teacher:  "王教授",
		},
	}

	for _, course := range courses {
		result := global.DB.Create(&course)
		if result.Error != nil {
			fmt.Printf("创建课程失败: %v\n", result.Error)
			continue
		}
		fmt.Printf("创建课程成功! ID: %d, 课程编号: %s, 课程名称: %s, 学分: %d, 教师: %s\n",
			course.ID, course.CourseNo, course.Name, course.Teacher)
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
		fmt.Printf("  - 基本信息: ID=%d, 学号=%s, 姓名=%s, 专业=%s\n",
			student.ID, student.StudentNo, student.Name, student.Major)
		fmt.Printf("  - 选修课程 (%d门):\n", len(student.Courses))
		for j, course := range student.Courses {
			fmt.Printf("    %d. %s (%s) - 学分:%d, 教师:%s\n",
				j+1, course.Name, course.CourseNo, course.Teacher)
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
		fmt.Printf("  - 基本信息: ID=%d, 课程编号=%s, 课程名称=%s\n",
			course.ID, course.CourseNo, course.Name)
		fmt.Printf("  - 课程详情:  教师=%s\n", course.Teacher)
		fmt.Printf("  - 选修学生 (%d人):\n", len(course.Students))
		for j, student := range course.Students {
			fmt.Printf("    %d. %s (%s) - 专业:%s\n",
				j+1, student.Name, student.StudentNo, student.Major)
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
		fmt.Printf("  - %s (%s) - 学分:%d, 教师:%s\n",
			course.Name, course.CourseNo, course.Teacher)
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
		fmt.Printf("  %d. %s (%s) - 学分:%d, 教师:%s\n",
			i+1, course.Name, course.CourseNo, course.Teacher)
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
		fmt.Printf("  %d. %s (%s) - 专业:%s\n",
			i+1, student.Name, student.StudentNo, student.Major)
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
