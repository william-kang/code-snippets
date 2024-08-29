import (
    "sort"
)

type Solution struct{}

func (s *Solution) searchTriplets(arr []int) [][]int {
    triplets := make([][]int, 0)
    sort.Ints(arr)
    for i := 0; i < len(arr) - 2; i++ {
        j, k := i + 1, len(arr) - 1
        for j < k {
            if arr[j] + arr[k] == -arr[i] {
                triplet := []int{arr[i], arr[j], arr[k]}
                triplets = append(triplets, triplet)
                for j < len(arr) - 1 && arr[j] == arr[j + 1] { j++ }
                j++
            } else if arr[j] + arr[k] < -arr[i] {
                j++
            } else {
                k--
            }
        }
        for i < len(arr) - 1 && arr[i] == arr[i + 1] { i++ }
    }
    return triplets
}
