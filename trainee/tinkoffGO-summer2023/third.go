package tinkoffGO_summer2023

//Дана строка s, состоящая из символов a, b, c, d.
//Подстрокой строки s называется строка, которую можно получить, взяв из s какие-то подряд идущие символы.
//Например, строки bcd, abcdcdab, cdcdab являются подстроками строки abcdcdab,
//а cc и cdcdcd — нет.
//Назовем строку хорошей, если каждый из символов a, b, c, d встречается в ней хотя бы один раз.
//Найдите длину самой короткой хорошей подстроки строки s или определите, что у s нет хороших подстрок.
//Формат входных данных
//Первая строка входных данных содержит число n (1≤n≤2⋅10^5) — длину строки.
//Во второй строке находится сама строка s, состоящую из символов a, b, c и d.
//Формат выходных данных
//Выведите длину самой короткой хорошей подстроки строки s. Если хороших подстрок нет, выведите 1.

//func Third() {
//	var str string
//	var ln, res, minIndex = 0, -1, -1
//	fmt.Scan(&ln, &str)
//	var aIndex, bIndex, cIndex, dIndex = -1, -1, -1, -1
//	for r, v := range str {
//		if v == 'a' {
//			aIndex = r
//		}
//		if v == 'b' {
//			bIndex = r
//		}
//		if v == 'c' {
//			cIndex = r
//		}
//		if v == 'd' {
//			dIndex = r
//		}
//		minIndex = int(math.Min(math.Min(float64(aIndex), float64(bIndex)), math.Min(float64(cIndex), float64(dIndex))))
//		if minIndex == -1 {
//			continue
//		}
//		temp := r - minIndex + 1
//		if res == -1 || res > temp {
//			res = temp
//		}
//	}
//	fmt.Println(res)
//}
