// 1375. 二进制字符串前缀一致的次数
// 给你一个长度为 n 、下标从 1 开始的二进制字符串，所有位最开始都是 0 。我们会按步翻转该二进制字符串的所有位（即，将 0 变为 1）。
// 给你一个下标从 1 开始的整数数组 flips ，其中 flips[i] 表示对应下标 i 的位将会在第 i 步翻转。
// 二进制字符串 前缀一致 需满足：在第 i 步之后，在 闭 区间 [1, i] 内的所有位都是 1 ，而其他位都是 0 。
// 返回二进制字符串在翻转过程中 前缀一致 的次数。
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a []int
	a = []int{3, 2, 4, 1, 5}
	//numTimesAllBlue1(a)
	result := numTimesAllBlue2(a)
	fmt.Println(result)
}
func numTimesAllBlue2(flips []int) int {

	var result int
	lenf := len(flips)
	comX := -1
	for i := 0; i < lenf; i++ {
		j := flips[i]
		if comX < j {
			comX = j
		}
		if comX == i+1 {
			result++
		}
	}
	return result

}

func numTimesAllBlue1(flips []int) int {
	var zero []int
	var result int
	for i := 0; i < len(flips); i++ {
		zero = append(zero, 0)
	}
	for i, v := range flips {
		zero[v-1] = 1
		first := zero[:i+1]
		last := zero[i+1:]
		expectedSlice1 := make([]int, i+1)
		for i := range expectedSlice1 {
			expectedSlice1[i] = 1
		}
		expectedSlice2 := make([]int, len(flips)-i-1)
		for i := range expectedSlice2 {
			expectedSlice2[i] = 0
		}
		//fmt.Printf("first:%s, expectedSlice1:%s;\n", first, expectedSlice1)
		//fmt.Printf("last:%s, expectedSlice2:%s;\n", last, expectedSlice2)
		if reflect.DeepEqual(first, expectedSlice1) && reflect.DeepEqual(last, expectedSlice2) {
			result += 1
			//fmt.Println(result)
		}
	}
	return result

}

func numTimesAllBlue(flips []int) int {
	var zero []int
	var result, x int
	for i := 0; i < len(flips); i++ {
		zero = append(zero, 0)
	}
	for i, v := range flips {
		zero[v-1] = 1
		x = 0
		for j, v1 := range zero {
			if v1 == 1 && j <= i {
				x += 1
			} else {
				x += 0
			}
			if x == i+1 && i == j {
				result += 1
			}
		}
	}
	return result
}
