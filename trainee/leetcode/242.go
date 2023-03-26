package leetcode

//Given two strings s and t, return true if t is an anagram of s, and false otherwise.
//An Anagram is a word or phrase formed by rearranging the
//letters of a different word or phrase, typically using all the original letters exactly once.
//
//func isAnagram(s string, t string) bool {
//	u := make(map[rune]int)
//	for _, val := range s {
//		u[val]++
//	}
//	for _, val := range t {
//		if getted, exist := u[val]; exist && getted > 0 {
//			u[val]--
//		} else {
//			return false
//		}
//	}
//	for _, val := range u {
//		if val != 0 {
//			return false
//		}
//	}
//	return true
//}
