package main

import "fmt"

// binarySearch ищет target в отсортированном слайсе a и возвращает индекс или -1, если не найден
func binarySearch(a []int, target int) int {
	left, right := 0, len(a)-1
	for left <= right {
		mid := left + (right-left)/2 // чтобы избежать переполнения
		if a[mid] == target {
			return mid
		} else if a[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1 // не найден
}

func main() {
	arr := []int{1, 3, 4, 6, 8, 10, 12}
	fmt.Println(binarySearch(arr, 6)) // 3
	fmt.Println(binarySearch(arr, 7)) // -1

}
