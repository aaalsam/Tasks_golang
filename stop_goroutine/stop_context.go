package main

import (
	"context"
	"fmt"
	"time"
)

// defer - откладывает выполнение какого-либо действия
// cancel (WifhTimeOut) - если прошло время, но контекст не отменился, потому что произошла ошибка (вызывается в defer)
// cancel (WithCancel) - в определенный момент вызываем cancel

func stopContext() {

	ctx, cancel := context.WithCancel(context.Background()) // Контекст с отменой

	ch := make(chan int)

	go function4(ch, ctx)

	ch <- 1

	cancel() // Отмена контекста

	time.Sleep(5 * time.Second) // Либо со слипом, либо с группой

}

func function4(ch chan int, ctx context.Context) {
	for {
		select {
		case a := <-ch:
			fmt.Println(a)
		case <-ctx.Done(): // Закрытие канала от имени контекст
			fmt.Println("мы тут")
			return
		}
	}
}
