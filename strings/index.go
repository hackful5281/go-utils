package strings

// In int 包含
func In(target int, nums []int) bool {
	for _, v := range nums {
		if target == v {
			return true
		}
	}
	return false
}
