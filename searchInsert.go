package main

import (
	"fmt"
)

// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
// 请必须使用时间复杂度为 O(log n) 的算法。
// 输入: nums = [1,3,5,6], target = 2
// 输出: 1
func searchInsert1(nums []int, target int) int {
	var subscript int
	if nums[0] >= target {
		subscript = 0
	} else if nums[len(nums)-1] < target {
		subscript = len(nums)

	} else {
		for i := 0; i < len(nums); i++ {
			if nums[i] == target {
				subscript = i
			} else if nums[i] < target && nums[i+1] > target {
				subscript = i + 1
			}
		}
	}
	return subscript
}

func searchInsert(nums []int, target int) int {
	var subscript int
	var left, right, merge int
	left = 0
	right = len(nums) - 1
	if nums[0] >= target {
		subscript = 0
	} else if nums[len(nums)-1] < target {
		subscript = len(nums)
	} else if nums[len(nums)-1] == target {
		subscript = len(nums) - 1
	} else {
		for right-left > 1 {
			merge = (left + right) / 2
			//fmt.Println(merge)
			if target < nums[merge] {
				right = merge
			} else if target > nums[merge] {
				left = merge
			} else {
				subscript = merge
				break
			}
		}
		if right-left == 1 && left < (len(nums)-1)/2 {

			if target <= nums[left] {
				subscript = left
			} else {
				subscript = left + 1
			}
		}
		if right-left == 1 && left >= (len(nums)-1)/2 {
			if target > nums[right] {
				subscript = right + 1
			} else {
				subscript = right
			}
		}

	}
	return subscript
}

//func searchInsert(nums []int, target int) int {
//	var subscript int
//	var left, right, merge int
//	left = 0
//	right = len(nums) - 1
//	if len(nums) == 1 {
//		if nums[0] >= target {
//			subscript = 0
//		} else {
//			subscript = 1
//		}
//
//	}
//	for right-left > 1 {
//		merge = (left + right) / 2
//		//fmt.Println(merge)
//		if target < nums[merge] {
//			right = merge
//		} else if target > nums[merge] {
//			left = merge
//		} else {
//			subscript = merge
//			break
//		}
//	}
//	if right-left == 1 && left < (len(nums)-1)/2 {
//
//		if target <= nums[left] {
//			subscript = left
//		} else {
//			subscript = left + 1
//		}
//	}
//	if right-left == 1 && left > (len(nums)-1)/2 {
//		if target > nums[right] {
//			subscript = right + 1
//		} else {
//			subscript = right
//		}
//	}
//	if len(nums) == 2 && right-left == 1 && left == (len(nums)-1)/2 {
//		if target <= nums[left] {
//			subscript = 0
//		} else if target > nums[left] && target <= nums[right] {
//			subscript = 1
//		} else {
//			subscript = 2
//		}
//	}
//
//	return subscript
//
//}

func main() {
	var nums []int = []int{1}
	result := searchInsert1(nums, 7)
	fmt.Println(result)
}
