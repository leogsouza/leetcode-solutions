package lc04

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lnums1, lnums2 := len(nums1), len(nums2)
	slc := make([]int, 0, lnums1+lnums2)
	i, j := 0, 0
	for i < lnums1 && j < lnums2 {
		n1 := nums1[i]
		n2 := nums2[j]
		if n1 <= n2 {
			slc = append(slc, n1)
			i++
		} else {
			slc = append(slc, n2)
			j++
		}
	}
	if i < lnums1 {
		slc = append(slc, nums1[i:]...)
	} else {
		slc = append(slc, nums2[j:]...)
	}
	c := lnums1 + lnums2
	if c&1 == 1 {
		t := (c - 1) / 2
		return float64(slc[t])
	}
	t := c / 2
	return (float64(slc[t-1]) + float64(slc[t])) / 2
}
