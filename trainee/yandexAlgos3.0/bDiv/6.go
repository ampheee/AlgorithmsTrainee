package main

import "fmt"

func sixth() {
	var units, taken int
	fmt.Scan(&units)
	fmt.Scan(&taken)
	m := make(map[[2]int]int)
	for i := 0; i < taken; i++ {
		var tempSlice [2]int
		fmt.Scan(&tempSlice[0], &tempSlice[1])
		if i == 0 {
			m[tempSlice] = 1
		} else {
			for k, _ := range m {
				if (tempSlice[0] <= k[0] && tempSlice[1] >= k[1]) ||
					(tempSlice[1] <= k[1] && tempSlice[1] >= k[0]) ||
					(tempSlice[0] >= k[0] && tempSlice[0] <= k[1]) {
					m[k] = 0
				}
			}
			m[tempSlice] = 1
		}
	}
	var res int
	for _, val := range m {
		res += val
	}
	fmt.Println(res)
}
