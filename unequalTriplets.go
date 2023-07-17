package main

import "fmt"

//给你一个下标从 0 开始的正整数数组 nums 。请你找出并统计满足下述条件的三元组 (i, j, k) 的数目：
//
//0 <= i < j < k < nums.length
//nums[i]、nums[j] 和 nums[k] 两两不同 。
//换句话说：nums[i] != nums[j]、nums[i] != nums[k] 且 nums[j] != nums[k] 。
//返回满足上述条件三元组的数目。

func unequalTriplets(nums []int) int {
	num := 0
	if len(nums) < 3 {
		return 0
	}
	for i, v1 := range nums {
		for j, v2 := range nums {
			for k, v3 := range nums {
				if v1 != v2 && v2 != v3 && v1 != v3 && i < j && i < k && j < k {
					num += 1
				}
			}
		}
	}
	return num
}

func main() {
	var num = []int{2, 4, 8, 9}
	a := unequalTriplets(num)
	fmt.Println(a)
	result := unequals(num)
	for k, v := range result {
		fmt.Printf("%d %d, ", k, v)
	}
}

func unequals(nums []int) map[int]int {
	num := map[int]int{}
	for _, v := range nums {
		num[v]++
	}
	return num
}
