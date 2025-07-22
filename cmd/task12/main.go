package main

import "fmt"

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})

	// как и в прошлый раз используем set для O(1) вставок уникальных значений
	for _, w := range words {
		set[w] = struct{}{}
	}

	// Выводим результат как множество
	fmt.Print("{")
	first := true
	for w := range set {
		if !first {
			fmt.Print(", ")
		}
		fmt.Print(w)
		first = false
	}
	fmt.Println("}")
}
