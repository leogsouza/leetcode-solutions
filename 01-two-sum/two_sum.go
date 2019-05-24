package twosum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for k, v := range nums {
		if val, ok := m[target-v]; ok {
			return []int{val, k}
		}
		m[v] = k
	}
	return nil
}
