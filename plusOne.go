package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func plusOne1(digits []int) []int {
	var length int
	length = len(digits)
	if digits[length-1] < 9 || length == 1 { //如果最后一位小于9或者长度等于1
		digits[length-1] += 1
	} else {
		for i := length - 1; i > 0; i-- {
			if digits[i] >= 9 {
				digits[i] = 0
				digits[i-1] += 1
				if digits[i-1] <= 9 { //如果前一位加一后还是小于10，就停止循环，避免类似于8989加1后一直进位成9000
					break
				}
			}
		}
	}
	if digits[0] == 10 { //判断第一位，如果是10，就进位
		digits = append([]int{1}, digits...)
		digits[1] = 0
	}
	return digits
}

func plusOne(digits []int) []int {
	var s string
	var num big.Int //转为大数
	var arrs []int
	nums := new(big.Int)
	s = ""
	for _, v := range digits {
		s = s + strconv.Itoa(v)
	}
	nums, _ = num.SetString(s, 10)
	one := big.NewInt(1)
	nums.Add(nums, one) //大数的加法
	s = nums.String()   //将大数转为string类型的

	for _, v := range s {
		v2 := string(v)
		v1, _ := strconv.Atoi(v2)
		arrs = append(arrs, v1)
		//fmt.Println(v1)
	}
	return arrs
}
func main() {
	digits := []int{8, 9, 8, 9}
	result := plusOne1(digits)
	fmt.Println(result)
}
