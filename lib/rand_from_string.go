package lib

import "time"

// RandFromString Placeholder comment
func RandFromString(text string, params ...interface{}) int64 {
	result := int64(0)
	for _, char := range text {
		result = (result + int64(char)) % 8096
	}
	if len(params) >= 1 {
		result += int64(params[0].(int))
	} else {
		result += time.Now().Unix() % 3600
	}
	return result
}
