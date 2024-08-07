package groupanagrams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := GroupAnagrams(strs)
	assert.Equal(t, 3, len(res))

	strs = []string{""}
	res = GroupAnagrams(strs)
	assert.Equal(t, 1, len(res))
	assert.Equal(t, "", res[0][0])

	strs = []string{"a"}
	res = GroupAnagrams(strs)
	assert.Equal(t, 1, len(res))
	assert.Equal(t, "a", res[0][0])
}
