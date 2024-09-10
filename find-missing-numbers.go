import (
    "fmt"
)

type Solution struct{}

func (s *Solution) findNumbers(nums []int) []int {
    for i := 1; i <= len(nums); i++ {
        fmt.Println(nums)
        for nums[i - 1] != i && nums[nums[i - 1] - 1] != nums[i - 1] {
            nums[nums[i - 1] - 1], nums[i - 1] =
                nums[i - 1], nums[nums[i - 1] - 1]
        }
    }
    ans := make([]int, 0)
    for i := 0; i < len(nums); i++ {
        if nums[i] != i + 1 {
            ans = append(ans, i + 1)
        }
    }
    return ans
}
