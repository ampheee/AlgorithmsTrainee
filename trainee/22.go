package main

//func twentytwo() {
//	in := bufio.NewReader(os.Stdin)
//	out := bufio.NewWriter(os.Stdout)
//	defer out.Flush()
//	temp := readSliceInt(in)
//	var res = make([]int, int(math.Max(float64(temp[0]), 2)))
//	res[0] = 1
//	res[1] = 1
//	for i := 2; i < temp[0]; i++ {
//		for j := 1; j <= temp[1]; j++ {
//			res[i] += res[i-j]
//			if i-j == 0 {
//				break
//			}
//		}
//	}
//	fmt.Println(res[temp[0]-1])
//}

//func readSliceInt(in *bufio.Reader) []int {
//	END := true
//	m := make([]byte, 0)
//	for END {
//		byteSlice, err, _ := in.ReadLine()
//		END = err
//		m = append(m, byteSlice...)
//	}
//	line := string(m)
//	ints := strings.Fields(line)
//	res := make([]int, 0, len(ints)/2)
//	for i := 0; i < len(ints); i++ {
//		temp, _ := strconv.Atoi(ints[i])
//		res = append(res, temp)
//	}
//	return res
//}
