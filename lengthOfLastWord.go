package main

import "fmt"

// 给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。
//
// 单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。
func lengthOfLastWord1(s string) int {
	var length int
	var left, right int
	left = 0
	right = len(s)
	var nr, nl int
	nr = 0
	nl = 0

	for i := len(s) - 1; i > 0; i-- {
		if (s[i] == ' ' && s[i-1] != ' ') || (s[i] != ' ' && i == len(s)-1) {
			fmt.Println(i)
			if s[i] == ' ' {
				right = i
			} else {
				right = i + 1
			}

			fmt.Println(right)
			nr++
		}
		if (s[i] != ' ' && s[i-1] == ' ') || (s[i] != ' ' && i == 0) {
			left = i
			fmt.Println(left)
			nl++
		}
		if nl == 1 && nr == 1 {
			break
		}

	}
	fmt.Println(right)
	fmt.Println(left)
	length = right - left
	return length
}

//	func lengthOfLastWord(s string) int {
//		var length int
//		var arrs []string = []string{}
//		//var s1 []string = []string{}
//		//将末尾的空格去除掉
//		for i := len(s) - 1; i > 0; i-- {
//			if s[i] == ' ' {
//				s = s[:i]
//			} else {
//				break
//			}
//		}
//		//将每个单词以及后面跟的空格一起加入数组
//		for k, v := range s {
//			if string(v) == " " && k != len(s)-1 {
//				arrs = append(arrs, s[:k+1])
//			} else {
//				continue
//			}
//		}
//		//将整体加入数组末尾
//		arrs = append(arrs, s)
//		//整体减去倒数第二长的。就是最后一个单词的长度
//		if len(arrs) == 1 {
//			length = len(s)
//		} else {
//			length = len(arrs[len(arrs)-1]) - len(arrs[len(arrs)-2])
//		}
//		return length
//	}
func main() {
	s := "World"
	result := lengthOfLastWord1(s)
	fmt.Println(result)
}
