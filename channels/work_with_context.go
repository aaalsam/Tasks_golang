package main

import (
	"fmt"
)

//Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.

func helloMyNameIsNastya()

// func writer(ctx context.Context, ch chan string) {
// 	for {
// 		select {
// 		case ch <- "message1": // Мы отправляем в канал "message1", "We send message" будет вызываться при чтении!
// 			fmt.Println("We send message")
// 		case <-ctx.Done(): // Мы ожидаем читение из канала, значит будет вызываться We cancel context", когда будем писать в канал
// 			fmt.Println("We cancel context")
// 			return
// 		}
// 	}
// }

func writer2(ch chan string, quit chan struct{}) {
	for {
		select {
		case ch <- "Hello, my name Nastya": // Мы отправляем в канал "message1", будет вызываться при чтении.
			fmt.Println("We send  some message")
		case <-quit: // Мы читаем из канала, значит будет вызываться, когда будем писать в канал
			fmt.Println("Hello")
			close(quit)
			return
		}
	}
}

// func someFunction(x int) int {
// 	if x==5 {
// 		fmt.Println("good")
// 		return 7
// 	}
// 	return 6
// }
// если мы читаем  сообщение из канала то пишем "we send some message"
// если мы пищем сообщение в канал quit то завершаем функцию
