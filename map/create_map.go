package main

import "fmt"

//Реализовать запись данных в map

func main() {
	
	map1 := map[int]string {
		1: "Nastya",
		2: "Dasha",
		3: "Sasha",
	}
	map1[4] = "Pobert"
	map1[5] = "Vasya"
	map1[6] = "Petya"

	iterationMap(map1)

	fmt.Println("")

	delete(map1, 2)

	iterationMap(map1)


	fmt.Println("\n",map1[3])

}

func iterationMap(map0 map[int]string) {
	for key, value := range map0 {
		fmt.Println(key, value)
	}
}
