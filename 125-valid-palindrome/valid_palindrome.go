package lc125

// Given a string s, determine if it is a palindrome, considering only
// alphanumeric characters and ignoring cases.

// Ex1: "A man, a plan, a canal: Panama" should return true
// because "amanaplanacanalpanama" is a palindrome.

// Ex2: "race a car" should return false
// because "raceacar" is not palindrome

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {

		if !isAlpha(s[l]) {
			l++
			continue
		}
		if !isAlpha(s[r]) {
			r--
			continue
		}

		a := lowerCase(s[l])

		b := lowerCase(s[r])

		if a != b {
			return false
		}

		l++
		r--
	}

	return true
}

func isAlpha(b byte) bool {
	return (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9')

}

// if b = 'C' b = 67 + 97 - 65 = 99 => 'c'
func lowerCase(b byte) byte {
	if b >= 'A' && b <= 'Z' {

		return b + 'a' - 'A'
	}

	return b
}
