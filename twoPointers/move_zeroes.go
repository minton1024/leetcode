package main

import "fmt"

/*
*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

请注意 ，必须在不复制数组的情况下原地对数组进行操作。

示例 1:
输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]

示例 2:
输入: nums = [0]
输出: [0]
*/
func moveZeroes(nums []int) {
	lastNonZero := 0
	for _, num := range nums {
		if num != 0 {
			nums[lastNonZero] = num
			lastNonZero++
		}
	}

	for i := lastNonZero; i < len(nums); i++ {
		nums[i] = 0
	}
	fmt.Println(nums)
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	nums2 := []int{0}
	moveZeroes(nums2)
	nums3 := []int{1, 2, 0, 4, 0, 5}
	moveZeroes(nums3)
	nums4 := []int{0, 0, 1}
	moveZeroes(nums4)
}
