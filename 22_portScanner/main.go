package main

import (
	"fmt"
	"pscan/port"
)

func main() {
	fmt.Println("Port scanner in go")

	open := port.ScanPort("tcp", "localhost", 8080)
	fmt.Printf("Port open: %t\n", open)
	results := port.InitialScan("localhost")
	fmt.Println(results)
}
