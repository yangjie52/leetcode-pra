package main

func main() {

}
func twoSum(nums []int, target int) []int {
	var result []int

	for i, value1 := range nums {
		for j, value2 := range nums {
			if value1+value2 == target {
				result = append(result, i)
				result = append(result, j)
			}
		}
	}
	return result
}
func sum(nums []int, target int) []int {
	result := map[int]int{}
	for i, value := range nums {
		if m, ok := result[target-value]; ok {
			return []int{i, m}
		}
		result[value] = i
	}
	return []int{}
}
