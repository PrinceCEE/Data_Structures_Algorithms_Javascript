package validanagram

// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
// typically using all the original letters exactly once.
func IsAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := make(map[rune]int)
	tMap := make(map[rune]int)

	PopulateMap(sMap, s)
	PopulateMap(tMap, t)

	for k, v := range sMap {
		if v != tMap[k] {
			return false
		}
	}

	return true
}

func PopulateMap(m map[rune]int, str string) {
	for _, char := range str {
		m[char]++
	}
}

// Time Complexity is O(n)
// Space Complexity is O(n)
