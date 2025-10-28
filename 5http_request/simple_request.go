package main

import (
	"encoding/json"
	"fmt"
	"goProjectBase/5http_request/models"
	"io"
	"net/http"
)

func main() { simpleRequest() }

func simpleRequest() {
	// 调用接口
	resp, err := http.Get("http://localhost:3000/api/ycd/randomBankerPlayer")
	if err != nil {
		fmt.Printf("请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		return
	}

	// 解析JSON响应
	var response models.Response
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Printf("JSON解析失败: %v\n", err)
		fmt.Printf("原始响应: %s\n", string(body))
		return
	}

	// 美化JSON输出
	prettyJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Printf("JSON格式化失败: %v\n", err)
		return
	}

	fmt.Printf("JSON响应:\n%s\n", string(prettyJSON))

	// 单独显示结果
	fmt.Printf("\n结果: %s\n", response.Data.Result)
}
