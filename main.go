package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func translate(source string, direction string) (string, error) {
	// API 请求地址
	url := "http://api.interpreter.caiyunai.com/v1/translator"

	// 你的彩云小译开发者令牌
	token := "Token"

	// 准备 API 请求的 payload
	payload := map[string]interface{}{
		"source":     source,
		"trans_type": direction,
		"request_id": "demo",
		"detect":     true,
	}

	// 将 payload 转换为 JSON 格式
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %v", err)
	}

	// 创建 HTTP POST 请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %v", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-authorization", "token "+token)

	// 发送请求并获取响应
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("HTTP request failed: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	// 解析 JSON 响应
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", fmt.Errorf("failed to parse translation result: %v", err)
	}

	// 提取翻译结果
	target, ok := result["target"].(string)
	if !ok {
		return "", fmt.Errorf("failed to parse translation target")
	}

	return target, nil
}

func main() {
	// 待翻译的文本
	source := "World War Z: Aftermath"

	// 翻译方向
	direction := "en2zh"

	// 调用翻译函数
	target, err := translate(source, direction)
	if err != nil {
		fmt.Println("translation error:", err)
		return
	}

	// 打印翻译结果
	fmt.Println("translated text:", target)
}
