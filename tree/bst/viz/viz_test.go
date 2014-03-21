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
	tr.Insert(int64(7))
	tr.Insert(int64(8))
	tr.Insert(int64(5))
	tr.Insert(int64(4))
	tr.Insert(int64(2))
	tr.Insert(int64(1))
	tr.Insert(int64(6))
	tr.Insert(int64(3))
	slice := []string{}
	Scan(tr, &slice)
	fmt.Println(slice)
	// Show(tr, "tree2.dot")
}
