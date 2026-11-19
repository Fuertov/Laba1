package main

import "fmt"

type DNode[T any] struct {
	Value T
	Next  *DNode[T]
	Prev  *DNode[T]
}

type DoublyList[T any] struct {
	Head *DNode[T]
	Tail *DNode[T]
}

func (l *DoublyList[T]) LAddHead(v T) {
	n := &DNode[T]{Value: v, Next: l.Head}
	if l.Head != nil {
		l.Head.Prev = n
	} else {
		l.Tail = n
	}
	l.Head = n
}

func (l *DoublyList[T]) LAddTail(v T) {
	n := &DNode[T]{Value: v, Prev: l.Tail}
	if l.Tail != nil {
		l.Tail.Next = n
	} else {
		l.Head = n
	}
	l.Tail = n
}

func (l *DoublyList[T]) LAddBefore(node *DNode[T], v T) {
	if node == nil || l.Head == node {
		l.LAddHead(v)
		return
	}
	cur := l.Head
	for cur != nil && cur.Next != node {
		cur = cur.Next
	}
	if cur != nil {
		n := &DNode[T]{Value: v, Next: node, Prev: node.Prev}
		node.Prev.Next = n
		node.Prev = n
	} else {
		fmt.Printf("ОШИБКА: Узла %v нет в списке\n", node.Value)
	}
}

func (l *DoublyList[T]) LAddAfter(node *DNode[T], v T) {
	if node == nil {
		l.LAddTail(v)
		return
	}
	cur := l.Head
	for cur != nil && cur != node {
		cur = cur.Next
	}
	if cur != nil {
		n := &DNode[T]{Value: v, Next: node.Next, Prev: node}
		if node.Next != nil {
			node.Next.Prev = n
		} else {
			l.Tail = n
		}
		node.Next = n
	} else {
		fmt.Printf("ОШИБКА: Узла %v нет в списке\n", node.Value)
	}
}

func (l *DoublyList[T]) LRemoveHead() {
	if l.Head == nil {
		fmt.Println("ОШИБКА: В списке нет элементов, удалять нечего")
		return
	}
	l.Head = l.Head.Next
	if l.Head != nil {
		l.Head.Prev = nil
	} else {
		l.Tail = nil
	}
}

func (l *DoublyList[T]) LRemoveTail() {
	if l.Tail == nil {
		fmt.Println("ОШИБКА: В списке нет элементов, удалять нечего")
		return
	}
	l.Tail = l.Tail.Prev
	if l.Tail != nil {
		l.Tail.Next = nil
	} else {
		l.Head = nil
	}
}

func (l *DoublyList[T]) LRemoveNode(node *DNode[T]) {
	if node == nil {
		fmt.Println("ОШИБКА: Узел не может быть nil")
		return
	}
	cur := l.Head
	for cur != nil && cur != node {
		cur = cur.Next
	}
	if cur == nil {
		fmt.Printf("ОШИБКА: Узла %v нет в списке\n", node.Value)
		return
	}
	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		l.Head = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		l.Tail = node.Prev
	}
}

func (l *DoublyList[T]) LRemoveAfter(node *DNode[T]) {
	if node == nil {
		fmt.Println("ОШИБКА: Узел не может быть nil")
		return
	}
	if node == l.Tail {
		fmt.Println("ОШИБКА: За последним элементом, нет элементов для удаления")
		return
	}
	cur := l.Head
	for cur != nil && cur != node {
		cur = cur.Next
	}
	if cur == nil {
		fmt.Printf("ОШИБКА: Узла %v нет в списке\n", node.Value)
		return
	}
	if cur.Next == l.Tail {
		l.LRemoveTail()
	} else {
		cur.Next = cur.Next.Next
		cur.Next.Prev = cur
	}
}

func (l *DoublyList[T]) LRemoveBefore(node *DNode[T]) {
	if node == nil {
		fmt.Println("ОШИБКА: Узел не может быть nil")
		return
	}
	if node == l.Head {
		fmt.Println("ОШИБКА: Перед первым элементом, нет элементов для удаления")
		return
	}
	cur := l.Head
	for cur != nil && cur != node {
		cur = cur.Next
	}
	if cur == nil {
		fmt.Printf("ОШИБКА: Узла %v нет в списке\n", node.Value)
		return
	}
	if cur.Prev == l.Head {
		l.LRemoveHead()
	} else {
		cur.Prev = cur.Prev.Prev
		cur.Prev.Next = cur
	}
}

func (l *DoublyList[T]) LRemoveValue(v T) {
	cur := l.Head
	for cur != nil && any(cur.Value) != any(v) {
		cur = cur.Next
	}
	if cur != nil {
		l.LRemoveNode(cur)
	} else {
		fmt.Printf("ОШИБКА: Значения %v нет в списке\n", v)
	}
}

func (l *DoublyList[T]) LSearch(v T) {
	i := 0
	for c := l.Head; c != nil; c = c.Next {
		if any(c.Value) == any(v) {
			fmt.Printf("Элемент был найден на позиции %d\n", i)
			return
		}
		i++
	}
	fmt.Println("Элемент не был найден в списке")
}

func (l *DoublyList[T]) LShow() {
	fmt.Print("L[")
	for c := l.Head; c != nil; c = c.Next {
		fmt.Printf("%v ", c.Value)
	}
	fmt.Println("]")
}

func (l *DoublyList[T]) LShowReverse() {
	fmt.Print("Lr[")
	for c := l.Tail; c != nil; c = c.Prev {
		fmt.Printf("%v ", c.Value)
	}
	fmt.Println("]")
}

func (l *DoublyList[T]) LShowRec(n *DNode[T]) {
	if n == nil {
		return
	}
	fmt.Printf("%v ", n.Value)
	l.LShowRec(n.Next)
}

func (l *DoublyList[T]) Clear() { l.Head = nil; l.Tail = nil }
