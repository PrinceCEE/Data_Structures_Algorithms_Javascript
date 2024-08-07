package groupanagrams

import (
	"sort"
	"strings"
)

// Given an array of strings strs, group the anagrams together.
// You can return the answer in any order.
func GroupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, str := range strs {
		sortedStr := SortStr(str)
		m[sortedStr] = append(m[sortedStr], str)
	}

	res := [][]string{}
	for _, v := range m {
		res = append(res, v)
	}

	return res
}

func SortStr(str string) string {
	split := strings.Split(str, "")
	sort.Strings(split)
	return strings.Join(split, "")
}

// Time complexity is O(n)
// Space complexity is O(n)
