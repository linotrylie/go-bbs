package sliceutil

import "strings"

func ContainsAny(data string, sonStr []string) bool {
	for _, str := range sonStr {
		if strings.Contains(data, str) {
			return true
		}
	}
	return false
}

// SliceEleIsDuplicated 切片元素是否重复
func SliceEleIsDuplicated(list []string) bool {
	tmpMap := make(map[string]int)
	for _, value := range list {
		tmpMap[value] = 1
	}

	var keys []interface{}
	for k := range tmpMap {
		keys = append(keys, k)
	}
	if len(keys) != len(list) {
		return true
	}
	return false
}
