// Package example is to show code usage.
// It is to be run with the command go test -v.
package example

import (
	"fmt"
	"testing"
	// go test -v github.com/gyuho/gotree/example
	// go test -v /Users/gyuho/go/src/github.com/gyuho/gotree/example/bi_test.go
	"github.com/gyuho/gotree/bi"
)

func Test_bi(test *testing.T) {
	t1 := bi.NewTree(5)
	// Perm returns, as a slice of n ints
	// , a pseudo-random permutation of the integers [0,n).
	// func (r *Rand) Perm(n int) []int
	for i := 0; i < 10; i++ {
		t1 = t1.Insert(int64(i))
	}

	fmt.Println(t1.TreePrint())
	// ((0 (0 (1 (2 (3 (4)))))) 5 (0 (5 (6 (7 (8 (9)))))))
	// 0 is nil

	ch1 := make(chan int64)
	fmt.Println(bi.StringInOrder(t1, ch1))
	// 0 0 1 2 3 4 5 0 5 6 7 8 9

	println()
	ch2 := make(chan int64)
	fmt.Println(bi.StringPreOrder(t1, ch2))
	// 5 0 0 1 2 3 4 0 5 6 7 8 9

	println()
	ch3 := make(chan int64)
	fmt.Println(bi.StringPostOrder(t1, ch3))
	// 4 3 2 1 0 0 9 8 7 6 5 0 5

	println()
	fmt.Println(bi.StringLevelOrder(t1))
	// 5 0 0 0 5 1 6 2 7 3 8 4 9
}
