package stringencodedecode

import (
	"fmt"
	"strconv"
)

// Design an algorithm to encode a list of strings to a single string.
// The encoded string is then decoded back to the original list of strings.

func Encode(arr []string) string {
	res := ""

	for _, str := range arr {
		res += fmt.Sprintf("%d#%s", len(str), str)
	}

	return res
}

func Decode(str string) []string {
	res := []string{}
	i := 0

	for i < len(str) {
		j := i
		for str[j] != '#' {
			j++
		}

		strLen, err := strconv.Atoi(str[i:j])
		if err != nil {
			panic(err)
		}

		res = append(res, str[j+1:j+1+strLen])
		i = j + 1 + strLen
	}

	return res
}

// Time and space complexities are O(n)
