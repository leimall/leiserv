package lianlianpay

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// ConvertStructToSignatureString - 将结构体转换为签名字符串
func ConvertStructToSignatureString(data interface{}) string {
	return convertMapToSignatureString(structToMap(data))
}

// structToMap - 递归将结构体转换为 map[string]interface{}
func structToMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := v.Type().Field(i)
		fieldName := fieldType.Tag.Get("json")
		if fieldName == "" {
			fieldName = fieldType.Name
		}

		switch field.Kind() {
		case reflect.Struct:
			result[fieldName] = structToMap(field.Interface())
		case reflect.Slice:
			var sliceItems []interface{}
			for j := 0; j < field.Len(); j++ {
				if field.Index(j).Kind() == reflect.Struct {
					sliceItems = append(sliceItems, structToMap(field.Index(j).Interface()))
				} else {
					sliceItems = append(sliceItems, field.Index(j).Interface())
				}
			}
			result[fieldName] = sliceItems
		default:
			result[fieldName] = field.Interface()
		}
	}
	return result
}

// convertMapToSignatureString - 递归地将 map 转为有序字符串
func convertMapToSignatureString(data map[string]interface{}) string {
	var parts []string
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		switch v := data[k].(type) {
		case map[string]interface{}:
			// 如果是嵌套 map，递归排序并生成字符串
			parts = append(parts, convertMapToSignatureString(v))
		case []interface{}:
			// 如果是数组类型，处理数组的每个元素
			arrParts := processArray(v)
			parts = append(parts, strings.Join(arrParts, "&"))
		default:
			// 基本类型，直接转为 key=value 格式
			parts = append(parts, fmt.Sprintf("%s=%v", k, v))
		}
	}
	return strings.Join(parts, "&")
}

// processArray - 处理数组中的每个对象元素
func processArray(arr []interface{}) []string {
	var sortedArr []string
	for _, item := range arr {
		if m, ok := item.(map[string]interface{}); ok {
			// 如果是 map 类型，递归处理
			sortedArr = append(sortedArr, convertMapToSignatureString(m))
		}
	}
	// 按照字母顺序对每个对象排序
	sort.Strings(sortedArr)
	return sortedArr
}
