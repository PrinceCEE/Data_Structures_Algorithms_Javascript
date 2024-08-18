package stringencodedecode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringEncodeDecode(t *testing.T) {
	arr := []string{"neet", "code", "love", "you"}
	str := "4#neet4#code4#love3#you"

	resStr := Encode(arr)
	assert.Equal(t, str, resStr)

	resArr := Decode(resStr)
	for i := range resArr {
		assert.Equal(t, arr[i], resArr[i])
	}
}
