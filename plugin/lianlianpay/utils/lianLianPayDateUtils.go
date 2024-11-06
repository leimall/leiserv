package utils

import (
	"time"
)

const defaultPattern = "20241024150405"

// 获取当前时间戳
func GetTimestamp() string {
	return Format(time.Now(), defaultPattern)
}

func Format(date time.Time, pattern string) string {
	return date.Format(pattern)
}
