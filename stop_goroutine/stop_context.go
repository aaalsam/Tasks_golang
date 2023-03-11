package main

import (
	"context"
	"fmt"
	"time"
)

func stopContext() {

	ctx, cancel := context.WithCancel(context.Background()) // Контекст с отменой

	defer cancel() // Отмена контекста

	ch := make(chan int)

	go function4(ch, ctx)

	ch <- 1

	time.Sleep(1 * time.Millisecond) // Либо со слипом, либо с группой

}

func function4(ch chan int, ctx context.Context) {
	for {
		select {
		case a := <-ch:
			fmt.Println(a)
		case <-ctx.Done(): // Закрытие канала от имени контекста
			return
		}
	}
}
