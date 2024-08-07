package datastructures

import (
	"errors"
)

var (
	ErrEmptyHeap = errors.New("heap is empty")
)

// Heaps are binary trees that are both complete and partially ordered in a certain way
// A complete binary tree is a binary tree that has all the levels filled with the exception
// of the last level which must be filled from the left
//
// Heaps are divided into Min/Max heap primarily. The min heap must have the node's value
// always less than its children, while the max node must have the node's value always
// greater than the its children
//
// The min and max heaps are implemented the same way, but only differ in their value
// comparisons.
type MinHeap struct {
	Data []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{Data: []int{}}
}

func (h *MinHeap) Size() int {
	return len(h.Data)
}

func (h *MinHeap) Peek() (int, error) {
	if h.Size() == 0 {
		return 0, ErrEmptyHeap
	}
	return h.Data[0], nil
}

func (h *MinHeap) Add(i int) {
	h.Data = append(h.Data, i)
	h.heapifyUp()
}

func (h *MinHeap) Pop() (int, error) {
	if h.Size() == 0 {
		return 0, ErrEmptyHeap
	}

	value := h.Data[0]
	h.Data[0] = h.Data[h.Size()-1]
	h.Data = h.Data[:h.Size()-1]
	h.heapifyDown()

	return value, nil
}

func (h *MinHeap) swap(i, j int) {
	h.Data[i], h.Data[j] = h.Data[j], h.Data[i]
}

func (h *MinHeap) heapifyUp() {
	i := h.Size() - 1
	for h.hasParent(i) {
		pi := h.getParentIndex(i)
		if h.Data[i] < h.Data[pi] {
			h.swap(i, pi)
		}
		i = pi
	}
}

func (h *MinHeap) heapifyDown() {
	i := 0
	for h.hasLeftChild(i) {
		smallestIdx := h.getLeftChildIndex(i)
		if h.hasRightChild(i) && h.Data[h.getRightChildIndex(i)] < h.Data[smallestIdx] {
			smallestIdx = h.getRightChildIndex(i)
		}

		if h.Data[i] <= h.Data[smallestIdx] {
			break
		} else {
			h.swap(i, smallestIdx)
			i = smallestIdx
		}
	}
}

func (h *MinHeap) getLeftChildIndex(i int) int {
	return 2*i + 1
}

func (h *MinHeap) getRightChildIndex(i int) int {
	return 2*i + 2
}

func (h *MinHeap) getParentIndex(i int) int {
	return (i - 1) / 2
}

func (h *MinHeap) hasLeftChild(i int) bool {
	return h.getLeftChildIndex(i) < len(h.Data)
}

func (h *MinHeap) hasRightChild(i int) bool {
	return h.getRightChildIndex(i) < len(h.Data)
}

func (h *MinHeap) hasParent(i int) bool {
	if i == 0 {
		return false
	}

	parentIdx := h.getParentIndex(i)
	return parentIdx >= 0
}
