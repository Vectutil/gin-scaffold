package utils

import (
	"strconv"
	"strings"
)

func String2Uint64(str string) uint64 {
	parseUint, _ := strconv.ParseUint(str, 10, 64)
	return parseUint
}
func String2int64(str string) int64 {
	parseUint, _ := strconv.ParseInt(str, 10, 64)
	return parseUint
}

func SplitTrim(str, splitMark string) []string {
	split := strings.Split(str, splitMark)
	strs := make([]string, 0)
	for _, v := range split {
		tv := strings.TrimSpace(v)
		if tv != "" {
			strs = append(strs, tv)
		}
	}
	return strs
}

// ExtractPercentageNumber 从百分比字符串中提取数字部分
func ExtractPercentageNumber(s string) float64 {
	// 检查字符串是否以百分号结尾
	if !strings.HasSuffix(s, "%") {
		return 0
	}

	// 去除末尾的百分号
	numberStr := s[:len(s)-1]

	// 简单验证数字格式（可以根据需要增强）
	if len(numberStr) == 0 {
		return 0
	}

	value, _ := strconv.ParseFloat(numberStr, 64)

	return value
}
