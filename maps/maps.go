package maps

import (
	"strings"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	arr := strings.Fields(s)
	for _, val := range arr {
		_, ok := result[val]
		if ok {
			result[val] += 1
		} else {
			result[val] = 1
		}
	}
	return result
}
