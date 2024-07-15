package datastructures

type Node[T comparable] struct {
	Data T
	Next *Node[T]
}

type LinkedList[T comparable] struct {
	Head *Node[T]
}

func (list *LinkedList[T]) Append(data T) {
	node := &Node[T]{Data: data}

	if list.Head == nil {
		list.Head = node
		return
	}

	current := list.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = node
}

func (list *LinkedList[T]) Prepend(data T) {
	node := &Node[T]{Data: data}

	node.Next = list.Head
	list.Head = node
}

func (list *LinkedList[T]) DeleteWithValue(data T) {
	if list.Head == nil {
		return
	}

	if list.Head.Data == data {
		list.Head = list.Head.Next
		return
	}

	current := list.Head
	for current.Next != nil {
		if current.Next.Data == data {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

func (list *LinkedList[T]) Find(data T) *Node[T] {
	current := list.Head
	for current != nil {
		if current.Data == data {
			return current
		}
		current = current.Next
	}
	return nil
}

func (list *LinkedList[T]) Length() int {
	length := 0
	current := list.Head
	for current != nil {
		length++
		current = current.Next
	}
	return length
}

func (list *LinkedList[T]) Reverse() {
	var prev, next *Node[T]

	current := list.Head
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	list.Head = prev
}
