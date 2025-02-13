package main

import (
	"fmt"
	"time"
)

func SendValues(c chan string) {
	fmt.Println("Executing Go routine")
	time.Sleep(1 * time.Second)
	c <- "Hello world"
	fmt.Println("Fininshed Executing Go routine")
}

func main() {
	fmt.Println("Channels tutorial in go")

	values := make(chan string, 2)
	defer close(values)

	go SendValues(values)
	go SendValues(values)

	value := <-values
	fmt.Println(value)

	time.Sleep(1 * time.Second)
}
