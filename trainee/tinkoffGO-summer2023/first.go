package tinkoffGO_summer2023

//Четыре человека выстроились в очередь друг за другом. Определите, правда ли, что они стоят по росту (неважно, в порядке неубывания или порядке невозрастания).
//Формат входных данных
//Единственная строка содержит числа h_1 h_2, ..., h_n(0≤h_n
// ≤300), где h_n рост человека, стоящего на i-й позиции.
//Формат выходных данных
//Выведите YES, если люди выстроены в порядке неубывания или невозрастания их роста, и NO в противном случае.

//import (
//	"bufio"
//	"os"
//)
//
//func first() {
//	in := bufio.NewReader(os.Stdin)
//	out := bufio.NewWriter(os.Stdout)
//	defer out.Flush()
//	sl := readIntSlice(in)
//	var toHigher, toLower = true, true
//	var previous int
//	for i, val := range sl {
//		if i == 0 {
//			previous = val
//			continue
//		}
//		if val > previous && toLower {
//			toLower = false
//		} else if val < previous && toHigher {
//			toHigher = false
//		}
//		previous = val
//	}
//	if toLower || toHigher {
//		out.WriteString("YES")
//	} else {
//		out.WriteString("NO")
//	}
//}
