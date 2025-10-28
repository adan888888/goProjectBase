package models

// API响应结构体
type Response struct {
	Code int `json:"code"`
	Data struct {
		BiasValue   int    `json:"biasValue"`
		RandomValue int    `json:"randomValue"`
		Result      string `json:"result"`
		Timestamp   int64  `json:"timestamp"`
		Value       int    `json:"value"`
	} `json:"data"`
	Msg string `json:"msg"`
}
