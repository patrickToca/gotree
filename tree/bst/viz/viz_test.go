package viz

import (
	"fmt"
	"testing"

	"github.com/gyuho/gotree/tree/bst"
)

// Show works as expected,
// but it is commented out
// just for the sake of Travis.org testing
func Test_Show1(test *testing.T) {
	tr := bst.NewTree(5)
	for i := 0; i < 10; i++ {
		if i != 5 {
			tr = tr.Insert(int64(i))
		}
	}
	slice := []string{}
	Scan(tr, &slice)
	fmt.Println(slice)
	// Show(tr, "tree1.dot")
}

func Test_Show2(test *testing.T) {
	tr := bst.NewTree(5)
	tr.Inserts(7, 8, 4, 2, 1, 6, 3)
	slice := []string{}
	Scan(tr, &slice)
	fmt.Println(slice)
	// Show(tr, "tree2.dot")
}

func Test_Show3(test *testing.T) {
	tr := bst.NewTree(5)
	tr.Inserts(7, 8, 4, 2, 1, 6, 3)
	tr.Delete(int64(6))
	// Show(tr, "tree3.dot")
}

func Test_Show4(test *testing.T) {
	tr := bst.NewTree(5)
	tr.Inserts(7, 8, 4, 2, 1, 3)
	tr.Delete(int64(7))
	// Show(tr, "tree4.dot")
}

func Test_Show5(test *testing.T) {
	tr := bst.NewTree(5)
	tr.Inserts(7, 8, 3, 4, 2, 1, 6)
	tr = tr.Delete(int64(5))
	// Show(tr, "tree5.dot")
}
