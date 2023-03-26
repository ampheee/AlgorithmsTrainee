package leetcode

//Given two integer arrays nums1 and nums2, return an array of their intersection.
//Each element in the result must appear as many times as it
//shows in both arrays and you may return the result in any order.

//func intersect(nums1 []int, nums2 []int) []int {
//	intersection := make(map[int]int)
//	for _, val := range nums1 {
//		intersection[val]++
//	}
//	var res []int
//	for _, val := range nums2 {
//		if intersection[val] > 0 {
//			res = append(res, val)
//			intersection[val]--
//		}
//	}
//	return res
//}
