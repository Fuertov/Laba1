package main

import "fmt"

type LifoStack[T any] struct {
	Head *SNode[T]
}

func (s *LifoStack[T]) SIsEmpty() bool { return s.Head == nil }

func (s *LifoStack[T]) SAdd(v T) {
	s.Head = &SNode[T]{Value: v, Next: s.Head}
}

func (s *LifoStack[T]) SRemove() {
	if s.SIsEmpty() {
		fmt.Println("ОШИБКА: Стек пуст, удалять нечего")
		return
	}
	s.Head = s.Head.Next
}

func (s *LifoStack[T]) SPeek() T {
	var zero T
	if s.SIsEmpty() {
		fmt.Println("ОШИБКА: Стек пуст")
		return zero
	}
	return s.Head.Value
}

func (s *LifoStack[T]) SShow() {
	fmt.Print("S[")
	for c := s.Head; c != nil; c = c.Next {
		fmt.Printf("%v ", c.Value)
	}
	fmt.Println("]")
}

func (s *LifoStack[T]) Clear() { s.Head = nil }
