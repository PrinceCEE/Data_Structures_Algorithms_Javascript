package validanagram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {
	assert.Equal(t, true, IsAnagram("listen", "silent"))
	assert.Equal(t, false, IsAnagram("car", "rat"))
}
