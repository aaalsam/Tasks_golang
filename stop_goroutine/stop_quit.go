package main

import (
	"fmt"
	"time"
)

// принято создавать канал пустых структур

func stopQuit() {

	quit := make(chan int)

	go function2(quit)

	quit <- 1 // Говорим, что он тру, при этом выполняется кейс
	//close(quit)

	time.Sleep(1 * time.Second)
}

func function2(quit chan int) {
	for {
		select {
		case <-quit:
			fmt.Println("мы тут")
			return
		default:
			fmt.Println("0")
		}
	}
}
