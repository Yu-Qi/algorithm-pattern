package algo

import "fmt"

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	need := map[byte]int{}
	for i := 0; i < len(t); i++ {
		need[t[i]] = need[t[i]] + 1
	}
	i := 0
	j := 0
	matchCount := 0
	result := ""
	for j < len(s) {
		// for match the condition
		for j < len(s) {
			c := s[j]
			if num, ok := count[c]; ok {
				if num == 0 {
					matchCount++
				}
				count[c] = num + 1
			}
			if matchCount != len(t) {
				j++
			} else {
				break
			}
		}
		if matchCount != len(t) {
			fmt.Println(i, j, matchCount, count)
			break
		}
	left:
		for i < len(s) {
			result = getMinSubstring(result, s, i, j+1)
			c := s[i]
			i++
			if num, ok := count[c]; ok {
				count[c] = num - 1
				if num == 1 {
					matchCount--
					break left
				}
			}
		}
		j++
	}
	return result
}

func getMinSubstring(result, newString string, startIdx, endIdx int) string {
	if len(newString) < endIdx {
		endIdx = len(newString)
	}
	newString = newString[startIdx:endIdx]

	if result != "" {
		if len(result) < len(newString) {
			return result
		} else {
			return newString
		}
	}
	return newString
}
