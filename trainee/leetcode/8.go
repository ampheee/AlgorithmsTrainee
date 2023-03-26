package leetcode

//func myAtoi(s string) int {
//	i := 0
//	for i < len(s) && s[i] == ' ' {
//		i++
//	}
//	sign := 1
//	if i < len(s) && (s[i] == '+' || s[i] == '-') {
//		if s[i] == '-' {
//			sign = -1
//		}
//		i++
//	}
//	num := 0
//	for i < len(s) && unicode.IsDigit(rune(s[i])) && (num > -1<<31 && num < 1<<31-1) {
//		num = num*10 + int(s[i]-'0')
//		i++
//	}
//	num *= sign
//	if num < -1<<31 {
//		return -1 << 31
//	} else if num > 1<<31-1 {
//		return 1<<31 - 1
//	} else {
//		return num
//	}
//}
