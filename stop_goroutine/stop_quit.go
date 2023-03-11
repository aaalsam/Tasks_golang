package main

import (
	"fmt"
)

func stopQuit(){

	quit := make(chan bool) // Пустой доп. канал

	ch := make(chan int)

	go function2(ch, quit)

	ch <- 1

	quit <- true // Говорим, что он тру, при этом выполняется кейс
}

func function2(ch chan int, quit chan bool) {
	for {
		select {
		case <-quit:
			return
		case a := <-ch:
			fmt.Println(a)
		}
	}
}
