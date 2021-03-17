package lc26

func removeDuplicates(nums []int) int {
	i, j, k := 0, 1, len(nums)

	for ; j < k; j++ {
		if nums[i] == nums[j] {
			continue
		}
		i++
		nums[i], nums[j] = nums[j], nums[i]
	}

	return i + 1
}
