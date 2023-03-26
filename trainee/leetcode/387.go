package leetcode

//Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.
//
//type test struct {
//	count int
//	index int
//}
//
//func firstUniqChar(s string) int {
//	u := make(map[int32]*test)
//	for i, val := range s {
//		if _, exist := u[val]; exist {
//			u[val].index = i
//			u[val].count++
//		} else {
//			u[val] = &test{count: 1, index: i}
//		}
//	}
//	var res = len(s)
//	for _, val := range u {
//		if val.count == 1 && val.index <= res {
//			res = val.index
//		}
//	}
//	if res != len(s) {
//		return res
//	}
//	return -1
//}
