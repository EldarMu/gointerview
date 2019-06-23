// Package gointerview is a collection of interview problems solved in go
package gointerview

import (
	"container/heap"
	dataStructs "gointerview/datastructs"
	"math"
	"sort"
	"strconv"
	"strings"
)

// TreeNode is a basic container for a binary tree with an int value
type TreeNode = dataStructs.TreeNode

// ListNode is a basic container for a singly-linked list with an int value
type ListNode = dataStructs.ListNode

// ValCount is a container for int values and their corresponding counts in some sort of array
type ValCount = dataStructs.ValCount

// PriorityQueue is a data structure for working with a maxheap
type PriorityQueue = dataStructs.PriorityQueue

// shortestToChar is a function that, given a string and a character,
// returns an int arr that shows the distance
// to the nearest instance of that character for each char in string
// we are guaranteed at least a single instance of the requested character
// https://leetcode.com/problems/shortest-distance-to-a-character/
// 4 ms, beats 100% of golang submissions
func shortestToChar(s string, c byte) []int {
	fromLeft := make([]int, len(s))
	fromRight := make([]int, len(s))
	lastLeftMatch := -1
	lastRightMatch := len(s)

	for i := range s {
		if s[i] == c {
			lastLeftMatch = i
		}
		fromLeft[i] = lastLeftMatch
		rx := len(s) - 1 - i
		if s[rx] == c {
			lastRightMatch = rx
		}
		fromRight[rx] = lastRightMatch
	}

	r := make([]int, len(s))
	for j := range s {
		if fromLeft[j] == -1 {
			r[j] = fromRight[j] - j
			continue
		}
		if fromRight[j] == len(s) {
			r[j] = j - fromLeft[j]
			continue
		}
		if j-fromLeft[j] <= fromRight[j]-j {
			r[j] = j - fromLeft[j]
			continue
		}
		r[j] = fromRight[j] - j
	}
	return r
}

// given two strings as a series of characters with # representing backspace (prev char was deleted)
// check if the two strings are identical
// https://leetcode.com/problems/backspace-string-compare/
// 0 ms, beats 100% of go submissions
func backspaceCompare(S string, T string) bool {
	sx := len(S) - 1
	tx := len(T) - 1
	for sx > -1 && tx > -1 {
		sx = nextCharX(S, sx)
		tx = nextCharX(T, tx)
		if sx == -1 || tx == -1 {
			break
		}
		if S[sx] != T[tx] {
			return false
		}
		sx--
		tx--
	}
	if sx > -1 && nextCharX(S, sx) > -1 {
		return false
	}
	if tx > -1 && nextCharX(T, tx) > -1 {
		return false
	}
	return true
}

func nextCharX(S string, X int) int {
	backCount := 0
	for X > -1 && (backCount != 0 || S[X] == '#') {
		if S[X] == '#' {
			backCount++
		} else {
			backCount--
		}
		X--
	}
	return X
}

// remove duplicates from sorted array in-place, return size of sub-array with sorted vals
// 60 ms for 160 test cases, beats 100% of golang submissions
func removeDuplicates(nums []int) int {
	nextWrite := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[nextWrite] = nums[i]
			nextWrite++
		}
	}
	return nextWrite
}

// given a binary tree of int values, return the sum of the left leaves
// beats 100% of golang submissions (0 ms for 102 unit tests)
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			sum += root.Left.Val
		} else {
			sum += sumOfLeftLeaves(root.Left)
		}
	}
	if root.Right != nil {
		sum += sumOfLeftLeaves(root.Right)
	}
	return sum
}

// redundant directions problem
// https://www.codewars.com/kata/directions-reduction/train/go
// no performance data from codewars
func dirReduc(inp []string) []string {
	revDir := map[string]string{
		"NORTH": "SOUTH",
		"SOUTH": "NORTH",
		"EAST":  "WEST",
		"WEST":  "EAST",
	}

	res := make([]string, 0, len(inp))
	lastWrite := -1
	for _, s := range inp {
		if lastWrite > -1 && s == revDir[res[lastWrite]] {
			res = res[:len(res)-1]
			lastWrite--
		} else {
			res = append(res, s)
			lastWrite++
		}
	}
	return res
}

// Given a string of distinct characters J, return a count of those characters appearing in string S
// https://leetcode.com/problems/jewels-and-stones/description/
// simply linear time solution with map
// beats 100% of go submissions (0 ms for 254 unit tests)
func numJewelsInStones(J string, S string) int {
	m := make(map[byte]int)
	for i := range J {
		m[J[i]] = 1
	}
	res := 0
	for i := range S {
		res += m[S[i]]
	}
	return res
}

