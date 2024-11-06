package lianlianpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

func PostRequest(url string, data map[string]interface{}) (string, error) {
	// 获取当前时间戳

	timestamp := time.Now().Format("20060102150405") // 格式为：yyyyMMddHHmmss
	data["Timestamp"] = timestamp
	// Marshal 请求体为 JSON

	signatureString := ConvertStructToSignatureString(data)
	fmt.Println("Signature String:", signatureString)

	bodyBytes, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %w", err)
	}

	// 生成签名
	signature, err := Sign(signatureString)
	if err != nil {
		log.Println("Error signing data:", err)
		return "", fmt.Errorf("failed to generate signature: %w", err)
	}
	// 创建 HTTP POST 请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// 设置请求头
	req.Header.Set("signature", signature)
	req.Header.Set("timezone", "Asia/Hong_Kong")
	req.Header.Set("timestamp", timestamp)
	req.Header.Set("sign-type", "RSA")
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(responseBody), nil
}

type TokenType struct {
	MerchantID string `json:"merchant_id"`
	Timestamp  string `json:"timestamp"`
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
