package lianlianpay

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func PostRequest(url string, sign string, timestamp string, body []byte) (string, error) {
	client := resty.New()
	client.SetDebug(true)
	resp, err := client.R().
		SetHeader("signature", sign).
		SetHeader("timezone", "Asia/Hong_Kong").
		SetHeader("timestamp", timestamp).
		SetHeader("sign-type", "RSA").
		SetHeader("Content-Type", "application/json").
		SetHeader("version", "1.0.0").
		SetBody(body).
		Post(url)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	return resp.String(), nil
}

// PostRequest 发送 HTTP GET 请求
func GetRequest(url string, sign string, timestamp string, body []byte) ([]byte, error) {
	// 发起 GET 请求
	client := resty.New()
	client.SetDebug(true)
	fmt.Println("Request URL:", body)
	resp, err := client.R().
		SetHeader("signature", sign).
		SetHeader("timezone", "Asia/Hong_Kong").
		SetHeader("timestamp", timestamp).
		SetHeader("sign-type", "RSA").
		SetHeader("Content-Type", "application/json").
		SetHeader("version", "1.0.0").
		SetBody(body).
		Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	fmt.Println("Response:", string(resp.Body()))
	// 检查响应
	return resp.Body(), nil
}
