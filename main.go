package main

import (
	"fmt"
	"strings"
)

func findSubstrings(s string) int32 {
	result := []string{}
	chars := strings.Split(s, "")
	result = chars

	for k, _ := range chars {
		for i := k; i <= len(chars); i++ {
			sbstr := s[k:i]
			if hasDiferentChars(sbstr) && !contains(result, sbstr) && sbstr != "" {
				result = append(result, s[k:i])
			}
		}
	}
	fmt.Println("substrings encontrados: " + fmt.Sprintf("%v", result))
	return int32(len(result))
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func hasDiferentChars(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}

	return true
}

func main() {
	// fmt.Println(findSubstrings("abac"))
	// fmt.Println(findSubstrings("baba"))
	// fmt.Println(findSubstrings("aa"))
	// fmt.Println(findSubstrings(""))
}
