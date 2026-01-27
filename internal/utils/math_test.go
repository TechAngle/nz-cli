package utils

import "testing"

var numbersSlice = [...]int{1, 2, 3} // Sum is 6, average is 2

func TestCalculateSum(t *testing.T) {
	sumResult := 6
	sum := CalculateSum[float32, int](numbersSlice[:])

	if sum != float32(sumResult) {
		t.Fatalf("Sum is not %d", sumResult)
	}
}

func TestCalculateAverage(t *testing.T) {
	averageResult := 2
	average := CalculateAverage[float32, int](numbersSlice[:])

	if average != float32(averageResult) {
		t.Fatalf("Average should be %d", averageResult)
	}
}
