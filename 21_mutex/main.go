package main

import (
	"fmt"
	"sync"
)

var (
	mutex   sync.Mutex
	balance int
)

func deposit(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Depositing %d in balance of %d\n", value, balance)
	balance += value
	mutex.Unlock()
	wg.Done()
}

func withdraw(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Withdrawing %d from balance of %d\n", value, balance)
	balance -= value
	mutex.Unlock()
	wg.Done()
}

func main() {
	fmt.Println("Hello world")

	balance = 1000

	var wg sync.WaitGroup

	wg.Add(2)
	go deposit(500, &wg)
	go withdraw(700, &wg)
	wg.Wait()

	fmt.Printf("Current balance: %d", balance)

}
