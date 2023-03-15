package main

//func twentyFive() {
//	in := bufio.NewReader(os.Stdin)
//	out := bufio.NewWriter(os.Stdout)
//	defer out.Flush()
//	temp, _, _ := in.ReadLine()
//	total, _ := strconv.Atoi(string(temp))
//	sl := readSliceInt(in)
//	sort.Slice(sl, func(i, j int) bool {
//		return sl[i] < sl[j]
//	})
//	m := make([]int, total)
//	out.WriteString(strconv.Itoa(sl[len(sl)-1] - sl[0] - int(math.Max(float64(max(1, total, &m, sl)),
//		float64(max(2, total, &m, sl))))))
//}
//
//func max(i, total int, temp *[]int, in []int) int {
//	if i > total-3 {
//		return 0
//	}
//	if (*temp)[i] == 0 {
//		(*temp)[i] = in[i+1] - in[i] + int(math.Max(float64(max(i+2, total, temp, in)), float64(max(i+3, total, temp, in))))
//	}
//	return (*temp)[i]
//}
