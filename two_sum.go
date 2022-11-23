package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
func twoSum(nums []int, target int) []int {
	var result []int
	// for i := 0; i < len(nums) - 1; i++ {
	//     for j := i+1; j < len(nums); j++ {
	//         if nums[i] + nums[j] == target {
	//         result = append(result, i)
	//         result = append(result, j)
	//     }
	//     }
	// }
	// return result
	res := make(map[int]int)

	for i, v := range nums {
		if val, ok := res[target-v]; ok {
			result = append(result, i)
			result = append(result, val)
		}
		res[v] = i
	}
	return result
}
