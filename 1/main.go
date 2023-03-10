package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

func main() {
	
	array := [5]int {2, 4, 6, 8, 10}
	
	for _, i := range array {
		
		
		go numberSquared(i)
	}

}

func numberSquared (number int) {
	fmt.Println(number *number)
	mu.Unlock()
}