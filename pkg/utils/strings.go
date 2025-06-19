package utils

import "strconv"

func String2Uint64(str string) uint64 {
	parseUint, _ := strconv.ParseUint(str, 10, 64)
	return parseUint
}
func String2int64(str string) int64 {
	parseUint, _ := strconv.ParseInt(str, 10, 64)
	return parseUint
}
