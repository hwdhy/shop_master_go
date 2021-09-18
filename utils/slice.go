package utils

import "time"

// SliceRepeat 切片去重
func SliceRepeat(array []int64) []int64 {
	arrayMap := make(map[int64]bool)
	var result []int64
	for _, item := range array {
		if !arrayMap[item] {
			arrayMap[item] = true
			result = append(result, item)
		}
	}
	return result
}

// ContainsInt 元素是否存在切片中
func ContainsInt(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// DateEqual 时间比较
func DateEqual(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}