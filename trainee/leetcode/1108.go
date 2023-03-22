package leetcode

//Given a valid (IPv4) IP address, return a defanged version of that IP address.
//A defanged IP address replaces every period "." with "[.]".

//func defangIPaddr(address string) string {
//	var str = make([]byte, 0)
//	for _, val := range address {
//		if val == '.' {
//			str = append(str, '[', '.', ']')
//			continue
//		}
//		str = append(str, byte(val))
//	}
//	return string(str)
//}
