package main

import "fmt"

// intersect возвращает пересечение двух слайсов a и b (без повторов в a)
func intersect(a, b []int) []int {
	set := make(map[int]struct{}) // Перегоняем элементы a в set
	for _, v := range a {
		set[v] = struct{}{}
	}
	var result []int
	for _, v := range b {
		// Перебор сета B
		if _, ok := set[v]; ok { // Если v есть в set -> это пересечение
			result = append(result, v)
		}
	}
	return result
}

func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}
	fmt.Println(intersect(A, B))
}
