package main

import (
	"4gorm/global"
	"4gorm/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func main() {
	global.InitDB()
	//global.MigrateDB()
	//
	//global.DB.Create(&models.Record{
	//	Name:  "张四",
	//	Money: 1000,
	//})

	//事务测试
	testTransaction()
}

func testTransaction() {
	var zhangSan = models.Record{ID: 1}
	var liSi = models.Record{ID: 2}
	////张三给李四转100元
	//global.DB.Model(&zhangSan).Update("money", gorm.Expr("money-100"))
	//panic("减钱失败")
	//global.DB.Model(&liSi).Update("money", gorm.Expr("money+100"))
	//假如第一个sql执行失败了。第二个就不会执行。导致账出现问题。  这个时候就引用了事务

	//一。自动事务 tx
	var err = global.DB.Transaction(func(tx *gorm.DB) error {
		//只要有一个没有执行成功，就会全部回滚
		err := tx.Model(&zhangSan).Update("money", gorm.Expr("money-100")).Error
		err = errors.New("出错了")
		if err != nil {
			return err
		}
		tx.Model(&liSi).Update("money", gorm.Expr("money+100"))
		return nil
	})
	fmt.Println(err)

	//二。手动控制
	// 开始事务
	tx := global.DB.Begin()
	if tx.Error != nil {
		fmt.Printf("开始事务失败: %v\n", tx.Error)
		return
	}

	// 执行第一个操作
	err = tx.Model(&zhangSan).Update("money", gorm.Expr("money-100")).Error
	if err != nil {
		tx.Rollback() // 回滚事务
		fmt.Printf("第一个操作失败: %v\n", err)
		return
	}

	// 模拟错误
	//err = errors.New("出错了")
	if err != nil {
		tx.Rollback() // 回滚事务
		fmt.Printf("业务逻辑错误: %v\n", err)
		return
	}

	// 执行第二个操作
	err = tx.Model(&liSi).Update("money", gorm.Expr("money+100")).Error
	if err != nil {
		tx.Rollback() // 回滚事务
		fmt.Printf("第二个操作失败: %v\n", err)
		return
	}

	// 提交事务
	err = tx.Commit().Error
	if err != nil {
		fmt.Printf("提交事务失败: %v\n", err)
	} else {
		fmt.Println("事务执行成功")
	}
}
