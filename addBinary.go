package main

import (
	"fmt"
	"math"
	"strconv"
)

func addBinary(a string, b string) string {
	var l, la, lb int //字符的长度，以及需要补的位数
	var zero string   //补位用的0
	la = len(a)
	lb = len(b)
	l = int(math.Abs(float64(la - lb))) //需要补零的个数
	var num1, num2 int                  //二进制数每位的数
	var arr []int                       //存入数组中
	var result string                   //返回的结果
	//补位，短的字符串前面补0，两段一样长
	for i := 0; i < l; i++ {
		zero += "0"
	}
	if la > lb {
		b = zero + b
	} else {
		a = zero + a
	}

	//两端逐位相加
	for i := 0; i < len(a); i++ {
		num1, _ = strconv.Atoi(string(a[i]))
		num2, _ = strconv.Atoi(string(b[i]))
		arr = append(arr, num1+num2)
	}

	//处理进位
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] > 1 {
			arr[i] -= 2
			arr[i-1] += 1
		}
	}

	//处理第一位，如果大于1的话，位数加一
	var one = []int{1}
	if arr[0] >= 2 {
		arr[0] -= 2
		arr = append(one, arr...)
	}

	//转为字符串返回
	for _, v := range arr {
		result += strconv.Itoa(v)
	}
	return result
}
func main() {
	a := "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101"
	b := "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"
	s := addBinary(a, b)
	fmt.Println(s)
}
