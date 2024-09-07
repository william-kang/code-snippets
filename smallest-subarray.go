import (
	"math"
)

type Solution struct{}

func (s *Solution) findMinSubArray(S int, arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	if sum < S { return 0 }

	minLength := math.MaxInt32
	windowSum := 0
	windowStart := 0
	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		windowSum += arr[windowEnd] // add the next element
		// slide the window
		for windowSum >= S {
			if minLength > windowEnd - windowStart + 1 {
				minLength = windowEnd - windowStart + 1
			}
			windowSum -= arr[windowStart]
			windowStart++
		}
	}
	return minLength
}
