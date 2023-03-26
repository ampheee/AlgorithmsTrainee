package leetcode

//
//A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and
//removing all non-alphanumeric characters, it reads the same forward and backward.
//	Alphanumeric characters include letters and numbers.
//Given a string s, return true if it is a palindrome, or false otherwise.
//
//func isPalindrome(s string) bool {
//	if len(s) == 1 {
//		return true
//	}
//	var ptr1, ptr2 = 0, len(s) - 1
//	for ptr1 < ptr2 {
//		r1, r2 := strings.ToLower(string(s[ptr1])), strings.ToLower(string(s[ptr2]))
//		if !isValidChar(r1[0]) {
//			ptr1++
//		} else if !isValidChar(r2[0]) {
//			ptr2--
//		} else if r1[0] != r2[0] {
//			return false
//		} else {
//			ptr1++
//			ptr2--
//		}
//
//	}
//	return true
//}
//
//func isValidChar(r uint8) bool {
//	if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
//		return true
//	}
//	return false
//}
