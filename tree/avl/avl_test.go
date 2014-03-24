package avl

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

func Test_Inserts(test *testing.T) {
	tr := NewTree(5)
	tr.Inserts(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	if tr.Size != 10 {
		test.Errorf("Should be size of 10 but %+v", tr.Size)
	}
}

func Test_Insert(test *testing.T) {
	tr := NewTree(5)
	for i := 0; i < 10; i++ {
		tr.Insert(int64(i))
		// tr = tr.Insert(int64(i))
	}
	if tr.GetSize(5) != 10 {
		test.Errorf("Should be size of 10 but %+v", tr.GetSize(5))
	}
	if tr.Size != 10 {
		test.Errorf("Should be size of 10 but %+v", tr.Size)
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
	tt.Inserts(7, 8, 5, 4, 2, 1, 6, 3)
	fr := tt.Find(4)
	if fr.Left.Value != int64(2) {
		test.Errorf("Should exist but %+v", tt.Find(int64(4)))
	}
}

func Test_Parent(test *testing.T) {
	tr := NewTree(5)
	tr.Inserts(7, 8, 5, 4, 2, 1, 6, 3)
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
	if tr.Parent(int64(5)).Value != 5 {
		test.Errorf("Parent should be 5 but\n%v", tr.Parent(int64(5)).Value)
	}
}

func Test_IsLeaf(test *testing.T) {
	tr := NewTree(5)
	tr.Inserts(7, 8, 5, 4, 2, 1, 6, 3)
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

func Test_IsRoot(test *testing.T) {
	tr := NewTree(5)
	tr.Inserts(7, 8, 5, 4, 2, 1, 6, 3)
	if !tr.IsRoot(int64(5)) {
		test.Errorf("IsRoot should return true but\n%v", tr.IsRoot(int64(5)))
	}
	if tr.IsRoot(int64(1)) {
		test.Errorf("IsRoot should return false but\n%v", tr.IsRoot(int64(1)))
	}
	if tr.IsRoot(int64(3)) {
		test.Errorf("IsRoot should return false but\n%v", tr.IsRoot(int64(3)))
	}
	if tr.IsRoot(int64(6)) {
		test.Errorf("IsRoot should return false but\n%v", tr.IsRoot(int64(6)))
	}
	if tr.IsRoot(int64(8)) {
		test.Errorf("IsRoot should return false but\n%v", tr.IsRoot(int64(8)))
	}
}

func Test_GetSize(test *testing.T) {
	tr := NewTree(5)
	tr.Inserts(7, 8, 4, 2, 1, 6, 3)
	if tr.GetSize(5) != 8 {
		test.Errorf("Size should be 8 but %v", tr.GetSize(5))
	}
	if tr.GetSize(4) != 4 {
		test.Errorf("Size should be 4 but %v", tr.GetSize(4))
	}
	if tr.GetSize(7) != 3 {
		test.Errorf("Size should be 3 but %v", tr.GetSize(7))
	}
	if tr.GetSize(2) != 3 {
		test.Errorf("Size should be 3 but %v", tr.GetSize(2))
	}
	if tr.GetSize(6) != 1 {
		test.Errorf("Size should be 1 but %v", tr.GetSize(6))
	}
	if tr.GetSize(8) != 1 {
		test.Errorf("Size should be 1 but %v", tr.GetSize(8))
	}
	if tr.GetSize(1) != 1 {
		test.Errorf("Size should be 1 but %v", tr.GetSize(1))
	}
	if tr.GetSize(3) != 1 {
		test.Errorf("Size should be 1 but %v", tr.GetSize(3))
	}
}

func Test_GetHeight(test *testing.T) {
	tr := NewTree(5)
	tr.Inserts(7, 8, 4, 2, 1, 6, 3)
	if tr.GetHeight(5) != 3 {
		test.Errorf("Height should be 3 but %v", tr.GetHeight(5))
	}
	if tr.GetHeight(4) != 2 {
		test.Errorf("Height should be 2 but %v", tr.GetHeight(4))
	}
	if tr.GetHeight(7) != 1 {
		test.Errorf("Height should be 1 but %v", tr.GetHeight(7))
	}
	if tr.GetHeight(2) != 1 {
		test.Errorf("Height should be 1 but %v", tr.GetHeight(2))
	}
	if tr.GetHeight(1) != 0 {
		test.Errorf("Height should be 0 but %v", tr.GetHeight(1))
	}
	if tr.GetHeight(3) != 0 {
		test.Errorf("Height should be 0 but %v", tr.GetHeight(3))
	}
	if tr.GetHeight(6) != 0 {
		test.Errorf("Height should be 0 but %v", tr.GetHeight(6))
	}
	if tr.GetHeight(8) != 0 {
		test.Errorf("Height should be 0 but %v", tr.GetHeight(8))
	}
}

func Test_GetHeightRightLeft(test *testing.T) {
	tr1 := NewTree(4)
	tr1.Inserts(3, 2)
	if tr1.GetHeightRight(4) != 0 {
		test.Errorf("Height should be 0 but %v", tr1.GetHeightRight(4))
	}
	if tr1.GetHeightRight(3) != 0 {
		test.Errorf("Height should be 0 but %v", tr1.GetHeightRight(3))
	}
	if tr1.GetHeightRight(2) != 0 {
		test.Errorf("Height should be 0 but %v", tr1.GetHeightRight(2))
	}
	tr2 := NewTree(5)
	tr2.Inserts(7, 8, 4, 2, 1, 6, 3)
	if tr2.GetHeightRight(5) != 2 {
		test.Errorf("Height should be 2 but %v", tr2.GetHeightRight(5))
	}
	if tr2.GetHeightLeft(5) != 3 {
		test.Errorf("Height should be 3 but %v", tr2.GetHeightLeft(5))
	}
	if tr2.GetHeightRight(7) != 1 {
		test.Errorf("Height should be 1 but %v", tr2.GetHeightRight(7))
	}
	if tr2.GetHeightLeft(7) != 1 {
		test.Errorf("Height should be 1 but %v", tr2.GetHeightLeft(7))
	}
	if tr2.GetHeightRight(4) != 0 {
		test.Errorf("Height should be 0 but %v", tr2.GetHeightRight(4))
	}
	if tr2.GetHeightLeft(4) != 2 {
		test.Errorf("Height should be 2 but %v", tr2.GetHeightLeft(4))
	}
	if tr2.GetHeightRight(2) != 1 {
		test.Errorf("Height should be 0 but %v", tr2.GetHeightRight(2))
	}
	if tr2.GetHeightLeft(2) != 1 {
		test.Errorf("Height should be 0 but %v", tr2.GetHeightLeft(2))
	}
	if tr2.GetHeightRight(1) != 0 {
		test.Errorf("Height should be 0 but %v", tr2.GetHeightRight(1))
	}
	if tr2.GetHeightLeft(1) != 0 {
		test.Errorf("Height should be 0 but %v", tr2.GetHeightLeft(1))
	}
	if tr2.GetHeightRight(3) != 0 {
		test.Errorf("Height should be 0 but %v", tr2.GetHeightRight(3))
	}
	if tr2.GetHeightLeft(3) != 0 {
		test.Errorf("Height should be 0 but %v", tr2.GetHeightLeft(3))
	}
	if tr2.GetHeightRight(6) != 0 {
		test.Errorf("Height should be 0 but %v", tr2.GetHeightRight(6))
	}
	if tr2.GetHeightLeft(6) != 0 {
		test.Errorf("Height should be 0 but %v", tr2.GetHeightLeft(6))
	}
	if tr2.GetHeightRight(8) != 0 {
		test.Errorf("Height should be 0 but %v", tr2.GetHeightRight(8))
	}
	if tr2.GetHeightLeft(8) != 0 {
		test.Errorf("Height should be 0 but %v", tr2.GetHeightLeft(8))
	}
}

func Test_Height(test *testing.T) {
	tr1 := NewTree(4)
	tr1.Inserts(3, 2)
	if tr1.Height(4) != 2 {
		test.Errorf("Height should be 2 but %v", tr1.Height(4))
	}
	tr2 := NewTree(4)
	tr2.Inserts(5, 6)
	if tr2.Height(4) != -2 {
		test.Errorf("Height should be -2 but %v", tr2.Height(4))
	}
}

func Test_IsBalanced(test *testing.T) {
	tr1 := NewTree(4)
	tr1.Inserts(3, 2)
	if tr1.IsBalanced(4) {
		test.Errorf("IsBalanced should be false but %v", tr1.IsBalanced(4))
	}
	tr2 := NewTree(4)
	tr2.Inserts(5, 6)
	if tr2.IsBalanced(4) {
		test.Errorf("IsBalanced should be false but %v", tr2.IsBalanced(4))
	}
	tr3 := NewTree(5)
	tr3.Inserts(7, 8, 4, 2, 1, 3)
	if !tr3.IsBalanced(5) {
		test.Errorf("IsBalanced should be true but %v", tr3.IsBalanced(5))
	}
	if tr3.IsBalanced(4) {
		test.Errorf("IsBalanced should be false but %v", tr3.IsBalanced(4))
	}
	if !tr3.IsBalanced(7) {
		test.Errorf("IsBalanced should be true but %v", tr3.IsBalanced(7))
	}
	if !tr3.IsBalanced(8) {
		test.Errorf("IsBalanced should be true but %v", tr3.IsBalanced(8))
	}
	if !tr3.IsBalanced(2) {
		test.Errorf("IsBalanced should be true but %v", tr3.IsBalanced(2))
	}
	if !tr3.IsBalanced(1) {
		test.Errorf("IsBalanced should be true but %v", tr3.IsBalanced(1))
	}
	if !tr3.IsBalanced(3) {
		test.Errorf("IsBalanced should be true but %v", tr3.IsBalanced(3))
	}
}

func Test_BalanceInsert(test *testing.T) {
	tr1 := NewTree(4)
	_, tr1_r1 := tr1.BalanceInsert(3)
	if tr1_r1 != "Balanced" {
		test.Errorf("BalanceInsert should return Balanced but %v", tr1_r1)
	}
	_, tr1_r2 := tr1.BalanceInsert(2)
	if tr1_r2 != "LL" {
		test.Errorf("BalanceInsert should return LL but %v", tr1_r2)
	}

	tr2 := NewTree(4)
	_, tr2_r1 := tr2.BalanceInsert(5)
	if tr2_r1 != "Balanced" {
		test.Errorf("BalanceInsert should return Balanced but %v", tr2_r1)
	}
	_, tr2_r2 := tr2.BalanceInsert(6)
	if tr2_r2 != "RR" {
		test.Errorf("BalanceInsert should return LL but %v", tr2_r2)
	}

	tr3 := NewTree(4)
	_, tr3_r1 := tr3.BalanceInsert(1)
	if tr3_r1 != "Balanced" {
		test.Errorf("BalanceInsert should return Balanced but %v", tr3_r1)
	}
	_, tr3_r2 := tr3.BalanceInsert(3)
	if tr3_r2 != "LR" {
		test.Errorf("BalanceInsert should return LR but %v", tr3_r2)
	}

	tr4 := NewTree(4)
	_, tr4_r1 := tr4.BalanceInsert(7)
	if tr4_r1 != "Balanced" {
		test.Errorf("BalanceInsert should return Balanced but %v", tr4_r1)
	}
	_, tr4_r2 := tr4.BalanceInsert(5)
	if tr4_r2 != "RL" {
		test.Errorf("BalanceInsert should return RL but %v", tr4_r2)
	}

	tr5 := NewTree(5)
	tr5.Inserts(7, 8, 4, 2, 1, 6, 3)
	_, tr5_r1 := tr5.BalanceInsert(9)
	if tr5_r1 != "Balanced" {
		test.Errorf("BalanceInsert should return Balanced but %v", tr5_r1)
	}
	_, tr5_r2 := tr5.BalanceInsert(10)
	if tr5_r2 != "RR" {
		test.Errorf("BalanceInsert should return RR but %v", tr5_r2)
	}
}

func Test_BalanceLL(test *testing.T) {
	tr := NewTree(4)
	tr.BalanceInsert(3)
	tr.BalanceInsert(2)
	// Now tr is unbalanced

	// Now the following line rebalances the tr
	tr.BalanceLL(2)
	if !tr.IsBalanced(4) {
		test.Errorf("BalanceLL should balance the tree but %v", tr.IsBalanced(4))
	}
	if tr.Find(4).GetSize(4) != 3 {
		test.Errorf("Should return 3 but %v", tr.Find(4).GetSize(4))
	}
	if tr.Find(3).GetSize(3) != 1 {
		test.Errorf("Should return 1 but %v", tr.Find(3).GetSize(3))
	}
	if tr.Find(2).GetSize(2) != 1 {
		test.Errorf("Should return 1 but %v", tr.Find(2).GetSize(2))
	}
}

func Test_BalanceLR(test *testing.T) {
	tr := NewTree(4)
	tr.BalanceInsert(2)
	tr.BalanceInsert(3)
	// Now tr is unbalanced

	// Now the following line rebalances the tr
	tr.BalanceLR(3)
	if !tr.IsBalanced(4) {
		test.Errorf("BalanceLR should balance the tree but %v", tr.IsBalanced(4))
	}
	if tr.Find(4).GetSize(4) != 3 {
		test.Errorf("Should return 3 but %v", tr.Find(4).GetSize(4))
	}
	if tr.Find(3).GetSize(3) != 1 {
		test.Errorf("Should return 1 but %v", tr.Find(3).GetSize(3))
	}
	if tr.Find(2).GetSize(2) != 1 {
		test.Errorf("Should return 1 but %v", tr.Find(2).GetSize(2))
	}
}

func Test_BalanceRR(test *testing.T) {
	tr := NewTree(4)
	tr.BalanceInsert(5)
	tr.BalanceInsert(6)
	// Now tr is unbalanced

	// Now the following line rebalances the tr
	tr.BalanceRR(6)
	if !tr.IsBalanced(4) {
		test.Errorf("BalanceRR should balance the tree but %v", tr.IsBalanced(4))
	}
	if tr.Find(4).GetSize(4) != 3 {
		test.Errorf("Should return 3 but %v", tr.Find(4).GetSize(4))
	}
	if tr.Find(5).GetSize(5) != 1 {
		test.Errorf("Should return 1 but %v", tr.Find(5).GetSize(5))
	}
	if tr.Find(6).GetSize(6) != 1 {
		test.Errorf("Should return 1 but %v", tr.Find(6).GetSize(6))
	}
}

func Test_BalanceRL(test *testing.T) {
	tr := NewTree(4)
	tr.BalanceInsert(6)
	tr.BalanceInsert(5)
	// Now tr is unbalanced

	// Now the following line rebalances the tr
	tr.BalanceRL(5)
	if !tr.IsBalanced(4) {
		test.Errorf("BalanceRL should balance the tree but %v", tr.IsBalanced(4))
	}
	if tr.Find(4).GetSize(4) != 3 {
		test.Errorf("Should return 3 but %v", tr.Find(4).GetSize(4))
	}
	if tr.Find(5).GetSize(5) != 1 {
		test.Errorf("Should return 1 but %v", tr.Find(5).GetSize(5))
	}
	if tr.Find(6).GetSize(6) != 1 {
		test.Errorf("Should return 1 but %v", tr.Find(6).GetSize(6))
	}
}
