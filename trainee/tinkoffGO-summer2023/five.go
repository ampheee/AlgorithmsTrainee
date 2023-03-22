package tinkoffGO_summer2023

//
//func third() {
//	in := bufio.NewReader(os.Stdin)
//	out := bufio.NewWriter(os.Stdout)
//	defer out.Flush()
//}
//
//func fourth() {
//	in := bufio.NewReader(os.Stdin)
//	out := bufio.NewWriter(os.Stdout)
//	defer out.Flush()
//}

//Саша ведет бюджет и анализирует, как изменялся баланс на его счету. Он выписал числа a1, a2, ..., an  — изменения баланса за последние n дней
//Отрезок из дней [i,j] Саша считает разумным, если суммарное изменение баланса с i-го по j-й день равно нулю, т.е. a_i + a_{i+1} + ... + a_j=0.
//Отрезок из дней [l,r] считается нормальным, если внутри данного отрезка можно найти разумный подотрезок.
//Помогите Саше проанализировать эффективность ведения бюджета и посчитайте количество нормальных отрезков в массиве изменений баланса его счета.
//Формат входных данных:
//Первая строка содержит число n (1 <= n <= 2 * 10^5) — количество дней, для которых Саша записывал изменения баланса.
//Вторая строка содержит 10^9 <= n <= 10^9 \leq 10^9) — значения изменений баланса.
//Формат выходных данных:
//Выведите одно число — количество нормальных отрезков для данного массива.

//import (
//	"bufio"
//	"strconv"
//	"strings"
//)
//
//func Five() {
//	in := bufio.NewReader(os.Stdin)
//	out := bufio.NewWriter(os.Stdout)
//	defer out.Flush()
//	temp, _, _ := in.ReadLine()
//	total, _ := strconv.Atoi(string(temp))
//	sl := readPrefixIntSlice(in)
//	var res = 0
//	for i := 0; i < total+1; i++ {
//		for j := i + 1; j < total+1; j++ {
//			if sl[j] == sl[i] {
//				res += total - j + 1
//			}
//		}
//	}
//	out.WriteString(strconv.Itoa(res))
//}
//
//func readPrefixIntSlice(in *bufio.Reader) []int {
//	var end = true
//	byteStr := make([]byte, 0)
//	for end {
//		temp, err, _ := in.ReadLine()
//		byteStr = append(byteStr, temp...)
//		end = err
//	}
//	finalStr := strings.Fields(string(byteStr))
//	res := make([]int, 0, len(finalStr)/2+1)
//	res = append(res, 0)
//	for i := 1; i <= len(finalStr); i++ {
//		val, _ := strconv.Atoi(finalStr[i-1])
//		res = append(res, res[i-1]+val)
//	}
//	return res
//}
