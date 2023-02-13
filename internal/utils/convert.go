package utils

import "strconv"

func ConvertStringToInt(s string) int64 {
	i, _ := strconv.Atoi(s)
	return int64(i)
}
