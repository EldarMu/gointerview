// Package gointerview is a collection of interview problems solved in go
package gointerview

import (
	dataStructs "gointerview/datastructs"
)

type TreeNode = dataStructs.TreeNode

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
