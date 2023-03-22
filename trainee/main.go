package main

import (
	"bufio"
	"strconv"
	"strings"
)

func readPrefixIntSlice(in *bufio.Reader) []int {
	var end = true
	byteStr := make([]byte, 0)
	for end {
		temp, err, _ := in.ReadLine()
		byteStr = append(byteStr, temp...)
		end = err
	}
	finalStr := strings.Fields(string(byteStr))
	res := make([]int, 0, len(finalStr)/2+1)
	res = append(res, 0)
	for i := 1; i <= len(finalStr); i++ {
		val, _ := strconv.Atoi(finalStr[i-1])
		res = append(res, res[i-1]+val)
	}
	return res
}
