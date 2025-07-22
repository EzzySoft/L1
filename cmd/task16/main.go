package main

import "fmt"

// quickSort возвращает новый отсортированный срез
func quickSort(a []int) []int {
	// Условие выхода: массив длины 0 или 1 уже отсортирован
	if len(a) < 2 {
		return a
	}
	pivot := a[0]
	var less, greater []int
	for _, v := range a[1:] {
		// less — всё, что меньше опорного, greater — всё остальное
		if v < pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}
	// Рекурсивно сортируем обе части и объединяем с опорным элементом по центру
	return append(append(quickSort(less), pivot), quickSort(greater)...)
}

func main() {
	a := []int{5, 3, 8, 4, 2, 7, 1, 6}
	fmt.Println(quickSort(a))
}
