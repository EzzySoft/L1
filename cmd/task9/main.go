package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	ch1 := make(chan int)
	ch2 := make(chan int)

	//  Генерация чисел — пишем x в ch1
	go func() {
		for _, x := range numbers {
			ch1 <- x
		}
		close(ch1) // закрываем
	}()

	//  Обработка — читаем x из ch1, пишем x*2 в ch2
	go func() {
		for x := range ch1 {
			ch2 <- x * 2
		}
		close(ch2) // обязательно закрыть!
	}()

	// Чтение из ch2 и вывод в stdout
	for y := range ch2 {
		fmt.Println(y)
	}
}
