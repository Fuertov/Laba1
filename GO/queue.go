package main

import "fmt"

type FifoQueue[T any] struct {
	Head *SNode[T]
	Tail *SNode[T]
}

func (q *FifoQueue[T]) QIsEmpty() bool { return q.Head == nil }

func (q *FifoQueue[T]) QAdd(v T) {
	n := &SNode[T]{Value: v}
	if q.Tail == nil {
		q.Head = n
		q.Tail = n
	} else {
		q.Tail.Next = n
		q.Tail = n
	}
}

func (q *FifoQueue[T]) QRemove() {
	if q.QIsEmpty() {
		fmt.Println("ОШИБКА: Очередь пуста, удалять нечего")
		return
	}
	q.Head = q.Head.Next
	if q.Head == nil {
		q.Tail = nil
	}
}

func (q *FifoQueue[T]) QPeek() T {
	var zero T
	if q.QIsEmpty() {
		fmt.Println("ОШИБКА: Очередь пуста")
		return zero
	}
	return q.Head.Value
}

func (q *FifoQueue[T]) QShow() {
	fmt.Print("Q[")
	for c := q.Head; c != nil; c = c.Next {
		fmt.Printf("%v ", c.Value)
	}
	fmt.Println("]")
}

func (q *FifoQueue[T]) Clear() { q.Head = nil; q.Tail = nil }
