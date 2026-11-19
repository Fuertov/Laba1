package main

import "fmt"

type TNode[T any] struct {
	Value T
	Left  *TNode[T]
	Right *TNode[T]
}

type CBTree[T any] struct {
	Root *TNode[T]
	Size int
}

func (t *CBTree[T]) TAdd(v T) {
	t.Root = t.placeAt(t.Root, t.Size, 0, v)
	t.Size++
}

func (t *CBTree[T]) placeAt(r *TNode[T], target, idx int, v T) *TNode[T] {
	if r == nil {
		if idx == target {
			return &TNode[T]{Value: v}
		}
		return nil
	}
	if 2*idx+1 <= target {
		r.Left = t.placeAt(r.Left, target, 2*idx+1, v)
	}
	if 2*idx+2 <= target {
		r.Right = t.placeAt(r.Right, target, 2*idx+2, v)
	}
	return r
}

func (t *CBTree[T]) TSearch(v T) {
	depth := 0
	if t.searchRec(t.Root, v, &depth) {
		fmt.Printf("Элемент был найден на глубине %d\n", depth)
	} else {
		fmt.Println("Элемент не был найден в дереве")
	}
}

func (t *CBTree[T]) searchRec(r *TNode[T], v T, depth *int) bool {
	if r == nil {
		return false
	}
	if any(r.Value) == any(v) {
		return true
	}
	if t.searchRec(r.Left, v, depth) {
		*depth++
		return true
	}
	if t.searchRec(r.Right, v, depth) {
		*depth++
		return true
	}
	return false
}

func (t *CBTree[T]) TShow() {
	fmt.Print("T[")
	t.inOrder(t.Root)
	fmt.Println("]")
}

func (t *CBTree[T]) inOrder(r *TNode[T]) {
	if r == nil {
		return
	}
	t.inOrder(r.Left)
	fmt.Printf("%v ", r.Value)
	t.inOrder(r.Right)
}

func (t *CBTree[T]) TShowTree() { t.drawTree(t.Root, 0) }

func (t *CBTree[T]) drawTree(r *TNode[T], level int) {
	if r == nil {
		return
	}
	t.drawTree(r.Right, level+1)
	for i := 0; i < level; i++ {
		fmt.Print("\t")
	}
	fmt.Println(r.Value)
	t.drawTree(r.Left, level+1)
}

func (t *CBTree[T]) TCheckComplete() bool {
	return t.checkCompleteRec(t.Root, 0, t.Size)
}

func (t *CBTree[T]) checkCompleteRec(r *TNode[T], i, n int) bool {
	if r == nil {
		return true
	}
	if i >= n {
		return false
	}
	return t.checkCompleteRec(r.Left, 2*i+1, n) &&
		t.checkCompleteRec(r.Right, 2*i+2, n)
}

func (t *CBTree[T]) TCheckFull() bool { return t.checkFullRec(t.Root) }

func (t *CBTree[T]) checkFullRec(r *TNode[T]) bool {
	if r == nil {
		return true
	}
	if r.Left == nil && r.Right == nil {
		return true
	}
	if r.Left != nil && r.Right != nil {
		return t.checkFullRec(r.Left) && t.checkFullRec(r.Right)
	}
	return false
}

func (t *CBTree[T]) Clear() {
	t.Root = nil
	t.Size = 0
}
