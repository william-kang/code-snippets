// Solution struct type
type Solution struct{}

// containsDuplicate checks for duplicates in a slice of integers
func (s *Solution) containsDuplicate(nums []int) bool {
    // ToDo: Write Your Code Here.
    m := make(map[int]int)
    for _, v := range nums {
        _, ok := m[v]
        if ok {
            return true
        } else {
            m[v] = 1
        }
    }
    return false
}
