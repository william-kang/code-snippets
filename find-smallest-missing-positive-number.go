type Solution struct{}

// findNumber finds the smallest missing positive number in the array.
func (s *Solution) findNumber(nums []int) int {
	for i := 1; i <= len(nums); i++ {
		for nums[i - 1] > 0 && nums[i - 1] <= len(nums) && nums[i - 1] != i && nums[i - 1] != nums[nums[i - 1] - 1] {
			nums[i - 1], nums[nums[i - 1] - 1] = nums[nums[i - 1] - 1], nums[i - 1]
		}
	}
	for i := 1; i <= len(nums); i++ {
		if nums[i - 1] != i { return i }
	}
	return len(nums) + 1
}
