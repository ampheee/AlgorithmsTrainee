package main

//func first() {
//	var file, _ = os.Open("input.txt")
//	scanner := bufio.NewReader(file)
//	m := make(map[string]int)
//	for {
//		var char rune
//		var err error
//		if char, _, err = scanner.ReadRune(); err == io.EOF {
//			break
//		}
//		if _, exist := m[string(char)]; exist && char != ' ' {
//			m[string(char)]++
//		} else if char != ' ' && char != '\n' {
//			m[string(char)] = 1
//		}
//	}
//	keys := getKeys(m)
//	max := getMax(m)
//	sort.Slice(keys, func(i, j int) bool {
//		return keys[j] > keys[i]
//	})
//	for ; max > 0; max-- {
//		for i := 0; i < len(keys); i++ {
//			if m[keys[i]] == max {
//				fmt.Print("#")
//				m[keys[i]]--
//				continue
//			}
//			fmt.Print(" ")
//		}
//		fmt.Println()
//	}
//	fmt.Println(strings.Join(keys, ""))
//}
//
//func getKeys(m map[string]int) []string {
//	keys := make([]string, 0, len(m))
//	for i := range m {
//		keys = append(keys, i)
//	}
//	return keys
//}

//func getMax(arr [150]int) int {
//	var max int
//	for _, val := range arr {
//		if val > max {
//			max = val
//		}
//	}
//	return max
//}
//
//func firstMy() {
//	//var file, _ = os.Open("input.txt")
//	scanner := bufio.NewScanner(os.Stdin)
//	var begin = 33
//	var arr [150]int
//	for scanner.Scan() {
//		for _, val := range scanner.Text() {
//			if val >= rune(begin) {
//				arr[val-rune(begin)]++
//			}
//		}
//	}
//	max := getMax(arr)
//	for i := 0; i < max; i++ {
//		for j := 0; j < 150; j++ {
//			if arr[j] != 0 && arr[j] < max-i {
//				fmt.Print(" ")
//			} else if arr[j] != 0 {
//				fmt.Print("#")
//			}
//		}
//		fmt.Println()
//	}
//	for i := 0; i < 150; i++ {
//		if arr[i] != 0 {
//			fmt.Printf("%c", i+begin)
//		}
//	}
//}
