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
