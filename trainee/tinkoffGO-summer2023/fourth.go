package tinkoffGO_summer2023

//Набор чисел ﻿x_1, x_2, x_k назовем "скучным", если возможно удалить из него один элемент так,
//чтобы каждое число в данном наборе встречалось одинаковое количество раз.
//Дан массив a_1, a_2,…,an длины n. Найдите максимальное число l(2≤l≤n), что префикс длины l является скучным.
//Формат входных данных:
//Первая строка содержит число n (2≤n≤2⋅10^5) - размер массива.
//Во второй строке находятся n чисел из массива a_1, a_2,…,an (1≤a<=i≤2⋅10^5)
//Формат выходных данных
//Выведите одно число — максимальное l, что префикс длины l массива a является скучным.

//func Fourth() {
//	in := bufio.NewReader(os.Stdin)
//	out := bufio.NewWriter(os.Stdout)
//	defer out.Flush()
//	temp, _, _ := in.ReadLine()
//	total, _ := strconv.Atoi(string(temp))
//	sl := readIntSlice(in)
//	for i := total - 1; i >= 1; i-- {
//		if isBad(sl) {
//			out.WriteString(strconv.Itoa(len(sl)))
//			break
//		} else {
//			sl = sl[:i-1]
//		}
//	}
//}
//
//func isBad(sl []int) bool {
//	m := make(map[int]int)
//	for _, val := range sl {
//		m[val]++
//	}
//	last := m[sl[0]]
//	val := 0
//	count := 0
//	for _, value := range m {
//		if last != value {
//			if count >= 1 {
//				return false
//			}
//			count++
//			val = value
//		} else {
//			last = value
//		}
//	}
//	if count == 1 {
//		if val == 1 {
//			return true
//		}
//		for _, el := range m {
//			if val-el == 1 {
//				return true
//			}
//		}
//	}
//	if count == 0 {
//		for _, el := range m {
//			if el == 1 {
//				return true
//			}
//		}
//	}
//	return false
//}
