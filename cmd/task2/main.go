package main

import (
	"fmt"
	"sync"
)

const N = 5

// Массив чисел для обработки
var nums = [N]int{2, 4, 6, 8, 10}

// Вычисляет квадрат числа и выводит результат
// После работы сообщает WaitGroup что задача завершена
func printSquare(num int, wg *sync.WaitGroup) {
	defer wg.Done() // Сообщаем о завершении даже при ошибках
	fmt.Println(num * num)
}

func main() {
	var wg sync.WaitGroup // Счётчик для синхронизации всех горутин

	// Для каждого числа запускаем отдельную горутину
	for _, n := range nums {
		wg.Add(1)              // Увеличиваем счётчик и ожидаем новую задачу
		go printSquare(n, &wg) // Запускаем горутину для вычисления квадрата
	}

	wg.Wait() // Ждём завершения всех горутин
}
