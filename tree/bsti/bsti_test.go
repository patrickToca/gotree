package bsti

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

func Test_Find(test *testing.T) {
	t1 := NewTree(5)
	for i := 0; i < 10; i++ {
		t1 = t1.Insert(int64(i))
	}
	tr := t1.Find(7)
	if tr.Value != 7 {
		test.Errorf("Should be true but %+v", tr)
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

	ch1 := make(chan int64)
	s1 := StringInOrder(t1, ch1)
	sc1 := "0 0 1 2 3 4 5 0 5 6 7 8 9 "

	if s1 != sc1 {
		test.Errorf("Should be same but %v, %v", s1, sc1)
	}
}

func Test_StringPreOrder(test *testing.T) {
	t1 := NewTree(5)
	for i := 0; i < 10; i++ {
		t1 = t1.Insert(int64(i))
	}

	ch1 := make(chan int64)
	s1 := StringPreOrder(t1, ch1)
	sc1 := "5 0 0 1 2 3 4 0 5 6 7 8 9 "

	if s1 != sc1 {
		test.Errorf("Should be same but %v, %v", s1, sc1)
	}
}

func Test_StringPostOrder(test *testing.T) {
	t1 := NewTree(5)
	for i := 0; i < 10; i++ {
		t1 = t1.Insert(int64(i))
	}

	ch1 := make(chan int64)
	s1 := StringPostOrder(t1, ch1)
	sc1 := "4 3 2 1 0 0 9 8 7 6 5 0 5 "

	if s1 != sc1 {
		test.Errorf("Should be same but %v, %v", s1, sc1)
	}
}

func Test_StringLevelOrder(test *testing.T) {
	t1 := NewTree(5)
	for i := 0; i < 10; i++ {
		t1 = t1.Insert(int64(i))
	}

	s1 := StringLevelOrder(t1)
	sc1 := "5 0 0 0 5 1 6 2 7 3 8 4 9 "

	if s1 != sc1 {
		test.Errorf("Should be same but %v, %v", s1, sc1)
	}
}
