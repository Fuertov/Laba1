package main

import "fmt"

type SNode[T any] struct {
	Value T
	Next  *SNode[T]
}

type SinglyList[T any] struct {
	Head *SNode[T]
}

func (l *SinglyList[T]) FAddHead(v T) {
	l.Head = &SNode[T]{Value: v, Next: l.Head}
}

func (l *SinglyList[T]) FAddTail(v T) {
	n := &SNode[T]{Value: v}
	if l.Head == nil {
		l.Head = n
		return
	}
	cur := l.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = n
}

func (l *SinglyList[T]) FAddAfter(node *SNode[T], v T) {
	if node == nil {
		l.FAddTail(v)
		return
	}
	cur := l.Head
	for cur != nil && cur != node {
		cur = cur.Next
	}
	if cur != nil {
		cur.Next = &SNode[T]{Value: v, Next: cur.Next}
	} else {
		fmt.Printf("ОШИБКА: Узла %v нет в списке\n", node.Value)
	}
}

func (l *SinglyList[T]) FAddBefore(node *SNode[T], v T) {
	if l.Head == node || node == nil {
		l.FAddHead(v)
		return
	}
	cur := l.Head
	for cur != nil && cur.Next != node {
		cur = cur.Next
	}
	if cur != nil {
		cur.Next = &SNode[T]{Value: v, Next: cur.Next}
	} else {
		fmt.Printf("ОШИБКА: Узла %v нет в списке\n", node.Value)
	}
}

func (l *SinglyList[T]) FRemoveHead() {
	if l.Head == nil {
		fmt.Println("ОШИБКА: В списке нет элементов, удалять нечего")
		return
	}
	l.Head = l.Head.Next
}

func (l *SinglyList[T]) FRemoveTail() {
	if l.Head == nil {
		fmt.Println("ОШИБКА: В списке нет элементов, удалять нечего")
		return
	}
	if l.Head.Next == nil {
		l.Head = nil
		return
	}
	cur := l.Head
	for cur.Next.Next != nil {
		cur = cur.Next
	}
	cur.Next = nil
}

func (l *SinglyList[T]) FRemoveAfter(node *SNode[T]) {
	if node == nil {
		fmt.Println("ОШИБКА: Узел не может быть nil")
		return
	}
	cur := l.Head
	for cur != nil && cur != node {
		cur = cur.Next
	}
	if cur != nil {
		if cur.Next == nil {
			fmt.Printf("ОШИБКА: После узла %v ничего нет\n", node.Value)
			return
		}
		cur.Next = cur.Next.Next
	} else {
		fmt.Printf("ОШИБКА: Узла %v нет в списке\n", node.Value)
	}
}

func (l *SinglyList[T]) FRemoveBefore(node *SNode[T]) {
	if node == nil {
		fmt.Println("ОШИБКА: Узел не может быть nil")
		return
	}
	if l.Head == node {
		fmt.Println("ОШИБКА: Перед головой ничего нет")
		return
	}
	if l.Head.Next == node {
		l.FRemoveHead()
		return
	}
	cur := l.Head
	for cur != nil && cur.Next != nil && cur.Next.Next != node {
		cur = cur.Next
	}
	if cur != nil {
		cur.Next = cur.Next.Next
	} else {
		fmt.Printf("ОШИБКА: Узла %v нет в списке\n", node.Value)
	}
}

func (l *SinglyList[T]) FRemoveNode(node *SNode[T]) {
	if node == nil {
		fmt.Println("ОШИБКА: Узел не может быть nil")
		return
	}
	if l.Head == node {
		l.FRemoveHead()
		return
	}
	cur := l.Head
	for cur != nil && cur.Next != node {
		cur = cur.Next
	}
	if cur != nil {
		cur.Next = node.Next
	} else {
		fmt.Printf("ОШИБКА: Узла %v нет в списке\n", node.Value)
	}
}

func (l *SinglyList[T]) FRemoveValue(v T) {
	cur := l.Head
	for cur != nil && any(cur.Value) != any(v) {
		cur = cur.Next
	}
	if cur != nil {
		l.FRemoveNode(cur)
	} else {
		fmt.Printf("ОШИБКА: Значения %v нет в списке\n", v)
	}
}

func (l *SinglyList[T]) FSearch(v T) {
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

func (l *SinglyList[T]) FShow() {
	fmt.Print("F[")
	for c := l.Head; c != nil; c = c.Next {
		fmt.Printf("%v ", c.Value)
	}
	fmt.Println("]")
}

func (l *SinglyList[T]) FShowRec(n *SNode[T]) {
	if n == nil {
		return
	}
	fmt.Printf("%v ", n.Value)
	l.FShowRec(n.Next)
}

func (l *SinglyList[T]) Clear() { l.Head = nil }
