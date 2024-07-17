package datastructures

type Trie struct {
	Root *TrieNode
}

type TrieNode struct {
	Children map[string]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{Children: map[string]*TrieNode{}}
}

func NewTrie() *Trie {
	return &Trie{Root: NewTrieNode()}
}

func (trie *Trie) Search(word string) bool {
	current := trie.Root

	for i := range word {
		char := string(word[i])
		node, ok := current.Children[char]
		if ok {
			current = node
		} else {
			return false
		}
	}
	return true
}

func (trie *Trie) Insert(word string) {
	current := trie.Root

	for i := range word {
		char := string(word[i])
		node, ok := current.Children[char]
		if ok {
			current = node
		} else {
			node := NewTrieNode()
			current.Children[char] = node
			current = node
		}
	}

	current.Children["*"] = nil
}
