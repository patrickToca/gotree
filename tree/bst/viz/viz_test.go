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
	tr.Inserts(7, 8, 5, 4, 2, 1, 6, 3)
	slice := []string{}
	Scan(tr, &slice)
	fmt.Println(slice)
	// Show(tr, "tree2.dot")
}

func Test_Show3(test *testing.T) {
	tr := bst.NewTree(5)
	tr.Inserts(7, 8, 5, 4, 2, 1, 6, 3)
	tr.Delete(int64(6))
	// Show(tr, "tree3.dot")
}