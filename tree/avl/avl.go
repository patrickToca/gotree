// Pacakge avl implements an AVL tree.
package avl

import "math"

// Tree is for binary tree.
type Tree struct {
	Left  *Tree
	Value int64
	Right *Tree
	Size  int64
}

// NewTree returns a new tree of input Value.
func NewTree(val int64) *Tree {
	return &Tree{
		Left:  nil, // (X) Left: new(Tree),
		Value: val,
		Right: nil, // (X) Right: new(Tree),
		Size:  1,
	}
}

// BalancingInserts implements Insert with a variadic function.
func (T *Tree) BalancingInserts(values ...int64) *Tree {
	for _, v := range values {
		T.Insert(v)
	}
	return T
}

// Inserts implements Insert with a variadic function.
func (T *Tree) Inserts(values ...int64) *Tree {
	for _, v := range values {
		T.Insert(v)
	}
	return T
}

// Insert inserts a new value(node) to the tree.
func (T *Tree) Insert(val int64) *Tree {
	if T != nil && T.Value != val {
		T.Size += 1
	}
	// To end recursion
	// set terminal node's left and right to nil
	if T == nil {
		return &Tree{nil, val, nil, int64(1)}
	}
	if val < T.Value {
		// Insert into the left tree
		T.Left = T.Left.Insert(val)
		// We don't need to do this
		// (X) to increase the size of the left sub-tree
		// (X) T.Left.Size += 1
	} else if val > T.Value {
		// Insert into the right tree
		T.Right = T.Right.Insert(val)
		// (X) to increase the size of the right sub-tree
		// (X) T.Right.Size += 1
	}
	return T
}

// Find does Binary Search to find the value
// and returns true if the value exists in the Tree.
func (T *Tree) Find(val int64) *Tree {
	if T == nil {
		return &Tree{nil, val, nil, int64(1)}
	}
	// To end recursion
	// set terminal node's left and right to nil
	if T.Value == val {
		return T
	}
	if val < T.Value {
		// Not working if we only have
		// T.Left.Find(val)
		return T.Left.Find(val)
	} else {
		return T.Right.Find(val)
	}
	return T
}

// GetSize returns the size of the node with the input value
// which is the number of the children node + 1.
func (T *Tree) GetSize(val int64) int64 {
	if T.Value == val {
		return T.Size
	}
	node := T.Find(val)
	return node.Size
}

// GetHeight returns the height of the node with the input value.
func (T *Tree) GetHeight(val int64) int64 {
	// Height h = ⌊log_2 n⌋
	// n = 15 (15 nodes), then the height of the root node
	// is = ⌊log_2 15⌋ = 3

	// in order to truncate
	// need to make a copy by passing as a parameter
	float64ToInt64 := func(num float64) int64 {
		return int64(num)
	}
	tree := T.Find(val)
	return float64ToInt64(math.Floor(math.Log2(float64(tree.GetSize(val)))))
}

// GetHeightLeft returns the height of the left sub-tree
// of the input value(node).
func (T *Tree) GetHeightLeft(val int64) int64 {
	tree := T.Find(val)
	if tree.Left == nil {
		return 0
	}
	h := T.GetHeight(tree.Left.Value)
	return h + 1
}

// GetHeightRight returns the height of the right sub-tree
// of the input value(node).
func (T *Tree) GetHeightRight(val int64) int64 {
	tree := T.Find(val)
	if tree.Right == nil {
		return 0
	}
	// (X) T.GetHeight(tree.Value)
	h := T.GetHeight(tree.Right.Value)
	return h + 1
}

// Height returns the difference between
// GetHeightLeft and GetHeightRight.
func (T *Tree) Height(val int64) int64 {
	return T.GetHeightLeft(val) - T.GetHeightRight(val)
}

// IsBalanced returns true if the Height of the Tree
// with the input value is balanced.
func (T *Tree) IsBalanced(val int64) bool {
	return -1 <= T.Height(val) && T.Height(val) <= 1
}

// Parent returns the parental Tree(node) of input value.
func (T *Tree) Parent(val int64) *Tree {
	// if the input value is nil
	if val == T.Value {
		return nil
	}
	if T == nil {
		return &Tree{nil, val, nil, int64(1)}
	}
	// we need to check if T.Left is nil or not
	// otherwise, it panics with the message:
	// panic: runtime error: invalid memory address
	// or nil pointer dereference [recovered]
	if T.Left != nil {
		if T.Left.Value == val {
			return T
		}
	}
	if T.Right != nil {
		if T.Right.Value == val {
			return T
		}
	}
	if val < T.Value {
		return T.Left.Parent(val)
	} else {
		return T.Right.Parent(val)
	}
	return T
}

// IsRoot returns true if the Node(tree) is a root of the tree.
func (T *Tree) IsRoot(val int64) bool {
	if T.Parent(val) == nil {
		return true
	} else {
		return false
	}
}

// IsLeaf returns true if the Node(tree) is a leaf.
func (T *Tree) IsLeaf(val int64) bool {
	nd := T.Find(val)
	if nd == nil {
		return false
	}
	if nd.Left == nil && nd.Right == nil {
		return true
	} else {
		return false
	}
}
