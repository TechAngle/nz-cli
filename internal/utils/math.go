package utils

// Calculates sum of numbers slice (int or float32)
func CalculateSum[K float32, V int | float32](numbers []V) float32 {
	var sum float32 = 0

	for _, i := range numbers {
		sum += float32(i)
	}

	return float32(sum)
}

// Calculates average of numbers slice (int or float32)
func CalculateAverage[A float32, V float32 | int](nums []V) float32 {
	if len(nums) == 0 {
		return 0
	}

	sum := CalculateSum[A, V](nums)
	return sum / float32(len(nums))
}
