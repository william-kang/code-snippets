// Solution struct
type Solution struct{}

// findNumbers finds duplicate numbers in an array.
func (s *Solution) findNumbers(nums []int) []int {
	duplicateNumbers := make([]int, 0)
	for i := 1; i <= len(nums); i++ {
		for nums[i - 1] != i && nums[i - 1] != nums[nums[i - 1] - 1] {
			nums[nums[i - 1] - 1], nums[i - 1] = nums[i - 1], nums[nums[i - 1] - 1]
		}
	}
	for i:= 0; i < len(nums); i++ {
		if (nums[i] != i + 1) {
			duplicateNumbers = append(duplicateNumbers, nums[i])
		}
	}
	return duplicateNumbers
}
