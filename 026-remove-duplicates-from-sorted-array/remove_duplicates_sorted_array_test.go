package lc26

import "testing"

type testCase struct {
	name string
	nums []int
	want int
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []testCase{
		{
			name: "should return 2 for array [1,1,2]",
			nums: []int{1, 1, 2},
			want: 2,
		},
		{
			name: "should return 5 for array [0,0,1,1,1,2,2,3,3,4]",
			nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.nums)

			if got != tt.want {
				t.Errorf("test: %s failed. got %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
