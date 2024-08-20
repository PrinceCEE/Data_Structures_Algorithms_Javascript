package palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	assert.Equal(t, true, IsPalindrome(s))

	s = "A man, a plan, a canal: Panam"
	assert.Equal(t, false, IsPalindrome(s))
}
