package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	t.Run("test append", func(t *testing.T) {
		list := &LinkedList[int]{}
		assert.Nil(t, list.head)

		list.append(1)
		list.append(2)
		list.append(3)

		assert.Equal(t, 1, list.head.data)
		assert.Equal(t, 2, list.head.next.data)
		assert.Equal(t, 3, list.head.next.next.data)
	})

	t.Run("test prepend", func(t *testing.T) {
		list := &LinkedList[int]{}
		assert.Nil(t, list.head)

		list.prepend(1)
		list.prepend(2)
		list.prepend(3)

		assert.Equal(t, 3, list.head.data)
		assert.Equal(t, 2, list.head.next.data)
		assert.Equal(t, 1, list.head.next.next.data)
	})

	t.Run("test delete with value", func(t *testing.T) {
		list := &LinkedList[int]{}
		assert.Nil(t, list.head)

		list.append(1)
		list.append(2)
		list.append(3)
		list.deleteWithValue(2)

		assert.Equal(t, 3, list.head.next.data)

		list.deleteWithValue(1)
		assert.Equal(t, 3, list.head.data)
	})

	t.Run("test find", func(t *testing.T) {
		list := &LinkedList[int]{}
		assert.Nil(t, list.head)

		list.append(1)
		list.append(2)
		list.append(3)

		node := list.find(2)
		assert.NotNil(t, node)
		assert.Equal(t, 2, node.data)

		node = list.find(4)
		assert.Nil(t, node)
	})

	t.Run("test length", func(t *testing.T) {
		list := &LinkedList[int]{}
		assert.Nil(t, list.head)

		list.append(1)
		list.append(2)
		list.append(3)

		assert.Equal(t, 3, list.length())

		list.deleteWithValue(2)
		assert.Equal(t, 2, list.length())
	})

	t.Run("test reverse", func(t *testing.T) {
		list := &LinkedList[int]{}
		assert.Nil(t, list.head)

		list.append(1)
		list.append(2)
		list.append(3)

		list.reverse()
		assert.NotNil(t, list.head)
		assert.Equal(t, 3, list.head.data)
		assert.Equal(t, 2, list.head.next.data)
		assert.Equal(t, 1, list.head.next.next.data)
	})
}
