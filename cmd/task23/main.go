package main

import "fmt"

func removeAt(slice []int, i int) []int {
	// Сдвигаем хвост на место i-го элемента
	copy(slice[i:], slice[i+1:])
	// Отрезаем лишние элементы (последний элемент теперь дублируется)
	slice = slice[:len(slice)-1]
	return slice
}

func main() {
	a := []int{10, 20, 30, 40, 50}
	i := 2 // удаляем элемент с индексом 2 (30)
	a = removeAt(a, i)
	fmt.Println(a) // [10 20 40 50]
}
