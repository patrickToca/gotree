package avlviz

import (
	"fmt"
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
	tr.BalanceInsert(7)
	tr.BalanceInsert(5)
	// Show(tr, "avl04.dot")
}

func Test_Show5(test *testing.T) {
	tr := avl.NewTree(4)
	tr.BalanceInsert(3)
	tr.BalanceInsert(2)
	tr.BalanceLL(2)
	// Show(tr, "avl05.dot")
}

func Test_Show6(test *testing.T) {
	tr := avl.NewTree(4)
	tr.BalanceInsert(2)
	tr.BalanceInsert(3)
	// Show(tr, "avl06.dot")

	tr.BalanceLR(3)
	// Show(tr, "avl06.dot")
}

func Test_Show7(test *testing.T) {
	tr := avl.NewTree(4)
	tr.BalanceInsert(6)
	tr.BalanceInsert(5)
	// Show(tr, "avl07.dot")

	tr.BalanceRL(5)
	// Show(tr, "avl07.dot")
}

func Test_Show8(test *testing.T) {
	tr := avl.NewTree(10)
	tr.BalanceInserts(13, 17, 5, 4, 7)
	// Show(tr, "avl08.dot")
}

func Test_Show9(test *testing.T) {
	tr := avl.NewTree(10)
	tr.BalanceInserts(13, 17)
	fmt.Println(tr.Right.Value)
	// Show(tr, "avl09.dot")
}

func Test_Show10(test *testing.T) {
	tr := avl.NewTree(13)
	tr.Inserts(5, 17, 4, 10)
	// Show(tr, "avl10.dot")
}
