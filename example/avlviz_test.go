package example

import (
	"testing"
	// go test -v github.com/gyuho/gotree/example
	// go test -v /Users/gyuho/go/src/github.com/gyuho/gotree/example/avlviz_test.go
	"github.com/gyuho/gotree/tree/avl"
	// "github.com/gyuho/gotree/tree/avlviz"
)

func Test_avlviz(test *testing.T) {
	tr := avl.NewTree(4)
	tr.BalanceInsert(6)
	tr.BalanceInsert(5)
	// avlviz.Show(tr, "avl-before.dot")

	tr.BalanceRL(5)
	// avlviz.Show(tr, "avl-after.dot")
}
