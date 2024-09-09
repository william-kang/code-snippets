type Solution struct{}

// sort rearranges the elements in the slice in ascending order.
func (s *Solution) sort(Nums []int) []int {
	i := 1
	for i <= len(Nums) {
		for Nums[i - 1] != i {
			Nums[i - 1], Nums[Nums[i - 1] - 1] = Nums[Nums[i - 1] - 1], Nums[i - 1]
		}
		i++
	}
	return Nums
}
