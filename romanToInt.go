// 罗马数字
package main

import "fmt"

func main() {
	s := "IV"
	a := romanToInt(s)
	fmt.Println(a)
	//romanToInt(s)
}

func romanToInt(s string) int {
	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	var result int
	for i, v := range s {
		if i+1 < len(s) && roman[string(int(s[i]))] < roman[string(int(s[i+1]))] {
			result += roman[string(v)] * (-1)
		} else {
			result = result + roman[string(v)]
		}
	}
	return result
}
