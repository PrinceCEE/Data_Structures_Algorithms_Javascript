package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearchTree(t *testing.T) {
	bst := BST{}

	assert.Nil(t, bst.Root)

	values := []int{2, 1, 5, 10, 3, 14, 4}

	t.Run("Insert into bst", func(t *testing.T) {
		var root *TreeNode
		for i, v := range values {
			if i == 0 {
				root = bst.Insert(nil, v)
				continue
			}

			bst.Insert(root, v)
		}

		arr := []int{}
		bst.TraverseInOrder(bst.Root, &arr)
		assert.Equal(t, 1, arr[0])
		assert.Equal(t, 14, arr[len(arr)-1])
	})

	t.Run("Search bst", func(t *testing.T) {
		node := bst.Search(bst.Root, 14)
		assert.NotNil(t, node)
		assert.Equal(t, 14, node.Value)

		node = bst.Search(bst.Root, 15)
		assert.Nil(t, node)
	})

	t.Run("Delete node", func(t *testing.T) {
		bst.Delete(bst.Root, 10)
		bst.Delete(bst.Root, 14)

		arr := []int{}
		bst.TraverseInOrder(bst.Root, &arr)
		for _, v := range arr {
			assert.NotEqual(t, 10, v)
			assert.NotEqual(t, 14, v)
		}
	})
}
