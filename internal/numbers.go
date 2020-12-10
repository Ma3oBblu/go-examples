package internal

// GetMaxInt генерирует максимальное значение для int
func GetMaxInt() int {
	return int(^uint(0) >> 1)
}
