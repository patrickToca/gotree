package avlviz

import (
	"testing"

	"github.com/gyuho/gotree/tree/avl"
)

// Show works as expected,
// but it is commented out
// just for the sake of Travis.org testing
func Test_Show1(test *testing.T) {
	tr := avl.NewTree(4)
	tr.Inserts(3, 2)
	// Show(tr, "avl01.dot")
}

func Test_Show2(test *testing.T) {
	tr := avl.NewTree(5)
	tr.Inserts(7, 8, 4, 2, 1, 3)
	// Show(tr, "avl02.dot")
}

func Test_Show3(test *testing.T) {
	tr := avl.NewTree(5)
	tr.Inserts(7, 8, 4, 2, 1, 6, 3, 9, 10)
	// Show(tr, "avl03.dot")
}

func Test_Show4(test *testing.T) {
	tr := avl.NewTree(4)
	tr.BalancingInsert(7)
	tr.BalancingInsert(5)
	// Show(tr, "avl04.dot")
}

func Test_Show5(test *testing.T) {
	tr := avl.NewTree(4)
	tr.BalancingInsert(3)
	tr.BalancingInsert(2)
	tr.RebalanceLL(2)
	// Show(tr, "avl05.dot")
}
