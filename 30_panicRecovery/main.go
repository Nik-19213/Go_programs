package main

import "fmt"

type SomeStruct struct {
	message string
}

func sillySusan() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Susan recovered from the panic")
			fmt.Printf("%+v\n", r)
			fmt.Println("silly susan handles the panic gracefully")
		}
	}()
	fmt.Println("silly susan called")
	panickingPeter()
	fmt.Println("silly susan finished")
}

func panickingPeter() {
	fmt.Println("panicking peter called")
	panic(SomeStruct{message: "hello from peter panic"})
	fmt.Println("panicking peter finished")
}

func main() {
	fmt.Println("cascading panics")

	sillySusan()
}
