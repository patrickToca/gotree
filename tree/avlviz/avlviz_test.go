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
	Show(tr, "avl01.dot")
}
