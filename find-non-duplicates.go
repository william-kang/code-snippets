import (
    "fmt"
)

type Solution struct {
}

func (sol *Solution) moveElements(arr []int) int {
    i := 1
    for i < len(arr) && arr[i - 1] < arr[i] {
        i++
    }
    j := i + 1
    for {
        for j < len(arr) && arr[j - 1] == arr[j] {
            j++
        }
        if j < len(arr) {
            arr[i] = arr[j]
            i++
            j++
        } else {
            break
        }
    }
    fmt.Println(arr)
    return i
}
