package leetcode

import "strings"

func longestCommonPrefix(strs []string) string {
	var res = ""
	if len(strs) != 0 {
		prefix := strs[0]
		for i := 1; i < len(strs); i++ {
			for !strings.HasPrefix(strs[i], prefix) {
				prefix = prefix[:len(prefix)-1]
			}
		}
		return prefix
	}
	return res
}
