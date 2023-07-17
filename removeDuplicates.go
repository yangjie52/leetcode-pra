package main

import "fmt"

func main() {
	var a []int
	a = []int{1, 2, 3, 1}
	//b = []int{}
	result := removeDuplicates(a)
	fmt.Println(result)
}
func removeDuplicates(nums []int) int {
	var reDup []int
	for _, v1 := range nums {
		exist := false
		for _, v2 := range reDup {
			if v1 == v2 {
				exist = true
				break
			}
		}
		if !exist {
			reDup = append(reDup, v1)
		}
	}
	copy(nums, reDup)
	//fmt.Println(nums)
	return len(reDup)
}
func insertIfNotExists(a []int, b []int) []int {
	for _, elementA := range a {
		exists := false
		for _, elementB := range b {
			if elementA == elementB {
				exists = true
				break
			}
		}
		if !exists {
			b = append(b, elementA)
		}
	}
	return b
}
