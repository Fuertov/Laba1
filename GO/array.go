package main

import "fmt"

type DynArray[T any] struct {
	Data     []T
	Capacity int
}

func (a *DynArray[T]) Init() {
	a.Capacity = 1
	a.Data = make([]T, 0, a.Capacity)
}

func (a *DynArray[T]) MAppend(n T) {
	if len(a.Data) == a.Capacity {
		a.Capacity *= 2
	}
	a.Data = append(a.Data, n)
}

func (a *DynArray[T]) MInsertAt(idx int, n T) {
	if idx < 0 || idx > len(a.Data) {
		fmt.Printf("ОШИБКА: индекс %d не существует\n", idx)
		return
	}
	a.Data = append(a.Data, *new(T))
	copy(a.Data[idx+1:], a.Data[idx:])
	a.Data[idx] = n
}

func (a *DynArray[T]) MGet(idx int) T {
	var zero T
	if idx < 0 || idx >= len(a.Data) {
		fmt.Printf("ОШИБКА: индекс %d не существует\n", idx)
		return zero
	}
	return a.Data[idx]
}

func (a *DynArray[T]) MRemoveAt(idx int) {
	if idx < 0 || idx >= len(a.Data) {
		fmt.Printf("ОШИБКА: индекс %d не существует\n", idx)
		return
	}
	a.Data = append(a.Data[:idx], a.Data[idx+1:]...)
}

func (a *DynArray[T]) MReplaceAt(idx int, v T) {
	if idx < 0 || idx >= len(a.Data) {
		fmt.Printf("ОШИБКА: индекс %d не существует\n", idx)
		return
	}
	a.Data[idx] = v
}

func (a *DynArray[T]) MSize() int { return len(a.Data) }

func (a *DynArray[T]) MShow() {
	fmt.Print("M[")
	for _, v := range a.Data {
		fmt.Printf("%v ", v)
	}
	fmt.Println("]")
}

func (a *DynArray[T]) Clear() {
	a.Data = nil
	a.Capacity = 0
}
