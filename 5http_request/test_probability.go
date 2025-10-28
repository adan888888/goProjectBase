package main

import (
	"encoding/json"
	"fmt"
	"goProjectBase/5http_request/models"
	"io"
	"net/http"
	"time"
)

func main() { testProbability() }
func testProbability() {
	url := "http://localhost:3000/api/ycd/randomBankerPlayer"

	// 统计变量
	totalCount := 100000
	bankerCount := 0

	fmt.Printf("开始测试，总共调用 %d 次接口...\n", totalCount)
	startTime := time.Now()

	for i := 0; i < totalCount; i++ {
		// 调用接口
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("第 %d 次请求失败: %v\n", i+1, err)
			continue
		}

		// 读取响应
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Printf("第 %d 次读取响应失败: %v\n", i+1, err)
			continue
		}

		// 解析JSON
		var response models.Response
		if err := json.Unmarshal(body, &response); err != nil {
			fmt.Printf("第 %d 次JSON解析失败: %v\n", i+1, err)
			continue
		}

		// 统计"庄"的出现次数
		if response.Data.Result == "庄" {
			bankerCount++
		}

		// 每10000次显示一次进度
		if (i+1)%10000 == 0 {
			fmt.Printf("已完成 %d 次调用，当前庄出现次数: %d\n", i+1, bankerCount)
		}
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)

	// 计算概率
	probability := float64(bankerCount) / float64(totalCount) * 100

	// 输出结果
	fmt.Printf("\n=== 测试结果 ===\n")
	fmt.Printf("总调用次数: %d\n", totalCount)
	fmt.Printf("庄出现次数: %d\n", bankerCount)
	fmt.Printf("庄出现概率: %.2f%%\n", probability)
	fmt.Printf("测试耗时: %v\n", duration)
}
