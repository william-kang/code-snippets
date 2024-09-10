type Solution struct{}

// findMissingNumber finds the first number missing from its index.
func (s *Solution) findMissingNumber(nums []int) int {
    nums = append(nums, -1)
    for i := 0; i < len(nums); i++ {
        for nums[i] != i && nums[i] != -1 {
            nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
        }
    }
    for i := 0; i < len(nums); i++ {
        if nums[i] == -1 {
            return i
        }
    }
    return -1
}
