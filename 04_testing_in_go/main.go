package main

import "fmt"

func calculate(x int) int {
	return x + 2
}

func main() {
	fmt.Println("Go testing tutorial")
	result := calculate(2)
	fmt.Println(result)
}
