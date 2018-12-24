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
	sIndex := len(S) - 1
	tIndex := len(T) - 1
	for sIndex > -1 && tIndex > -1 {
		sIndex = getIndOfNextChar(S, sIndex)
		tIndex = getIndOfNextChar(T, tIndex)
		if sIndex == -1 || tIndex == -1 {
			break
		}
		if S[sIndex] != T[tIndex] {
			return false
		}
		sIndex--
		tIndex--
	}
	if sIndex > -1 && getIndOfNextChar(S, sIndex) > -1 {
		return false
	}
	if tIndex > -1 && getIndOfNextChar(T, tIndex) > -1 {
		return false
	}
	return true
}

func getIndOfNextChar(S string, index int) int {
	backCount := 0
	for index > -1 && (backCount != 0 || S[index] == '#') {
		if S[index] == '#' {
			backCount++
		} else {
			backCount--
		}
		index--
	}
	return index
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

//redundant directions problem
//https://www.codewars.com/kata/directions-reduction/train/go
//no performance data from codewars
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

//Given a string of distinct characters J, return a count of those characters appearing in string S
//https://leetcode.com/problems/jewels-and-stones/description/
//simply linear time solution with map
//beats 100% of go submissions (0 ms for 254 unit tests)
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