// given a 2d array of positive ints, find by how much all the elements in the array can be increased
// while preserving the maximum value of each row and column
// https://leetcode.com/problems/max-increase-to-keep-city-skyline/description/
// this solution assumes a non-jagged slice of slices
// beats 100% of go solutions at 8 ms for 133 test cases
func maxIncreaseKeepingSkyline(grid [][]int) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	row := make([]int, len(grid))
	col := make([]int, len(grid[0]))

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] > row[i] {
				row[i] = grid[i][j]
			}
			if grid[i][j] > col[j] {
				col[j] = grid[i][j]
			}
		}
	}

	res := 0
	for i := range grid {
		for j := range grid[i] {
			if row[i] <= col[j] {
				if grid[i][j] < row[i] {
					res += row[i] - grid[i][j]
				}
			} else {
				if grid[i][j] < col[j] {
					res += col[j] - grid[i][j]
				}
			}
		}
	}

	return res
}

// given a list of words and the morse code for the letters, many words will have the same representation
// find all unique representations in the list
// https://leetcode.com/problems/unique-morse-code-words/description/
// beats 100% of golang submissions at 0 ms for 83 test cases
func uniqueMorseRepresentations(words []string) int {
	c := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---",
		"-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--",
		"-..-", "-.--", "--.."}
	m := make(map[string]bool)
	for i := range words {
		w := make([]byte, 0, len(words[i])*4)
		for j := range words[i] {
			w = append(w, c[int(words[i][j])-int('a')]...)
		}
		if !m[string(w)] {
			m[string(w)] = true
		}
	}
	return len(m)
}

// calculate sum of a and b without using + or -
// using bit operations is pretty straightforward here
// https://leetcode.com/problems/sum-of-two-integers/
func getSum(a int, b int) int {
	carry := 0
	for b != 0 {
		carry = (a & b) << 1
		a = a ^ b
		b = carry
	}
	return a
}

// return k most frequent elements in slice
// https://leetcode.com/problems/top-k-frequent-elements/description/
// let's try with a maxheap
// beats 40% of golang submissions at 24 ms for 21 test cases
// also pretty awful to read
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]*ValCount)
	ok := false
	pq := make(PriorityQueue, 0, len(nums))
	for i := range nums {
		_, ok = m[nums[i]]
		if ok {
			vc := m[nums[i]]
			pq.Update(vc, m[nums[i]].Count+1)

		} else {
			item := &ValCount{Value: nums[i], Count: 1, Index: -1}
			m[nums[i]] = item
			pq.Push(item)
		}
	}
	res := make([]int, 0, k)
	for j := 0; j < k; j++ {
		res = append(res, heap.Pop(&pq).(*ValCount).Value)
	}
	return res
}

// beats 100% of golang submissions
// at 20 ms for 21 test cases
func altTopKFrequent(nums []int, k int) []int {
	vMap := make(map[int]int)
	for _, i := range nums {
		vMap[i]++
	}

	counts := make([]int, 0)
	cMap := make(map[int][]int)
	for v, c := range vMap {
		if len(cMap[c]) == 0 {
			counts = append(counts, c)
		}
		cMap[c] = append(cMap[c], v)
	}

	res := make([]int, 0, k)

	sort.Ints(counts)
	add := 0
	ind := len(counts) - 1
	for add < k && ind >= 0 {
		vs := cMap[counts[ind]]
		sort.Ints(vs)
		for _, i := range vs {
			if add == k {
				return res
			}
			res = append(res, i)
			add++
		}
		ind--
	}
	return res
}

// given an array of numbers, build a maximum binary tree in the following way:
// the max value becomes the root, the max value of the array slice to its left becomes
// the left child, and so forth..
// https://leetcode.com/problems/maximum-binary-tree/
// beats 100% of golang submissions at 68 ms for 107 test cases
func constructMaximumBinaryTree(nums []int) *TreeNode {
	return recurBuildMaxBinTree(nums, 0, len(nums)-1)
}

func recurBuildMaxBinTree(nums []int, start int, end int) *TreeNode {
	mxInd := start
	for i := start; i <= end; i++ {
		if nums[i] > nums[mxInd] {
			mxInd = i
		}
	}

	cur := &TreeNode{Val: nums[mxInd]}
	if start <= mxInd-1 {
		cur.Left = recurBuildMaxBinTree(nums, start, mxInd-1)
	}
	if mxInd+1 <= end {
		cur.Right = recurBuildMaxBinTree(nums, mxInd+1, end)
	}
	return cur
}

// given a two dimensional matrix of ones and zeros, flip it horizontally, then invert the values
// https://leetcode.com/problems/flipping-an-image
// beats 100% of golang submissions at 4 ms for 82 test cases
func flipAndInvertImage(A [][]int) [][]int {
	m := []int{1, 0}
	var t int
	l := len(A[0])
	odd := l%2 != 0
	for i := range A {
		for j := 0; j < l/2; j++ {
			t = A[i][l-1-j]
			A[i][l-1-j] = m[A[i][j]]
			A[i][j] = m[t]
		}
		if odd {
			A[i][l/2] = m[A[i][l/2]]
		}
	}
	return A
}

