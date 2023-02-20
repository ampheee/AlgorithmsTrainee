package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func seven() {
	var A, B, C string
	fmt.Scan(&A, &B, &C)
	aSeconds := convertToMS(A)
	bSeconds := convertToMS(B)
	cSeconds := convertToMS(C)
	if cSeconds < aSeconds {
		cSeconds += 24 * 3600
	}
	res := convertToTime((bSeconds + int(math.Round(float64(cSeconds-aSeconds)/2))) % (24 * 3600))
	fmt.Printf("%.2d:%.2d:%.2d", res[0], res[1], res[2])
}

func convertToMS(str string) int {
	sl := strings.Split(str, ":")
	var res int
	hours, _ := strconv.Atoi(sl[0])
	minutes, _ := strconv.Atoi(sl[1])
	seconds, _ := strconv.Atoi(sl[2])
	res += 3600*hours + 60*minutes + seconds
	return res
}

func convertToTime(bRes int) [3]int {
	var res [3]int
	res[0] = bRes / 3600
	res[1] = (bRes % 3600) / 60
	res[2] = (bRes % 3600) % 60
	return res
}
