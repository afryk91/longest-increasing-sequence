package main

import (
	"fmt"
)

func findLIS(list []int) []int {
	if len(list) == 0 {
		return []int{}
	}

	var lengthOfLongest int = 1
	var subsequenceEndingsIdxs = make([]int, len(list))
	var predecesorIdxs = make([]int, len(list) + 1)
	

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
	fmt.Println(findLIS(input))
}
