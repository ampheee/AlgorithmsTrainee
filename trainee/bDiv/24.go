package main

//func twentyFourth() {
//	in := bufio.NewReader(os.Stdin)
//	out := bufio.NewWriter(os.Stdout)
//	defer out.Flush()
//	temp, _, _ := in.ReadLine()
//	total, _ := strconv.Atoi(string(temp))
//	var aSl []int
//	var bSl []int
//	var cSl []int
//	for i := 0; i < total; i++ {
//		sl := readSliceInt(in)
//		aSl = append(aSl, sl[0])
//		bSl = append(bSl, sl[1])
//		cSl = append(cSl, sl[2])
//	}
//	d := make([]int, total+1)
//	d[1] = aSl[0]
//	if total > 1 {
//		d[2] = int(math.Min(float64(aSl[0]+aSl[1]), float64(bSl[0])))
//	}
//	if total > 2 {
//		d[3] = int(math.Min(float64(d[2]+aSl[2]), float64(d[1]+bSl[1])))
//		d[3] = int(math.Min(float64(d[3]), float64(cSl[0])))
//	}
//	for i := 0; i <= total; i++ {
//		if i >= 4 {
//			d[i] = int(math.Min(float64(d[i-1]+aSl[i-1]), float64(d[i-2]+bSl[i-2])))
//			d[i] = int(math.Min(float64(d[i]), float64(d[i-3]+cSl[i-3])))
//		}
//	}
//	out.WriteString(strconv.Itoa(d[total]))
//}
