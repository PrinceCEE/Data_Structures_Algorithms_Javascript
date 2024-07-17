package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()

	words := []string{"ace", "act", "sponge", "bob"}
	for i := range words {
		trie.Insert(words[i])
	}

	for i := range words {
		word := words[i]
		isFound := trie.Search(word)
		assert.Equal(t, true, isFound)
	}

	isFound := trie.Search("acts")
	assert.Equal(t, false, isFound)
}
