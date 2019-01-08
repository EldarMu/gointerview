//Package gointerview is a collection of interview problems solved in go
package gointerview

import (
	"testing"
)

//the following are unit tests for the functions in interviewProblems
//in the same order as the functions
//
func TestShortestToChar(t *testing.T) {
	expected := [12]int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0}
	var result = shortestToChar("loveleetcode", 'e')
	for i := 0; i < len(result); i++ {
		if expected[i] != result[i] {
			t.Error("Incorrect value, expected", expected[i], "but got", result[i], "at index", i)
		}
	}
}

func TestBackspaceCompare(t *testing.T) {
	vals := [5][2]string{{"ab#d", "ad#d"}, {"ab##", "c#d#"}, {"a##c", "#a#c"}, {"a#c", "b"}, {"nzp#o#g", "b#nzp#o#g"}}
	expected := [5]bool{true, true, true, false, true}

	for i := 0; i < len(vals); i++ {
		if expected[i] != backspaceCompare(vals[i][0], vals[i][1]) {
			t.Error(vals[i][0], "and", vals[i][1], "match was", !expected[i], "expected", expected[i])
		}
	}
}

func TestRemoveDuplicates(t *testing.T) {
	input := []int{1, 1, 2}
	expected := []int{1, 2}

	result := removeDuplicates(input)
	if len(expected) != result {
		t.Error("incorrect result length, expected", len(expected), "but got", result)
	}

	for i := 0; i < len(expected); i++ {
		if input[i] != expected[i] {
			t.Error("incorrectly sorted at index", i, "expected", expected[i], "but got", input[i])
		}
	}

	input = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	expected = []int{0, 1, 2, 3, 4}
	result = removeDuplicates(input)
	if len(expected) != result {
		t.Error("incorrect result length, expected", len(expected), "but got", result)
	}

	for i := 0; i < len(expected); i++ {
		if input[i] != expected[i] {
			t.Error("incorrectly sorted at index", i, "expected", expected[i], "but got", input[i])
		}
	}
}

func TestSumOfLeftLeaf(t *testing.T) {
	root := TreeNode{Val: 3}
	left := TreeNode{Val: 9}
	right := TreeNode{Val: 20}
	root.Left = &left
	root.Right = &right
	rleft := TreeNode{Val: 15}
	rright := TreeNode{Val: 7}
	root.Right.Left = &rleft
	root.Right.Right = &rright

	expected := 24
	result := sumOfLeftLeaves(&root)
	if expected != result {
		t.Error("incorrect sum of left leaves, expected", expected, "but got", result)
	}

	if sumOfLeftLeaves(nil) != 0 {
		t.Error("null root should return 0")
	}

	lright := TreeNode{Val: 4}
	left.Right = &lright
	lrleft := TreeNode{Val: -4}
	lright.Left = &lrleft
	expected = 11
	result = sumOfLeftLeaves(&root)
	if expected != result {
		t.Error("incorrect sum of left leaves, expected", expected, "but got", result)
	}
}

func TestDirReduc(t *testing.T) {
	var inp []string
	var expec []string
	var res []string

	inp = []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "NORTH"}
	expec = []string{"NORTH"}
	res = dirReduc(inp)
	compareStrArrResult(expec, res, t)

	inp = []string{"SOUTH", "SOUTH", "WEST", "NORTH", "WEST"}
	expec = []string{"SOUTH", "SOUTH", "WEST", "NORTH", "WEST"}
	res = dirReduc(inp)
	compareStrArrResult(expec, res, t)
}

func compareStrArrResult(expec, res []string, t *testing.T) {
	if len(expec) != len(res) {
		t.Error("size mismatch between expected array size", len(expec), "and result size", len(res))
	} else {
		for i := range expec {
			if expec[i] != res[i] {
				t.Error("result mismatch at index", i, "expected", expec[i], "got", res[i])
			}
		}
	}
}

func compareIntArrResult(expec, res []int, t *testing.T) {
	if len(expec) != len(res) {
		t.Error("size mismatch between expected array size", len(expec), "and result size", len(res))
	} else {
		for i := range expec {
			if expec[i] != res[i] {
				t.Error("result mismatch at index", i, "expected", expec[i], "got", res[i])
			}
		}
	}
}

func genLN(a []int) *ListNode {
	head := new(ListNode)
	cur := head
	for i, v := range a {
		cur.Val = v
		if i != len(a)-1 {
			cur.Next = new(ListNode)
			cur = cur.Next
		}
	}
	return head
}

func compareIntMtrxResult(exp, res [][]int, t *testing.T) {
	if len(exp) != len(res) {
		t.Error("matrix size mismatch, expected", len(exp), "got", len(res))
	} else {
		for i := range exp {
			compareIntArrResult(exp[i], res[i], t)
		}
	}
}

func compareIntResult(exp, res int, t *testing.T) {
	if exp != res {
		t.Error("result mismatch, expected", exp, "but got", res)
	}
}

func compareTreeResult(exp, res *TreeNode, t *testing.T) {
	if exp == nil && res == nil {
		return
	} else if exp == nil && res != nil {
		t.Error("tree mismatch, expected end of branch, got", res.Val)
	} else if exp != nil && res == nil {
		t.Error("tree mismatch at tree node with value", exp.Val, "result cuts off branch")
	} else if exp.Val != res.Val {
		t.Error("tree mismatch at tree node with val", exp.Val, "result was", res.Val)
	} else {
		if exp.Left != nil {
			compareTreeResult(exp.Left, res.Left, t)
		}
		if exp.Right != nil {
			compareTreeResult(exp.Right, res.Right, t)
		}
	}
}

