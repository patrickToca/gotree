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
