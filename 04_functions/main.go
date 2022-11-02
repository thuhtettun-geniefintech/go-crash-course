package main

import "fmt"

func hello(name string) string {
	return "my name is " + name
}

func main() {
	var content = hello("thu htet tun")
	fmt.Println(content)
}
