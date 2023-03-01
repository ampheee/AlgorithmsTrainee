package main

import "fmt"

func nine() {
	var N, M, K int
	fmt.Scan(&N, &M, &K)
	var arr = make([][]int, N)
	for i := 0; i < N; i++ {
		arr[i] = make([]int, M)
		for j := 0; j < M; j++ {
			fmt.Scan(&arr[i][j])
		}
	}
	var prefixes = make([][]int, N)
	for i := 0; i < N; i++ {
		prefixes[i] = make([]int, M)
		for j := 0; j < M; j++ {
			prefixes[i][j] = arr[i][j]
			if i > 0 {
				prefixes[i][j] += prefixes[i-1][j]
			}
			if j > 0 {
				prefixes[i][j] += prefixes[i][j-1]
			}
			if j > 0 && i > 0 {
				prefixes[i][j] -= prefixes[i-1][j-1]
			}
		}
	}
	for z := 0; z < K; z++ {
		var x1, x2, y1, y2 int
		fmt.Scan(&x1, &y1, &x2, &y2)
		x1--
		x2--
		y1--
		y2--
		res := prefixes[x2][y2]
		if x1 > 0 {
			res -= prefixes[x1-1][y2]
		}
		if y1 > 0 {
			res -= prefixes[x2][y1-1]
		}
		if x1 > 0 && y1 > 0 {
			res += prefixes[x1-1][y1-1]
		}
		fmt.Println(res)
	}
}
