package main

import (
	"errors"
	"fmt"
	"sync"
)

//Реализовать конкурентную запись данных в map

type Storage struct { // Представляем новую структуру, которая предоставляет собственный мьютек
	sync.RWMutex
	map1 map[int]int
}

func main() {

	val := &Storage{ // Создаем объект структуры
		map1: map[int]int{},
	}

	writeAndReadMap(val)

}

func writeAndReadMap(val *Storage) {
	wg := new(sync.WaitGroup)

	for i := 0; i < 5; i++ {
		wg.Add(1) // Добавляем в группу
		go func(i int) {
			defer wg.Done() // По завершению функции, то есть после return
			val.Set(i)      // Добавляем значение i в мапу
		}(i)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			num, err := val.Get(i)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(num)
		}(i)
	}

	wg.Wait()
}

func (s *Storage) Set(num int) {
	s.Lock()
	defer s.Unlock()
	s.map1[num] = num
}

func (val *Storage) Get(num int) (int, error) {
	val.RLock()
	defer val.RUnlock()
	if number, ok := val.map1[num]; ok {
		return number, nil
	}
	return 0, errors.New("number does not exist")
}
