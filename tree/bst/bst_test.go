package bst

import (
	"fmt"
	"testing"
)

func Test_NewTree(test *testing.T) {
	tr := NewTree(10)
	v1 := fmt.Sprintf("%v", tr.Value)
	vc1 := fmt.Sprintf("%v", 10)
	if v1 != vc1 {
		test.Errorf("Should be same but %v, %v", v1, vc1)
	}
}

func Test_Insert(test *testing.T) {
	tr := NewTree(5)
	for i := 0; i < 10; i++ {
		tr = tr.Insert(int64(i))
	}
	if tr.Size != 11 {
		test.Errorf("Should be size of 11 but %+v", tr.Size)
	}
}

func Test_Find(test *testing.T) {
	tr := NewTree(5)
	for i := 0; i < 10; i++ {
		tr = tr.Insert(int64(i))
	}
	for i := 0; i < 10; i++ {
		if tr.Find(int64(i)) == nil {
			test.Errorf("Should exist but %+v", tr.Find(int64(i)))
		}
	}

	tt := NewTree(5)
	tt.Insert(int64(7))
	tt.Insert(int64(8))
	tt.Insert(int64(5))
	tt.Insert(int64(4))
	tt.Insert(int64(2))
	tt.Insert(int64(1))
	tt.Insert(int64(6))
	tt.Insert(int64(3))

	fr := tt.Find(4)
	if fr.Left.Value != int64(2) {
		test.Errorf("Should exist but %+v", tt.Find(int64(4)))
	}
}

func Test_Parent(test *testing.T) {
	tr := NewTree(5)
	tr.Insert(int64(7))
	tr.Insert(int64(8))
	tr.Insert(int64(5))
	tr.Insert(int64(4))
	tr.Insert(int64(2))
	tr.Insert(int64(1))
	tr.Insert(int64(6))
	tr.Insert(int64(3))
	if tr.Parent(int64(6)).Value != 7 {
		test.Errorf("Parent should be 7 but\n%v", tr.Parent(int64(6)).Value)
	}
	if tr.Parent(int64(8)).Value != 7 {
		test.Errorf("Parent should be 7 but\n%v", tr.Parent(int64(8)).Value)
	}
	if tr.Parent(int64(7)).Value != 5 {
		test.Errorf("Parent should be 5 but\n%v", tr.Parent(int64(7)).Value)
	}
	if tr.Parent(int64(1)).Value != 2 {
		test.Errorf("Parent should be 2 but\n%v", tr.Parent(int64(1)).Value)
	}
	if tr.Parent(int64(5)) != nil {
		test.Errorf("Parent should be nil but\n%v", tr.Parent(int64(5)))
	}
}

func Test_FindMinMax(test *testing.T) {
	tr := NewTree(5)
	tr.Insert(int64(7))
	tr.Insert(int64(8))
	tr.Insert(int64(5))
	tr.Insert(int64(4))
	tr.Insert(int64(2))
	tr.Insert(int64(1))
	tr.Insert(int64(6))
	tr.Insert(int64(3))
	if tr.FindMin() != 1 {
		test.Errorf("Should be 1 but\n%v", tr.FindMin())
	}
	if tr.FindMax() != 8 {
		test.Errorf("Should be 8 but\n%v", tr.FindMax())
	}
}

func Test_IsLeaf(test *testing.T) {
	tr := NewTree(5)
	tr.Insert(int64(7))
	tr.Insert(int64(8))
	tr.Insert(int64(5))
	tr.Insert(int64(4))
	tr.Insert(int64(2))
	tr.Insert(int64(1))
	tr.Insert(int64(6))
	tr.Insert(int64(3))
	if tr.IsLeaf(int64(5)) {
		test.Errorf("IsLeaf should return false but\n%v", tr.IsLeaf(int64(5)))
	}
	if !tr.IsLeaf(int64(1)) {
		test.Errorf("IsLeaf should return true but\n%v", tr.IsLeaf(int64(1)))
	}
	if !tr.IsLeaf(int64(3)) {
		test.Errorf("IsLeaf should return true but\n%v", tr.IsLeaf(int64(3)))
	}
	if !tr.IsLeaf(int64(6)) {
		test.Errorf("IsLeaf should return true but\n%v", tr.IsLeaf(int64(6)))
	}
	if !tr.IsLeaf(int64(8)) {
		test.Errorf("IsLeaf should return true but\n%v", tr.IsLeaf(int64(8)))
	}
}

func Test_Same(test *testing.T) {
	tr := NewTree(5)
	for i := 0; i < 10; i++ {
		tr = tr.Insert(int64(i))
	}

	t2 := NewTree(5)
	for i := 0; i < 10; i++ {
		t2 = t2.Insert(int64(i))
	}

	if !Same(tr, t2) {
		test.Errorf("Should be same but %v, %v", tr, t2)
	}
}

func Test_StringInOrder(test *testing.T) {
	tr := NewTree(5)
	for i := 0; i < 10; i++ {
		tr = tr.Insert(int64(i))
	}

	sc1 := "0 1 2 3 4 5 6 7 8 9 "
	ch1 := make(chan int64)
	s1 := StringInOrder(tr, ch1)

	if s1 != sc1 {
		test.Errorf("Should be\n%v\n\nbut\n%v", sc1, s1)
	}
}

func Test_StringPreOrder(test *testing.T) {
	tr := NewTree(5)
	for i := 0; i < 10; i++ {
		tr = tr.Insert(int64(i))
	}

	sc1 := "5 0 1 2 3 4 6 7 8 9 "
	ch1 := make(chan int64)
	s1 := StringPreOrder(tr, ch1)

	if s1 != sc1 {
		test.Errorf("Should be\n%v\n\nbut\n%v", sc1, s1)
	}
}

func Test_StringPostOrder(test *testing.T) {
	tr := NewTree(5)
	for i := 0; i < 10; i++ {
		tr = tr.Insert(int64(i))
	}

	sc1 := "4 3 2 1 0 9 8 7 6 5 "
	ch1 := make(chan int64)
	s1 := StringPostOrder(tr, ch1)

	if s1 != sc1 {
		test.Errorf("Should be\n%v\n\nbut\n%v", sc1, s1)
	}
}

func Test_StringLevelOrder(test *testing.T) {
	tr := NewTree(5)
	for i := 0; i < 10; i++ {
		tr = tr.Insert(int64(i))
	}

	sc1 := "5 0 6 1 7 2 8 3 9 4 "
	s1 := StringLevelOrder(tr)

	if s1 != sc1 {
		test.Errorf("Should be\n%v\n\nbut\n%v", sc1, s1)
	}
}
