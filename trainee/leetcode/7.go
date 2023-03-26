package leetcode

//Given a signed 32-bit integer x, return x with its digits reversed.
//If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

func reverse(x int) int {
	var res, neg int
	if x < 0 {
		x, neg = -x, 1
	}
	for x > 0 {
		res *= 10
		res += x % 10
		x /= 10
	}
	if neg == 1 && -res > -1<<31 {
		return -res
	} else if res > 0 && res < 1<<31-1 {
		return res
	}
	return 0
}
