package datastructures

type Node[T comparable] struct {
	data T
	next *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
}

func (list *LinkedList[T]) append(data T) {
	node := &Node[T]{data: data}

	if list.head == nil {
		list.head = node
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = node
}

func (list *LinkedList[T]) prepend(data T) {
	node := &Node[T]{data: data}

	node.next = list.head
	list.head = node
}

func (list *LinkedList[T]) deleteWithValue(data T) {
	if list.head == nil {
		return
	}

	if list.head.data == data {
		list.head = list.head.next
		return
	}

	current := list.head
	for current.next != nil {
		if current.next.data == data {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}

func (list *LinkedList[T]) find(data T) *Node[T] {
	current := list.head
	for current != nil {
		if current.data == data {
			return current
		}
		current = current.next
	}
	return nil
}

func (list *LinkedList[T]) length() int {
	length := 0
	current := list.head
	for current != nil {
		length++
		current = current.next
	}
	return length
}

func (list *LinkedList[T]) reverse() {
	var prev, next *Node[T]

	current := list.head
	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}

	list.head = prev
}
