package main

import "fmt"

func main() {

	// map
	emails := make(map[string]string)

	// Assing key value
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))

	//loop
	var fruitArr [4]string

	fruitArr[0] = "Apple"
	fruitArr[1] = "orange"

	fmt.Println(fruitArr)

	fruit := [2]string{"Apple", "GreenTea"}

	fmt.Println(fruit)
}
