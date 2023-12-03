package utils

func Clamp(index int, _min int, _max int) int {
	return min(_max, max(index, _min))
}
