// Package bi implements binary tree.
package bi

import (
	"container/list"
	"fmt"
)

// Tree is for binary tree.
type Tree struct {
	Left  *Tree
	Value int64
	Right *Tree
}

// NewTree returns a new tree of input Value.
func NewTree(val int64) *Tree {
	return &Tree{
		Left:  new(Tree),
		Value: val,
		Right: new(Tree),
	}
}

// Insert inserts a new value(node) to the tree.
func (T *Tree) Insert(val int64) *Tree {
	if T == nil {
		return &Tree{nil, val, nil}
		// by doing this
		// the terminal node's left and right
		// are automatically set to nil
	}

	if val < T.Value {
		T.Left = T.Left.Insert(val)
	} else {
		T.Right = T.Right.Insert(val)
	}

	return T
}

// TreePrint prints out the tree.
func (T *Tree) TreePrint() string {
	if T == nil {
		return "()"
	}
	s := ""
	if T.Left != nil {
		s += T.Left.TreePrint() + " "
	}
	s += fmt.Sprintf("%v", T.Value)
	if T.Right != nil {
		s += " " + T.Right.TreePrint()
	}
	return "(" + s + ")"
}

// walkInOrder traverses the tree in the order of
// Left, Root, Right.
func walkInOrder(T *Tree, ch chan int64) {
	// if left sub-tree does exist
	// recursively traverse the left sub-tree
	if T.Left != nil {
		walkInOrder(T.Left, ch)
	}

	// send the value of the present root node
	ch <- T.Value

	if T.Right != nil {
		walkInOrder(T.Right, ch)
	}
}

// WalkInOrder traverses the tree in the order of
// Left, Root, Right.
func WalkInOrder(T *Tree, ch chan int64) {
	walkInOrder(T, ch)
	close(ch)
}

// Same returns true if the two trees are same.
func Same(T1, T2 *Tree) bool {
	ch1, ch2 := make(chan int64), make(chan int64)
	go WalkInOrder(T1, ch1)
	go WalkInOrder(T2, ch2)

	for {
		// if the two trees are the same
		// all values that are sent to channel
		// 				should be equal
		// and the time that the channel gets closed
		// 				should also be equal
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		// TRUE when TWO trees are different
		if v1 != v2 || ok1 != ok2 {
			return false
		}

		// TRUE here only when TWO trees are SAME
		// if one tree gets closed first
		// break and return true
		if !ok1 {
			break
		}
	}
	return true
}

// walkPreOrder traverses the tree in the order of
// Root, Left, Right.
func walkPreOrder(T *Tree, ch chan int64) {
	ch <- T.Value

	if T.Left != nil {
		walkPreOrder(T.Left, ch)
	}

	if T.Right != nil {
		walkPreOrder(T.Right, ch)
	}
}

// WalkPreOrder traverses the tree in the order of
// Root, Left, Right.
func WalkPreOrder(T *Tree, ch chan int64) {
	walkPreOrder(T, ch)
	close(ch)
}

// walkPostOrder traverses the tree in the order of
// Left, Right, Root.
func walkPostOrder(T *Tree, ch chan int64) {
	if T.Left != nil {
		walkPostOrder(T.Left, ch)
	}

	if T.Right != nil {
		walkPostOrder(T.Right, ch)
	}

	ch <- T.Value
}

// WalkPostOrder traverses the tree in the order of
// Left, Right, Root.
func WalkPostOrder(T *Tree, ch chan int64) {
	walkPostOrder(T, ch)
	close(ch)
}

// StringInOrder returns the traversed string
// of the tree in the order of Left, Root, Right.
func StringInOrder(T *Tree, ch chan int64) string {
	go func() {
		defer close(ch)
		walkInOrder(T, ch)
	}()
	s := ""
	for v := range ch {
		s += fmt.Sprintf("%v ", v)
	}
	return s
}

// StringPreOrder returns the traversed string
// of the tree in the order of Root, Left, Right.
func StringPreOrder(T *Tree, ch chan int64) string {
	go func() {
		defer close(ch)
		walkPreOrder(T, ch)
	}()
	s := ""
	for v := range ch {
		s += fmt.Sprintf("%v ", v)
	}
	return s
}

// StringPreOrder returns the traversed string
// of the tree in the order of Left, Right, Root.
func StringPostOrder(T *Tree, ch chan int64) string {
	go func() {
		defer close(ch)
		walkPostOrder(T, ch)
	}()
	s := ""
	for v := range ch {
		s += fmt.Sprintf("%v ", v)
	}
	return s
}

// Copy returns a copy of the tree.
func (T *Tree) Copy() *Tree {
	t := NewTree(T.Value)
	t.Left = T.Left
	t.Right = T.Right
	return t
}

// WalkLevelOrder traverses the tree from the top.
func WalkLevelOrder(T *Tree) *list.List {
	result := list.New()
	queue := list.New()
	result.PushBack(T)
	queue.PushBack(T)
	for queue.Len() > 0 {
		elem := queue.Front()
		tn := elem.Value.(*Tree).Copy()
		queue.Remove(elem)
		if tn.Left != nil {
			result.PushBack(tn.Left)
			queue.PushBack(tn.Left)
		}

		if tn.Right != nil {
			result.PushBack(tn.Right)
			queue.PushBack(tn.Right)
		}
	}
	return result
}

// StringLevelOrder traverses the tree from the top.
func StringLevelOrder(T *Tree) string {
	list := WalkLevelOrder(T)
	fmt.Println(list.Len())
	s := ""
	for elem := list.Front(); elem != nil; elem = elem.Next() {
		s += fmt.Sprintf("%v ", elem.Value.(*Tree).Value)
	}
	return s
}
