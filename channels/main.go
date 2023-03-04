package main

import (
	"fmt"
)

func main() {
	fmt.Println()

	c := make(chan int) 

	arr := [5]int {2, 4, 6, 8, 10}

	go number(c, arr)

	for {
        num, ok := <-c
        if !ok {
            break
        } else {
            fmt.Println(num, ok)
        }
    }

	fmt.Println()
}

func number (c chan int, arr [5]int) {
	for _, i := range arr {
		c <- i * i
	}
	close(c)
}