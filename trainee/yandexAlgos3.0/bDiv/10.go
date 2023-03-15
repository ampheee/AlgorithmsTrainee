package main

//func ten() {
//	var str string
//	fmt.Scan(&str)
//	size := len(str)
//	lstr, count, rstr := 0, 0, size-1
//	m := make(map[uint8]int)
//	for lstr < rstr {
//		temp := (size + size - count) * (lstr + 1) / 2
//		m[str[lstr]] += temp
//		m[str[rstr]] += temp
//		count += 2
//		lstr++
//		rstr--
//	}
//	if size%2 != 0 {
//		m[str[size/2]] += (size + 1) * (size/2 + 1) / 2
//	}
//	keys := getKeys(m)
//	sort.Slice(keys, func(i, j int) bool {
//		return keys[i] < keys[j]
//	})
//	for i := range keys {
//		fmt.Printf("%c: %d\n", keys[i], m[keys[i]])
//	}
//}
//
//func getKeys(m map[uint8]int) (keys []uint8) {
//	for k := range m {
//		keys = append(keys, k)
//	}
//	return
//}

//func recursy(str, dividing string, arr *[26]int) {
//	var lstr = len(str)
//	if lstr > 1 {
//		var count = 0
//		for i := 0; i < lstr; i++ {
//			for j := 0; j < lstr-count; j++ {
//				arr[str[j]-start]++
//			}
//			count++
//		}
//		recursy(dividing[:len(str)-1], dividing[1:], arr)
//	} else {
//		arr[str[0]-start]++
//		return
//	}
//}
