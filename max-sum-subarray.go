import (
    "fmt"
)

// Solution struct
type Solution struct{}

// findMaxSumSubArray method to find the maximum sum of a subarray of size k
func (s *Solution) findMaxSumSubArray(k int, arr []int) int {
    maxSum := 0
    sum := 0
    windowStart := 0
    for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
        sum += arr[windowEnd]
        if (windowEnd >= k - 1) {
            fmt.Println(sum)
            if sum > maxSum { maxSum = sum }
            sum -= arr[windowStart]
            windowStart++
        }
    }
    return maxSum
}
