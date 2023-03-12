package main

import (
	"errors"
	"fmt"
	"sync"
)

//Реализовать конкурентную запись данных в map

type Value struct { // Представляем новую структуру, которая предоставляет собственный мьютек
	sync.RWMutex
	map1 map[int]int
}

func main() {

	val := &Value{ // Создаем объект структуры
		map1: map[int]int{},
	}

	writeAndReadeMap(val)

}

func writeAndReadeMap(val *Value) {
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1) // Добавляем в группу
		go func(i int) {
			defer wg.Done()
			val.Add(i) // Добавляем значение i в мапу
		}(i)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			num, err := val.Get(i)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(num)
			}
		}(i)
	}

	wg.Wait()
}

func (val *Value) Add (num int) {
	val.Lock()
	defer val.Unlock()
	val.map1[num] = num
}

func (val *Value) Get (num int) (int, error) {
	val.RLock()
	defer val.RUnlock()
	if number, ok := val.map1[num]; ok {
		return number, nil
	}
	return 0, errors.New("number does not exists")
}
