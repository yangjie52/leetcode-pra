package main

import "fmt"

// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

// 递归方式
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

// 非递归方式
func climbStairs1(n int) int {
	var methods = []int{1, 2}
	var result int
	if n <= 2 {
		result = n
	}
	for i := 2; i < n; i++ {
		methods = append(methods, methods[i-1]+methods[i-2])
	}
	result = methods[n-1]
	return result
}

func main() {
	n := 44
	result := climbStairs1(n)
	fmt.Println(result)
}
