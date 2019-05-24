package twosum

import (
	"reflect"
	"testing"
)

type testCase struct {
	nums     []int
	target   int
	expected []int
}

func TestTwoSum(t *testing.T) {
	data := [...]testCase{
		{
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			nums:     []int{2, 9, 13, 22},
			target:   5,
			expected: nil,
		},
		{
			nums:     []int{3, 4, -4, 2},
			target:   0,
			expected: []int{1, 2},
		},
	}

	for _, tt := range data {
		got := twoSum(tt.nums, tt.target)
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("twoSum input %v expected %d but got %v", tt.nums, tt.expected, got)
		}
	}
}
