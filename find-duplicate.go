type Solution struct{}

func (s *Solution) findNumber(nums []int) int {
	for i := 1; i <= len(nums); i++ {
		for nums[i - 1] != i {
			if nums[i - 1] == nums[nums[i - 1] - 1] { return nums[i - 1] }
			nums[nums[i - 1] - 1], nums[i - 1] = nums[i - 1], nums[nums[i - 1] - 1]
		}
	}
	return -1
}
