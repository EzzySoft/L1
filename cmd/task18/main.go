package main

import (
	"fmt"
	"sync"
)

// Counter — потокобезопасный счётчик
// Использует sync.Mutex для защиты от гонок данных
type Counter struct {
	mu sync.Mutex
	n  int
}

// Inc увеличивает значение счётчика на 1 с захватом мьютекса
func (c *Counter) Inc() {
	c.mu.Lock()
	c.n++
	c.mu.Unlock()
}

// Value возвращает текущее значение счётчика с захватом мьютекса
func (c *Counter) Value() int {
	c.mu.Lock()
	v := c.n
	c.mu.Unlock()
	return v
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	// Запускаем 1000 горутин — каждая увеличивает счётчик на 1
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter.Inc()
			wg.Done()
		}()
	}

	wg.Wait() // Ожидаем завершения всех горутин
	fmt.Println("Итоговое значение:", counter.Value())
}
