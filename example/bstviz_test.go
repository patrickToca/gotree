package example

import (
	"fmt"
	"testing"
	// go test -v github.com/gyuho/gotree/example
	// go test -v /Users/gyuho/go/src/github.com/gyuho/gotree/example/bstviz_test.go
	"github.com/gyuho/gotree/tree/bst"
	// "github.com/gyuho/gotree/tree/bstviz"
)

func Test_bstviz(test *testing.T) {
	t1 := bst.NewTree(5)
	// Perm returns, as a slice of n ints
	// , a pseudo-random permutation of the integers [0,n).
	// func (r *Rand) Perm(n int) []int
	for i := 0; i < 10; i++ {
		t1 = t1.Insert(int64(i))
	}
	fmt.Println(t1.Find(6)) // &{<nil> 6 0x210323140 1}
	// bstviz.Show(t1, "bstviz01.dot")
}
