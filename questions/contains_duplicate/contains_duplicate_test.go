package containsduplicate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type intSlice []int

func TestContainsDuplicate(t *testing.T) {
	withDuplicate := intSlice{1, 2, 1, 3, 4}
	noDuplicate := intSlice{1, 2, 3, 4, 5}

	assert.Equal(t, true, ContainsDuplicate((withDuplicate)))
	assert.Equal(t, false, ContainsDuplicate(noDuplicate))
}
