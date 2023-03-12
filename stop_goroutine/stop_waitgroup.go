package main

import (
	"fmt"
	"sync"
	"syscall"
	"time"
)

func stopWaitGroup() {

	wg := new(sync.WaitGroup)

	ch := make(chan int)

	go function1(ch, wg)

	syscall.

	time.Sleep(10 * time.Second)

	ch <- 1

	wg.Wait() // Завершается группа и с ней горутина, когда счетчик в группе не станет 0
}

func function1(ch chan int, wg *sync.WaitGroup) {
	wg.Add(1) // Добавляем в группу
	fmt.Println(<-ch)
	wg.Done() // Минус 1 из группы
	
	fmt.Println("мы тут")
}
