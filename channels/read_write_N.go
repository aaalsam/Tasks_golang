package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
//По истечению N секунд программа должна завершиться.
//N нужно ввести из консоли.

func writesAndReading() {
	var n int //Завели переменную для сканера

	fmt.Println("Введите N:")
	fmt.Scan(&n) //Считали данные с консоли, записали в n

	wg := new(sync.WaitGroup) // Выделяет память для переменной типа sync.WaitGroup и возвращет указатель на эту память

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(n)*time.Second) // Создали контекст, который отменяется по истечению времени или при отмене функции
	defer cancel()                                                                         // Откладываем фукнцию отмены контекста, если программа закончится раньше тайм-аута

	ch := make(chan string) //Создали канал

	go writes(ctx, ch, wg) //Создали для функции горутину

	go reading(ctx, ch, wg)

	time.Sleep(1 * time.Millisecond)

	wg.Wait()
}

// функция write(читать) бесконечно пишет в канал
func writes(ctx context.Context, ch chan string, wg *sync.WaitGroup) {
	wg.Add(1)
	for {
		select { // Оператор select позволяет go-процедуре находиться в ожидании нескольких операций передачи данных.
		case <-ctx.Done(): // Проверяем не завершился ли контекст
			wg.Done()
			fmt.Println("Завершено")
			return // Если тру, то выходим из функции
		case ch <- "message": //Если мы можем писать 
			fmt.Println("Мы записали в канал")
		}
	}
}

// Функция read(писать) бесконечно читает из канала
func reading(ctx context.Context, ch chan string, wg *sync.WaitGroup) {
	wg.Add(1)
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			fmt.Println("Программа завершена")
			return
		case <-ch:
			fmt.Println("Мы прочитали из канала")
		}
	}
}
