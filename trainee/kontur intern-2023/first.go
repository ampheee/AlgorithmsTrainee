package main

//
//import (
//	"bufio"
//	"os"
//	"strconv"
//	"strings"
//)
//
//func main() {
//	first()
//}
//
//func first() {
//	in := bufio.NewReader(os.Stdin)
//	out := bufio.NewWriter(os.Stdout)
//	defer out.Flush()
//	temp, _, _ := in.ReadLine()
//	total, _ := strconv.Atoi(string(temp))
//	sl := readSliceInt(in)
//	var max, min, left, right = sl[0], sl[0], 0, 0
//	for i := 0; i < total; i++ {
//		if sl[i] >= max {
//			max = sl[i]
//			right = i
//		}
//		if sl[i] < min {
//			min = sl[i]
//			left = i
//		}
//	}
//	out.WriteString(strconv.Itoa(right+1) + " " + strconv.Itoa(left+1))
//}
