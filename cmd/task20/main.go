package main

import "fmt"

// Разворачивает срез рун между left и right
func reverse(runes []rune, left, right int) {
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
}

// Разворачивает слова в предложении in-place
func reverseWordsInPlace(runes []rune) {
	reverse(runes, 0, len(runes)-1) // Разворачиваем всю строку

	start := 0
	for i := 0; i <= len(runes); i++ {
		// Разделяем по пробелу
		if i == len(runes) || runes[i] == ' ' {
			reverse(runes, start, i-1) // Разворачиваем каждое слово
			start = i + 1
		}
	}
}

func main() {
	s := "sun dog snow"
	runes := []rune(s)
	reverseWordsInPlace(runes)
	fmt.Println(string(runes)) // snow dog sun
}
