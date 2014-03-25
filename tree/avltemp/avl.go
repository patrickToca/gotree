// Pacakge avl implements an AVL tree.
package avltemp

import (
	"math"
	"sort"
)

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

// Inserts implements Insert with a variadic function.
func (T *Tree) Inserts(values ...int64) *Tree {
	for _, v := range values {
		T.Insert(v)
	}
	return T
}

// Insert inserts a new value(node) to the tree.
func (T *Tree) Insert(val int64) *Tree {
	if T == nil {
		return &Tree{nil, val, nil, int64(1)}
	}
	if T != nil && T.Value != val {
		T.Size += 1
	}
	if val < T.Value {
		T.Left = T.Left.Insert(val)
	} else if val > T.Value {
		T.Right = T.Right.Insert(val)
	}
	return T
}

// Find does Binary Search to find the value
// and returns Tree with the value as a root node.
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

// Parent returns the parental Tree(node) of input value.
func (T *Tree) Parent(val int64) *Tree {
	// if the input value is root
	if val == T.Value {
		return T
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
	if T.Parent(val).Value == val {
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

// Detect returns the Height and detects the balancing status.
func (T *Tree) Detect(val int64) (int64, string) {
	num := T.Height(val)
	Tr := T.Find(val)
	result := ""
	if T.IsBalanced(val) {
		result = "Balanced"
	} else {
		switch T.Height(Tr.Value) {
		case 2: // LL or LR
			if Tr.Left.Left != nil && Tr.Left.Right == nil {
				result = "LL"
			}
			if Tr.Left.Left == nil && Tr.Left.Right != nil {
				result = "LR"
			}
		case -2: // RR or RL
			if Tr.Right.Right != nil && Tr.Right.Left == nil {
				result = "RR"
			}
			if Tr.Right.Right == nil && Tr.Right.Left != nil {
				result = "RL"
			}
		}
	}
	return num, result
}

// IsBalanced returns true if the Height of the Tree
// with the input value is balanced.
func (T *Tree) IsBalanced(val int64) bool {
	Tree := T.Find(val)
	return -1 <= Tree.Height(val) && Tree.Height(val) <= 1
}

// BalanceInsert inserts one value to a Tree
// and tells if the Tree is LL, RR, LR, RL.
func (T *Tree) BalanceInsert(val int64) (*Tree, string) {
	T.Insert(val)
	pt := T.Parent(val)
	Parent := T.Parent(pt.Value)
	if T.IsBalanced(Parent.Value) {
		return T, "Balanced"
	} else {
		switch T.Height(Parent.Value) {
		case 2: // LL or LR
			if Parent.Left.Left != nil && Parent.Left.Right == nil {
				return T, "LL"
			}
			if Parent.Left.Left == nil && Parent.Left.Right != nil {
				return T, "LR"
			}
		case -2: // RR or RL
			if Parent.Right.Right != nil && Parent.Right.Left == nil {
				return T, "RR"
			}
			if Parent.Right.Right == nil && Parent.Right.Left != nil {
				return T, "RL"
			}
		}
	}
	return T, "None"
}

type Int64Slice []int64

func (p Int64Slice) Len() int {
	return len(p)
}
func (p Int64Slice) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p Int64Slice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// int64Sort sorts 3 integers and return the slice in order.
func int64Sort(v1, v2, v3 int64) []int64 {
	slice := []int64{v1, v2, v3}
	sort.Sort(Int64Slice(slice))
	return slice
}

// SetValue updates the Value of the Tree.
func (T *Tree) SetValue(val int64) {
	T.Value = val
}

// BalanceLL balances a LL tree with the val of a leaf node.
// Not only changes the children, it also changes the order of the root.
func (T *Tree) BalanceLL(val int64) *Tree {
	//
	//     Parent
	//      /
	//     pt
	//    /
	//  node
	//
	// to
	//
	//     Parent
	//      /  \
	//    pt   node
	//
	node := T.Find(val)
	pt := T.Parent(val)
	Parent := T.Parent(pt.Value)
	// Now balanced!
	Parent.Right = node

	// Update values (Balance)
	slice := int64Sort(node.Value, pt.Value, Parent.Value)

	// 2nd Biggest should be Parent(Root)
	Parent.SetValue(slice[1])

	// Biggest should be Right
	Parent.Right.SetValue(slice[2])

	// Smallest Should be Left
	Parent.Left.SetValue(slice[0])

	// Deletion
	// 1. Delete as a child
	pt.Size -= 1
	pt.Left = nil

	// 2. Delete the moved node itself
	// (X) T.Size -= 1
	// we just move the node not deleting from the whole tree
	node = nil

	return T
}

// BalanceLR balances a LR tree with the val of a leaf node.
// Not only changes the children, it also changes the order of the root.
func (T *Tree) BalanceLR(val int64) *Tree {
	//
	//     Parent
	//      /
	//     pt
	//       \
	//       node
	//
	// to
	//
	//     Parent
	//      /  \
	//     pt  node
	//
	node := T.Find(val)
	pt := T.Parent(val)
	Parent := T.Parent(pt.Value)
	// Now balanced!
	Parent.Right = node

	// Update values (Balance)
	slice := int64Sort(node.Value, pt.Value, Parent.Value)

	// 2nd Biggest should be Parent(Root)
	Parent.SetValue(slice[1])

	// Biggest should be Right
	Parent.Right.SetValue(slice[2])

	// Smallest Should be Left
	Parent.Left.SetValue(slice[0])

	// Deletion
	// 1. Delete as a child
	pt.Size -= 1
	pt.Right = nil

	// 2. Delete the moved node itself
	node = nil

	return T
}

// BalanceRR balances a RR tree with the val of a leaf node.
// Not only changes the children, it also changes the order of the root.
func (T *Tree) BalanceRR(val int64) *Tree {
	//
	//     Parent
	//        \
	//         pt
	//          \
	//          node
	//
	// to
	//
	//     Parent
	//      /  \
	//   node   pt
	//
	node := T.Find(val)
	pt := T.Parent(val)
	Parent := T.Parent(pt.Value)
	// Now balanced!
	Parent.Left = node

	// Update values (Balance)
	slice := int64Sort(node.Value, pt.Value, Parent.Value)

	// 2nd Biggest should be Parent(Root)
	Parent.SetValue(slice[1])

	// Biggest should be Right
	Parent.Right.SetValue(slice[2])

	// Smallest Should be Left
	Parent.Left.SetValue(slice[0])

	// Deletion
	// 1. Delete as a child
	pt.Size -= 1
	pt.Right = nil

	// 2. Delete the moved node itself
	// (X) T.Size -= 1
	// we just move the node not deleting from the whole tree
	node = nil

	return T
}

// BalanceRL balances a RL tree with the val of a leaf node.
// Not only changes the children, it also changes the order of the root.
func (T *Tree) BalanceRL(val int64) *Tree {
	//
	//     Parent
	//         \
	//          pt
	//          /
	//        node
	//
	// to
	//
	//     Parent
	//      /  \
	//   node   pt
	//
	node := T.Find(val)
	pt := T.Parent(val)
	Parent := T.Parent(pt.Value)
	// Now balanced!
	Parent.Left = node

	// Update values (Balance)
	slice := int64Sort(node.Value, pt.Value, Parent.Value)

	// 2nd Biggest should be Parent(Root)
	Parent.SetValue(slice[1])

	// Biggest should be Right
	Parent.Right.SetValue(slice[2])

	// Smallest Should be Left
	Parent.Left.SetValue(slice[0])

	// Deletion
	// 1. Delete as a child
	pt.Size -= 1
	pt.Left = nil

	// 2. Delete the moved node itself
	node = nil

	return T
}

// RotateLeft rotates the Tree in the counter clockwise direction.
func (T *Tree) RotateLeft(val int64) *Tree {
	//
	//      P
	//     /  \
	//    A    Q
	//        / \
	//       B   C
	//
	// to
	//
	//     Parent
	//      /  \
	//   node   pt
	//
	node := T.Find(val)
	pt := T.Parent(val)
	Parent := T.Parent(pt.Value)
	// Now balanced!
	Parent.Left = node

	// Update values (Balance)
	slice := int64Sort(node.Value, pt.Value, Parent.Value)

	// 2nd Biggest should be Parent(Root)
	Parent.SetValue(slice[1])

	// Biggest should be Right
	Parent.Right.SetValue(slice[2])

	// Smallest Should be Left
	Parent.Left.SetValue(slice[0])

	// Deletion
	// 1. Delete as a child
	pt.Size -= 1
	pt.Left = nil

	// 2. Delete the moved node itself
	node = nil

	return T
}

// BalanceInserts implements Insert with a variadic function.
func (T *Tree) BalanceInserts(values ...int64) *Tree {
	for _, val := range values {
		_, rs := T.BalanceInsert(val)

		switch rs {

		case "LL":
			T.BalanceLL(val)

		case "LR":
			T.BalanceLR(val)

		case "RR":
			T.BalanceRR(val)

		case "RL":
			T.BalanceRL(val)

		}
	}
	return T
}
