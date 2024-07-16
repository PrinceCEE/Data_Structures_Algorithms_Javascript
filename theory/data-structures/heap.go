package datastructures

// Heaps implementations use arrays for ease of getting and inserting the last data.
// To get a node's left child, we use (i x 2) + 1, while we use (i x 2) + 2 for the
// right child.
// To get a node's parent, we use (i - 1)/2 using integer division
type Heap struct {
	Data []int
}

func NewHeap() *Heap {
	return &Heap{Data: []int{}}
}

func (h *Heap) GetRootNode() int {
	return h.Data[0]
}

func (h *Heap) GetLastNode() int {
	return h.Data[len(h.Data)-1]
}

func (h *Heap) GetLeftChildIndex(index int) int {
	return (2 * index) + 1
}

func (h *Heap) GetRightChildIndex(index int) int {
	return (2 * index) + 2
}

func (h *Heap) GetParentIndex(index int) int {
	return (index - 1) / 2
}

// Appends the data to the end of the array, and trickles it up
// its correct position by iteratively check it against its parent and swapping
func (h *Heap) Insert(data int) {
	h.Data = append(h.Data, data)

	newNodeIndex := len(h.Data) - 1

	for newNodeIndex > 0 && h.Data[newNodeIndex] > h.Data[h.GetParentIndex(newNodeIndex)] {
		h.Data[newNodeIndex], h.Data[h.GetParentIndex(newNodeIndex)] =
			h.Data[h.GetParentIndex(newNodeIndex)], h.Data[newNodeIndex]

		newNodeIndex = h.GetParentIndex(newNodeIndex)
	}
}

// Replaces the root with the last element in the array,
// then trickle the value down to its right position by
// iteratively swapping it with its child that is greater
func (h *Heap) Delete() int {
	n := len(h.Data)
	if n == 0 {
		return -1
	}

	max := h.Data[0]
	h.Data[0] = h.Data[n-1]
	h.Data = h.Data[:n-1]

	index := 0
	for h.HasGreaterChild(index) {
		largerIndex := h.GetLargerIndex(index)
		h.Data[index], h.Data[largerIndex] = h.Data[largerIndex], h.Data[index]
		index = largerIndex
	}

	return max
}

func (h *Heap) HasGreaterChild(index int) bool {
	l := h.GetLeftChildIndex(index)
	r := h.GetRightChildIndex(index)
	n := len(h.Data)

	if l <= n-1 && h.Data[l] > h.Data[index] {
		return true
	}

	if r <= n-1 && h.Data[r] > h.Data[index] {
		return true
	}

	return false
}

func (h *Heap) GetLargerIndex(index int) int {
	l := h.GetLeftChildIndex(index)
	r := h.GetRightChildIndex(index)
	n := len(h.Data)

	if l <= n-1 && h.Data[l] > h.Data[index] && r <= n-1 && h.Data[l] > h.Data[r] {
		return l
	}

	return r
}
