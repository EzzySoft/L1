package main

import (
	"fmt"
	"time"
)

// MySleep блокирует выполнение на duration
func MySleep(duration time.Duration) {
	done := make(chan struct{})
	go func() {
		time.AfterFunc(duration, func() {
			close(done) // Через duration закрываем канал
		})
	}()
	<-done // Блокируемся до закрытия канала
}

func main() {
	fmt.Println("Start")
	MySleep(2 * time.Second)
	fmt.Println("After 2 seconds")
}
