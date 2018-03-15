package lib

func RandFromString(text string) int64 {
	result := int64(0)
	for _, char := range text {
		result += int64(char)
	}
	return result
}
