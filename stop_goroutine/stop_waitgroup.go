package main

import (
	"fmt"
	"sync"
)

func stopWautGroup() {

	var wg sync.WaitGroup

	ch := make(chan int)

	go function1(ch, &wg)

	ch <- 1

	wg.Wait() // Завершается группа и с ней горутина
}

func function1(ch chan int, wg *sync.WaitGroup) {
	wg.Add(1) // Добавляем в группу
	fmt.Println(<-ch)
	wg.Done() // Минус 1 из группы
}
