package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	m := make(map[int]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	// Стартуем 10 горутин, каждая пишет 100 значений
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				mu.Lock() // Защищаем map на запись и чтение
				m[id*100+j] = id
				mu.Unlock()
				time.Sleep(time.Microsecond)
			}
		}(i)
	}

	wg.Wait()

	// Проверка: сколько всего записей
	mu.Lock()
	fmt.Println("Всего записей в map:", len(m))
	mu.Unlock()
}
