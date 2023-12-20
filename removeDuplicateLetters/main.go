package main

import "fmt"

func main() {

	result := removeDuplicateLetters("cbacdcbc")

	fmt.Println(result)
}

func removeDuplicateLetters(s string) string {
	freqMap := make(map[uint8]int)
	resulted := make(map[uint8]bool)
	result := ""

	for i := 0; i < len(s); i++ {
		if freqMap[s[i]] == 0 {
			freqMap[s[i]] = 1
		} else {
			freqMap[s[i]]++
		}

	}

	for i := 0; i < len(s); i++ {
		if freqMap[s[i]] >= 1 {
			freqMap[s[i]]--
		}

		if resulted[s[i]] == false {
			for len(result) > 0 && result[len(result)-1] > s[i] && freqMap[result[len(result)-1]] > 0 {
				delete(resulted, result[len(result)-1])
				result = result[:len(result)-1]
			}

			result += string(s[i])
			resulted[s[i]] = true
		}

	}

	return result
}
