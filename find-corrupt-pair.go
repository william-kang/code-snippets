type Solution struct{}

func (s *Solution) findNumbers(nums []int) []int {
	for i := 1; i <= len(nums); i++ {
		for nums[i - 1] != i && nums[i - 1] != nums[nums[i - 1] - 1] {
			nums[nums[i - 1] - 1], nums[i - 1] = nums[i - 1], nums[nums[i - 1] - 1]
		}
	}
	for i := 1; i <= len(nums); i++ {
		if nums[i - 1] != i {
			return []int{nums[i - 1], i}
		}
	}
	return []int{-1, -1}
}
