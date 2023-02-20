package main

//
//func three() {
//	start := time.Now()
//	var count int64
//	var uniqs []int64
//	fmt.Scan(&count)
//	m := make(map[int64]int64, 100001)
//	for i := int64(0); i < count; i++ {
//		var temp int64
//		fmt.Scan(&temp)
//		m[temp] = 1
//	}
//	uniqs = getUniqs(m)
//	sort.Slice(uniqs, func(i, j int) bool {
//		return uniqs[i] < uniqs[j]
//	})
//	var collectioners int
//	fmt.Scan(&collectioners)
//	for i := 0; i < collectioners; i++ {
//		var temp int64
//		fmt.Scan(&temp)
//		index := binarySearch(0, int64(len(uniqs)-1), int64(temp), uniqs)
//		if uniqs[index] < temp {
//			fmt.Printf("%d\n", index+1)
//		} else {
//			fmt.Printf("0\n")
//		}
//	}
//	end := time.Since(start)
//	fmt.Println(end.Seconds())
//}
//
//func AppendIfMissing(slice []int64, i int64) []int64 {
//	for _, ele := range slice {
//		if ele == i {
//			return slice
//		}
//	}
//	return append(slice, i)
//}
//
//func binarySearch(l, r, val int64, slice []int64) int64 {
//	for l < r {
//		m := (l + r + 1) / 2
//		if slice[m] < val {
//			l = m
//		} else {
//			r = m - 1
//		}
//	}
//	return l
//}
//
//func getUniqs(m map[int64]int64) (uniqs []int64) {
//	for k, _ := range m {
//		uniqs = append(uniqs, k)
//	}
//	return
//}
