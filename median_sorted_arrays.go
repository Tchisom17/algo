package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)

	if len(nums1)%2 != 0 {
		return float64(nums1[len(nums1)/2])
	}

	return float64(nums1[(len(nums1)-1)/2]+nums1[len(nums1)/2]) / 2.0
}
