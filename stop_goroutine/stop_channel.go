package main

import (
	"fmt"
	"time"
)

func stopChannel() {

	ch := make(chan int)

	go function3(ch)

	ch <- 1

	close(ch) // Закрываем канал

	time.Sleep(1 * time.Millisecond) // Ожидание 1 млс, чтобы горутина успели запуститься

}

func function3(ch chan int) {
	fmt.Println(<-ch)
}
