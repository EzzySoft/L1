package main

import (
	"fmt"
	"strings"
)

func isUnique(s string) bool {
	s = strings.ToLower(s)          // игнорируем регистр
	seen := make(map[rune]struct{}) // set для рун
	for _, r := range s {
		if _, ok := seen[r]; ok {
			return false // уже встречался
		}
		seen[r] = struct{}{}
	}
	return true
}

func main() {
	fmt.Println(isUnique("abcd"))  // true
	fmt.Println(isUnique("abcdD")) // false

}
