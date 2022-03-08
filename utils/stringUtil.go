package utils

// 处理字符串的函数

import (
	"bytes"
	"strings"
)

// 替换字符串 参数分别为 原始字符串,目标字符串,要替换的目标
func InsteadString(str, source, target string) string {
	index := strings.Index(str, source)
	sourceLen := len(source)
	start := str[:index]
	end := str[index+sourceLen:]
	var result bytes.Buffer
	result.WriteString(start)
	result.WriteString(target)
	result.WriteString(end)
	return result.String()
}

// 返回unicode字符的位置
func UnicodeIndex(str, substr string) int {
	// 子串在字符串的字节位置
	result := strings.Index(str, substr)
	if result >= 0 {
		// 获得子串之前的字符串并转换成[]byte
		prefix := []byte(str)[0:result]
		// 将子串之前的字符串转换成[]rune
		rs := []rune(string(prefix))
		// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
		result = len(rs)
	} else {
		return -1
	}
	return result
}