// given an array of ints and a sum,
// return indices of the two ints that together produce the sum
// guaranteed only one valid sol'n
// https://leetcode.com/problems/two-sum
// beats 100% of golang submissions at 4 ms for 29 test cases
func twoSum(nums []int, target int) []int {
	mi := make(map[int]int, len(nums))
	for i, v := range nums {
		x, ok := mi[target-v]
		if ok {
			return []int{x, i}
		}
		mi[v] = i

	}
	return []int{-1, -1}
}

// given two numbers in reverse order as two singly linked lists (e.g. 3->4->2 == 243)
// add them
// https://leetcode.com/problems/add-two-numbers
// 1563 test cases in 20 ms
// unable to verify code speed due to graphs being down
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := false
	head := new(ListNode)
	cur := head
	v1 := 0
	v2 := 0
	vt := 0
	for l1 != nil || l2 != nil || carry {
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		if carry {
			vt = 1
		}
		vt += v1 + v2
		if vt > 9 {
			vt %= 10
			carry = true
		} else {
			carry = false
		}
		cur.Val = vt
		if l1 != nil || l2 != nil || carry {
			cur.Next = new(ListNode)
			cur = cur.Next
		}
		v1 = 0
		v2 = 0
		vt = 0
	}
	return head
}

// given a string, return the length of the longest substring with no repeating characters
// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
// refactored to use a bool array of size 128
// 987 unit tests in 4 ms, can't get speed results due to graphs being down, but quadrupled speed
func lengthOfLongestSubstring(s string) int {
	var m [128]bool
	l := 0
	t := 0
	res := 0
	for i := range s {
		c := s[i]
		for m[c] {
			m[s[t]] = false
			t++
			l--
		}
		m[c] = true
		l++
		if res < l {
			res = l
		}
	}
	return res
}

// convert zig-zag text (string, with number of rows it's spread on)
// into the resulting flat string with rows read top to bottom
// https://leetcode.com/problems/zigzag-conversion/description/
// text first goes down, with 1 per row, then up diagonally, and repeat
// 8 ms sol'n, this took 5 times longer in java
func zigZagConvert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	rows := make([]strings.Builder, len(s))

	row := 0
	down := true
	for _, c := range s {
		rows[row].WriteRune(c)
		if row == 0 {
			down = true
		} else if row == numRows-1 {
			down = false
		}

		if down {
			row++
		} else {
			row--
		}
	}

	for i := 1; i < len(rows); i++ {
		rows[0].WriteString(rows[i].String())
	}
	return rows[0].String()
}

// HighAndLow is a function that,
// given a string consisting of valid space separated int32 values
// return a string with the max value and the min value, space separated
// solution to https://www.codewars.com/kata/highest-and-lowest/go
// unfortunately codewars has no speed measurements
func HighAndLow(in string) string {
	min := math.MaxInt32
	max := math.MinInt32
	minStr := ""
	maxStr := ""

	strArr := strings.Split(in, " ")

	for _, v := range strArr {
		num, _ := strconv.Atoi(v)
		if num <= min {
			min = num
			minStr = v
		}
		if num >= max {
			max = num
			maxStr = v
		}
	}

	return maxStr + " " + minStr
}

// DigPow should, given a number n consisting of digits a,b,c,d (e.g. 1,3,4 = 134) and an int p
// find if there is a number k such that (a^p + b^(p+1) + c ^(p+2))/k = n
// https://www.codewars.com/kata/playing-with-digits/train/go
func DigPow(n, p int) int {
	dig := p
	for tmp := n; tmp >= 10; tmp /= 10 {
		dig++
	}

	var sum int

	for tmp := n; tmp > 0; tmp /= 10 {
		v := tmp % 10
		cv := 1
		for i := 0; i < dig; i++ {
			cv *= v
		}
		sum += cv
		dig--
	}

	if sum%n == 0 {
		return sum / n
	}

	return -1
}

func ToCamelCase(s string) string {
	sb := strings.Builder{}

	f := func(c rune) bool {
		return c == '_' || c == '-'
	}
	strArr := strings.FieldsFunc(s, f)
	sb.WriteString(strArr[0])

	for i := 1; i < len(strArr); i++ {
		sb.WriteString(strings.Title(strArr[i]))
	}

	return sb.String()
}

// given a total volume, find out if the volume is precisely a sum of cubes
// of progressively larger numbers 1...n
// return n if it exists, else -1
func CalculateCubeSum(m int) int {
	if m == 0 {
		return 0
	} else if m < 0 {
		return -1
	}
	rem := m
	vol := 1
	for rem > 0 {
		rem -= vol * vol * vol
		vol++
	}

	if rem == 0 {
		return vol - 1
	} else {
		return -1
	}
}
