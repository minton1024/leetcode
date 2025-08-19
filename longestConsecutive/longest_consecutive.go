package main

import "fmt"

/*
*
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1：

输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
示例 2：

输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9
示例 3：

输入：nums = [1,0,1,2]
输出：3
*/
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}
	longestStreak := 0
	for num := range numSet {
		if !numSet[num-1] { // Only check for the start of a sequence
			currentNum := num
			currentStreak := 1
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}
			longestStreak = max(longestStreak, currentStreak)
		}
	}

	return longestStreak
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	result := longestConsecutive(nums)
	fmt.Println("Longest consecutive sequence length:", result)
	nums2 := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	result2 := longestConsecutive(nums2)
	fmt.Println("Longest consecutive sequence length:", result2)
}
