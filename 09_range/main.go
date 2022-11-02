package main

import "fmt"

func main() {

	// range
	ids := []int64{1, 2, 3, 4, 5, 56, 6, 6, 4, 234}

	// I think it is more like a foreach loop
	for _, id := range ids {
		fmt.Println("ID:", id)
	}

}
