package main

import "fmt"

func fifth() {
	var (
		count   = 0
		longest = 0
		arr     []int
	)
	fmt.Scan(&count)
	for i := 0; i < count; i++ {
		var temp int
		fmt.Scan(&temp)
		arr = append(arr, temp)
	}
	for i := 0; i < count-1; i++ {
		longest += getMin(arr[i], arr[i+1])
	}
	fmt.Print(longest)
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
