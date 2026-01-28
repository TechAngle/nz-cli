package utils

func CalculateSum[K float32, V int | float32](numbers []V) float32 {
	// sum
	var sum float32
	sum = 0

	for _, i := range numbers {
		sum += float32(i)
	}

	// log.Println(sum)

	return float32(sum)
}

func CalculateAverage[A float32, V float32 | int](nums []V) float32 {
	if len(nums) == 0 {
		return 0
	}

	sum := CalculateSum[A, V](nums)
	return sum / float32(len(nums))
}
