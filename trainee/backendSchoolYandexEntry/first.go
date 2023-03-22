package backendSchoolYandexEntry

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func first() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	total := readIntSlice(in)
	for i := 0; i < total[2]; i++ {
		temp, _, _ := in.ReadLine()
		str := strings.Fields(string(temp))
		switch str[0] {
		case "RESET":
		case "DISABLE":
		case "GETMAX":
		case "GETMIN":

		}
	}

}

func readIntSlice(in *bufio.Reader) []int {
	var end = true
	byteStr := make([]byte, 0)
	for end {
		temp, err, _ := in.ReadLine()
		byteStr = append(byteStr, temp...)
		end = err
	}
	finalStr := strings.Fields(string(byteStr))
	res := make([]int, 0, len(finalStr)/2)
	for i := 0; i < len(finalStr); i++ {
		val, _ := strconv.Atoi(finalStr[i])
		res = append(res, val)
	}
	return res
}
