import ()

// Solution struct to mimic a class.
type Solution struct{}

func (s* Solution) makeSquares(arr []int) []int {
    squares := make([]int, len(arr))
    i, j := 0, len(arr) - 1
    n := len(arr) - 1
    for i <= j {
        if arr[i]*arr[i] > arr[j]*arr[j] {
            squares[n] = arr[i]*arr[i]
            i++
            n--
        } else {
            squares[n] = arr[j]*arr[j]
            j--
            n--
        }
    }
    return squares
}
