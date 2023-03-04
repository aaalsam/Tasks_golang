package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

func main() {

	c := make(chan int) 

	arr := [5]int {2, 4, 6, 8, 10}

	mu.Lock()
	defer mu.Unlock()

	go number(c, arr)

	for {
        num, ok := <-c
        if !ok {
            break
        } else {
            fmt.Println(num, ok)
        }
    }
}

func number (c chan int, arr [5]int) {
	for _, i := range arr {
		c <- i * i
	}
	close(c)
}