package lc125

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		data string
		want bool
	}{
		{
			name: "should return true for string 'abba'",
			data: "abba",
			want: true,
		},
		{
			name: "should return false for string 'love'",
			data: "love",
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := isPalindrome(test.data)
			if got != test.want {
				t.Errorf("test: %v failed. got %v, want %v", test.name, got, test.want)
			}
		})
	}
}
