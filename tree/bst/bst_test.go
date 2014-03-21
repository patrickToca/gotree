package bst

import (
	"fmt"
	"testing"
)

func Test_NewTree(test *testing.T) {
	t1 := NewTree(10)
	v1 := fmt.Sprintf("%v", t1.Value)
	vc1 := fmt.Sprintf("%v", 10)
	if v1 != vc1 {
		test.Errorf("Should be same but %v, %v", v1, vc1)
	}
}

func Test_Insert(test *testing.T) {
	t1 := NewTree(5)
	for i := 0; i < 10; i++ {
		t1 = t1.Insert(int64(i))
	}
	if t1.Size != 11 {
		test.Errorf("Should be size of 11 but %+v", t1.Size)
	}
}

func Test_Find(test *testing.T) {
	t1 := NewTree(5)
	for i := 0; i < 10; i++ {
		t1 = t1.Insert(int64(i))
	}
	for i := 0; i < 10; i++ {
		if t1.Find(int64(i)) == nil {
			test.Errorf("Should exist but %+v", t1.Find(int64(i)))
		}
	}
}

func Test_Same(test *testing.T) {
	t1 := NewTree(5)
	for i := 0; i < 10; i++ {
		t1 = t1.Insert(int64(i))
	}

	t2 := NewTree(5)
	for i := 0; i < 10; i++ {
		t2 = t2.Insert(int64(i))
	}

	if !Same(t1, t2) {
		test.Errorf("Should be same but %v, %v", t1, t2)
	}
}

func Test_StringInOrder(test *testing.T) {
	t1 := NewTree(5)
	for i := 0; i < 10; i++ {
		t1 = t1.Insert(int64(i))
	}

	sc1 := "0 1 2 3 4 5 6 7 8 9 "
	ch1 := make(chan int64)
	s1 := StringInOrder(t1, ch1)

	if s1 != sc1 {
		test.Errorf("Should be\n%v\n\nbut\n%v", sc1, s1)
	}
}

func Test_StringPreOrder(test *testing.T) {
	t1 := NewTree(5)
	for i := 0; i < 10; i++ {
		t1 = t1.Insert(int64(i))
	}

	sc1 := "5 0 1 2 3 4 6 7 8 9 "
	ch1 := make(chan int64)
	s1 := StringPreOrder(t1, ch1)

	if s1 != sc1 {
		test.Errorf("Should be\n%v\n\nbut\n%v", sc1, s1)
	}
}

func Test_StringPostOrder(test *testing.T) {
	t1 := NewTree(5)
	for i := 0; i < 10; i++ {
		t1 = t1.Insert(int64(i))
	}

	sc1 := "4 3 2 1 0 9 8 7 6 5 "
	ch1 := make(chan int64)
	s1 := StringPostOrder(t1, ch1)

	if s1 != sc1 {
		test.Errorf("Should be\n%v\n\nbut\n%v", sc1, s1)
	}
}

func Test_StringLevelOrder(test *testing.T) {
	t1 := NewTree(5)
	for i := 0; i < 10; i++ {
		t1 = t1.Insert(int64(i))
	}

	sc1 := "5 0 6 1 7 2 8 3 9 4 "
	s1 := StringLevelOrder(t1)

	if s1 != sc1 {
		test.Errorf("Should be\n%v\n\nbut\n%v", sc1, s1)
	}
}
