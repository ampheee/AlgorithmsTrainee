package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func twentyone() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	b, _, _ := in.ReadLine()
	total, _ := strconv.Atoi(string(b))
	var temp = make([]int, int(math.Max(float64(total), 3)))
	temp[0] = 2
	temp[1] = 4
	temp[2] = 7
	for i := 3; i < total; i++ {
		temp[i] = temp[i-3] + temp[i-2] + temp[i-1]
	}
	fmt.Println(temp[total-1])
}
