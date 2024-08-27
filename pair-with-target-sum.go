
import ()
type Solution struct {
}
// search - finds a pair in arr with a given targetSum.
func (sol *Solution) search(arr []int, targetSum int) []int {
    i, j := 0, len(arr) - 1
    for i < j {
      if sum := arr[i] + arr[j]; sum > targetSum {
        j--
      } else if sum < targetSum {
        i++
      } else {
        return []int{i, j}
      }
    }
    return []int{-1, -1}
}
