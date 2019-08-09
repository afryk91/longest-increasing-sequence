package main

import (
	"fmt"
)

func findLIS(list []int) []int {
	if len(list) == 0 {
		return []int{}
	}

	var lengthOfLongest int = 0
	var subsequenceEndingsIdxs = make([]int, len(list) + 1)
	var predecesorIdxs = make([]int, len(list))
	

	constructLIS := func () []int {
		var subseq = make([]int, lengthOfLongest)
		k := subsequenceEndingsIdxs[lengthOfLongest]
	
		for j:= lengthOfLongest - 1; j >= 0; j-- {
			subseq[j] = list[k]
			k = predecesorIdxs[k]
		}
		return subseq
	}

	lengthOfLongestSubseqEndingWithItem := func (item int) int {
		left := 1
		right := lengthOfLongest
		for left <= right {
			var mid = int((left + right) / 2)
			if list[subsequenceEndingsIdxs[mid]] < item {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	
		return left
	}

	for i, item := range list {
		newLength := lengthOfLongestSubseqEndingWithItem(item)

		subsequenceEndingsIdxs[newLength] = i
		predecesorIdxs[i] = subsequenceEndingsIdxs[newLength - 1]

		if newLength > lengthOfLongest {
			lengthOfLongest = newLength
		}
	}

	return constructLIS()
}

func main() {
	var input = []int{1, 9, 5, 13, 3, 11, 7, 15, 2, 10, 6, 14, 4, 12, 8, 16}
	var increasing = []int{1,2,3,4,5,6}
	var decreasing = []int{6,5,4,3,2,1}
	var constant = []int{2,2,2,2,2,2}
	var empty = []int{}
	var repeating = []int{1,2,3,1,2,3,1,2,3}
	var withNegative = []int{2,-1,3,6,4,8}
	
	
	fmt.Println(findLIS(input), "should be [1,3,7,10,12,16]")
	fmt.Println(findLIS(increasing), "should be [1,2,3,4,5,6]")
	fmt.Println(findLIS(decreasing), "should be [1]")
	fmt.Println(findLIS(constant), "should be [2]")
	fmt.Println(findLIS(empty), "should be []")
	fmt.Println(findLIS(repeating), "should be [1,2,3]")
	fmt.Println(findLIS(withNegative), "should be [-1,3,4,8]")
}
