package leetcode

//You are given a large integer represented as an integer array digits, where each digits[i]
//is the ith digit of the integer.
//The digits are ordered from most significant to least significant in left-to-right order.
//The large integer does not contain any leading 0's.
//
//Increment the large integer by one and return the resulting array of digits.

//func plusOne(digits []int) []int {
//	if digits[len(digits)-1] == 9 {
//		var temp = 1
//		for i := len(digits) - 1; i >= 0; i-- {
//			if digits[i] == 9 && temp == 1 {
//				digits[i] = 0
//				temp = 1
//			} else if digits[i] != 9 {
//				digits[i] += 1
//				temp = 0
//				break
//			}
//		}
//		if temp != 0 {
//			var res = []int{1}
//			return append(res, digits...)
//		}
//	} else {
//		digits[len(digits)-1]++
//	}
//	return digits
//}
