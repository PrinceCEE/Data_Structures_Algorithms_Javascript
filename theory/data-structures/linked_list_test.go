package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	t.Run("test append", func(t *testing.T) {
		list := &LinkedList[int]{}
		assert.Nil(t, list.Head)

		list.Append(1)
		list.Append(2)
		list.Append(3)

		assert.Equal(t, 1, list.Head.Data)
		assert.Equal(t, 2, list.Head.Next.Data)
		assert.Equal(t, 3, list.Head.Next.Next.Data)
	})

	t.Run("test prepend", func(t *testing.T) {
		list := &LinkedList[int]{}
		assert.Nil(t, list.Head)

		list.Prepend(1)
		list.Prepend(2)
		list.Prepend(3)

		assert.Equal(t, 3, list.Head.Data)
		assert.Equal(t, 2, list.Head.Next.Data)
		assert.Equal(t, 1, list.Head.Next.Next.Data)
	})

	t.Run("test delete with value", func(t *testing.T) {
		list := &LinkedList[int]{}
		assert.Nil(t, list.Head)

		list.Append(1)
		list.Append(2)
		list.Append(3)
		list.DeleteWithValue(2)

		assert.Equal(t, 3, list.Head.Next.Data)

		list.DeleteWithValue(1)
		assert.Equal(t, 3, list.Head.Data)
	})

	t.Run("test find", func(t *testing.T) {
		list := &LinkedList[int]{}
		assert.Nil(t, list.Head)

		list.Append(1)
		list.Append(2)
		list.Append(3)

		node := list.Find(2)
		assert.NotNil(t, node)
		assert.Equal(t, 2, node.Data)

		node = list.Find(4)
		assert.Nil(t, node)
	})

	t.Run("test length", func(t *testing.T) {
		list := &LinkedList[int]{}
		assert.Nil(t, list.Head)

		list.Append(1)
		list.Append(2)
		list.Append(3)

		assert.Equal(t, 3, list.Length())

		list.DeleteWithValue(2)
		assert.Equal(t, 2, list.Length())
	})

	t.Run("test reverse", func(t *testing.T) {
		list := &LinkedList[int]{}
		assert.Nil(t, list.Head)

		list.Append(1)
		list.Append(2)
		list.Append(3)

		list.Reverse()
		assert.NotNil(t, list.Head)
		assert.Equal(t, 3, list.Head.Data)
		assert.Equal(t, 2, list.Head.Next.Data)
		assert.Equal(t, 1, list.Head.Next.Next.Data)
	})
}
