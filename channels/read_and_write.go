package main

import (
	"fmt"
	"sync"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.


func readAndWrite() {
	ch := make(chan string)
	wg := new(sync.WaitGroup) // Выделяет память для переменной типа sync.WaitGroup и возвращет указатель на эту память

	go write(ch, wg) 

	go read(ch, wg)

	time.Sleep(1*time.Millisecond) // Ожидание 1 млс, чтобы горутины успели запуститься
	wg.Wait() // Ожидаем 0 в wg, говорящее о завершении горутин 
}

// функция write бесконечно пишет в канал
func write(ch chan string, wg *sync.WaitGroup) {
	wg.Add(1) // Добавляем в wg горутину
	for {
		select {
		case ch <- "message":
			fmt.Println("Мы записали в канал")
		}
	}
}

// Функция read бесконечно читает из канала 
func read(ch chan string, wg *sync.WaitGroup) {
	wg.Add(1)
	for {
		select {
		case a := <-ch:
			fmt.Println("Мы прочитали из канала", a)
		}
	}
}
