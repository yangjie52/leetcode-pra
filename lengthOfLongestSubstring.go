package main

import "fmt"

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
//
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
func main() {
	s := "abac"
	result := lengthOfLongestSubstring(s)
	fmt.Println(result)
}
func lengthOfLongestSubstring(s string) int {
	var char []string
	for _, v := range char {
		char = append(char, string(v))
	}
	return 0
}

//func lengthOfLongestSubstring(s string) int {
//	var result []int
//	char := make(map[string]int)
//	if len(s) == 0 {
//		return 0
//	}
//	for _, v1 := range s {
//		char[string(v1)]++
//		if char[string(v1)] == 2 {
//			result = append(result, len(char))
//			for key := range char {
//				if char[key] == 1 {
//					delete(char, key)
//				} else {
//					char[key] = 1
//				}
//			}
//		}
//		result = append(result, len(char))
//	}
//	max := result[0]
//	for _, v := range result {
//		if int(v) > max {
//			max = v
//		}
//	}
//	return max
//}