func compareListNodeResult(exp, res *ListNode, t *testing.T) {
	for exp != nil || res != nil {
		if exp == nil {
			t.Error("Result list goes too long, final value of result was", res.Val)
			return
		} else if res == nil {
			t.Error("Result list cuts short, last expected value was", exp.Val)
			return
		}
		if exp.Val != res.Val {
			t.Error("list value mismatch, expected", exp.Val, "got", res.Val)
		}
		exp = exp.Next
		res = res.Next
	}
}

func TestNumJewelsInStones(t *testing.T) {
	var res int
	var exp int

	res = numJewelsInStones("aAbBZ", "afuinefdsnBdasZA")
	exp = 5
	compareIntResult(exp, res, t)

	res = numJewelsInStones("", "")
	exp = 0
	compareIntResult(exp, res, t)
}

func TestMaxIncreaseKeepingSkyline(t *testing.T) {
	var inp [][]int
	var res int
	var exp int

	inp = [][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}
	res = maxIncreaseKeepingSkyline(inp)
	exp = 35
	compareIntResult(exp, res, t)
}

func TestUniqueMorseRepresentations(t *testing.T) {
	var inp []string
	var exp int
	var res int

	inp = []string{"gin", "zen", "gig", "msg"}
	exp = 2
	res = uniqueMorseRepresentations(inp)
	compareIntResult(exp, res, t)
}

func TestGetSum(t *testing.T) {
	var res int
	var exp int

	exp = 0
	res = getSum(2, -2)
	compareIntResult(exp, res, t)

	exp = 5
	res = getSum(2, 3)
	compareIntResult(exp, res, t)

	exp = 3
	res = getSum(1, 2)
	compareIntResult(exp, res, t)
}

func TestTopKFrequent(t *testing.T) {
	var inp []int
	var res []int
	var exp []int

	inp = []int{5, 2, 5, 3, 5, 3, 1, 1, 3}
	exp = []int{3, 5}
	res = topKFrequent(inp, 2)
	compareIntArrResult(exp, res, t)

	inp = []int{1, 1, 1, 2, 2, 3}
	exp = []int{1, 2}
	res = topKFrequent(inp, 2)
	compareIntArrResult(exp, res, t)

	inp = []int{1}
	exp = []int{1}
	res = topKFrequent(inp, 1)
	compareIntArrResult(exp, res, t)
}

func TestAltTopKFrequent(t *testing.T) {
	var inp []int
	var res []int
	var exp []int

	inp = []int{1, 1, 1, 2, 2, 3}
	exp = []int{1, 2}
	res = altTopKFrequent(inp, 2)
	compareIntArrResult(exp, res, t)

	inp = []int{5, 2, 5, 3, 5, 3, 1, 1, 3}
	exp = []int{3, 5}
	res = altTopKFrequent(inp, 2)
	compareIntArrResult(exp, res, t)

	inp = []int{1}
	exp = []int{1}
	res = altTopKFrequent(inp, 1)
	compareIntArrResult(exp, res, t)
}

func TestConstructMaximumBinaryTree(t *testing.T) {
	var inp []int
	var exp *TreeNode
	var res *TreeNode

	inp = []int{3, 2, 1, 6, 0, 5}

	exp = &TreeNode{Val: 6}
	exp.Left = &TreeNode{Val: 3}
	exp.Right = &TreeNode{Val: 5}
	exp.Left.Right = &TreeNode{Val: 2}
	exp.Right.Left = &TreeNode{Val: 0}
	exp.Left.Right.Right = &TreeNode{Val: 1}

	res = constructMaximumBinaryTree(inp)
	compareTreeResult(exp, res, t)
}

func TestFlipAndInvertImage(t *testing.T) {
	var inp [][]int
	var exp [][]int
	var res [][]int

	inp = [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}
	exp = [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}}
	res = flipAndInvertImage(inp)
	compareIntMtrxResult(exp, res, t)

	inp = [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}
	exp = [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}}
	res = flipAndInvertImage(inp)
	compareIntMtrxResult(exp, res, t)
}

func TestTwoSum(t *testing.T) {
	var inp []int
	var exp []int
	var res []int

	inp = []int{2, 7, 11, 15}
	exp = []int{0, 1}
	res = twoSum(inp, 9)
	compareIntArrResult(exp, res, t)
}

func TestAddTwoNumbers(t *testing.T) {
	var l1 *ListNode
	var l2 *ListNode
	var exp *ListNode
	var res *ListNode

	l1 = &ListNode{Val: 0}
	l2 = genLN([]int{1, 8})
	exp = genLN([]int{1, 8})
	res = addTwoNumbers(l1, l2)
	compareListNodeResult(exp, res, t)

	l1 = &ListNode{Val: 5}
	l2 = &ListNode{Val: 5}
	exp = genLN([]int{0, 1})
	res = addTwoNumbers(l1, l2)
	compareListNodeResult(exp, res, t)

	l1 = genLN([]int{2, 4, 7})
	l2 = genLN([]int{5, 6, 4})
	exp = genLN([]int{7, 0, 2, 1})
	res = addTwoNumbers(l1, l2)
	compareListNodeResult(exp, res, t)

}
