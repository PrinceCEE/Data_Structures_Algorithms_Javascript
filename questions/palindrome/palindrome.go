package palindrome

import "strings"

// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters
// and removing all non-alphanumeric characters, it reads the same forward and backward.
// Alphanumeric characters include letters and numbers.
// Given a string s, return true if it is a palindrome, or false otherwise.
func IsPalindrome(str string) bool {
	l, r := 0, len(str)-1

	for l < r {
		for l < r && !IsAlphaNum(rune(str[l])) {
			l++
		}

		for l < r && !IsAlphaNum(rune(str[r])) {
			r--
		}

		if !strings.EqualFold(string(str[l]), string(str[r])) {
			return false
		}

		l, r = l+1, r-1
	}

	return true
}

func IsAlphaNum(char rune) bool {
	return ('0' <= char && char <= '9') || ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z')
}
