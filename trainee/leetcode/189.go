package leetcode

//Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

//func rotate(nums []int, k int) {
//	n := len(nums)
//	k = k % n
//	if k != 0 {
//		reverse(nums[:n-k])
//		reverse(nums[n-k:])
//		reverse(nums)
//	}
//	fmt.Println(nums)
//}
//
//func reverse(nums []int) {
//	n := len(nums)
//	for i := 0; i < n/2; i++ {
//		j := n - i - 1
//		nums[i], nums[j] = nums[j], nums[i]
//	}
//}
