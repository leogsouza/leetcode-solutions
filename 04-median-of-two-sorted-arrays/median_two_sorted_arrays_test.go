package lc04

import "testing"

type testCase struct {
	nums1    []int
	nums2    []int
	expected float64
}

func TestFindMedianSortedArrays(t *testing.T) {
	data := [...]testCase{
		{
			nums1:    []int{1, 3},
			nums2:    []int{2},
			expected: 2.0,
		},
		{
			nums1:    []int{1, 2},
			nums2:    []int{3, 4},
			expected: 2.50000,
		},
		{
			nums1:    []int{0, 0},
			nums2:    []int{0, 0},
			expected: 0.0,
		},
		{
			nums1:    []int{},
			nums2:    []int{1},
			expected: 1.0,
		},
	}

	for _, tt := range data {
		got := findMedianSortedArrays(tt.nums1, tt.nums2)
		if got != tt.expected {
			t.Fatalf("input: %v %v, got: %v, expected: %v", tt.nums1, tt.nums2, got, tt.expected)
		}
	}
}
