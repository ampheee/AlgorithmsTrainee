package leetcode

//Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
//You must implement a solution with a linear runtime complexity and use only constant extra space.

//func singleNumber(nums []int) int {
//	n := len(nums)
//	if n == 1 {
//		return nums[0]
//	}
//	var xor int
//	for _, val := range nums {
//		xor ^= val
//	}
//	return xor
//}
