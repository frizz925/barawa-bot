package lib

import "time"

func RandFromString(params ...interface{}) int64 {
	text := params[0].(string)
	result := int64(0)
	for _, char := range text {
		result += int64(char)
	}
	if len(params) >= 2 {
		result += int64(params[1].(int))
	} else {
		result += time.Now().Unix() % 3600
	}
	return result
}
