package main

import (
	"fmt"
	"time"
)

// Закрытый канал не останавливает горутину! В закрытый канал нельзя ничего записывать, но можно читать, если в нём не будет сообщений, то будет возвращаться nil.
// Проверка на закрытый канал (переменная ok) говорит о том, что канал закрыт и нет возможно ничего прочитать из него

func cancelChannel() {

	ch := make(chan int)

	go func(ch chan int) {
		ch <- 3
		close(ch) // Закрываем канал
	}(ch)

	a, ok := <-ch
	fmt.Println(a, ok) // 3, true (канал не закрыт)

	b, ok1 := <-ch
	fmt.Println(b, ok1) // 0, false (канал закрыт и в нем нет сообщений)

	time.Sleep(1 * time.Second)

}

// func function3(ch chan int) {
// 	for {
// 		fmt.Println(<- ch)
// 		time.Sleep(1 * time.Second)
// 	}
// }
