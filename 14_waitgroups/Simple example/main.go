package main

import (
	"fmt"
	"sync"
	"time"
)

func myfunc(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("finished executing my go routine ")
	wg.Done()
}
func main() {
	fmt.Println("Go tutorial on wait groups ")

	var wg sync.WaitGroup
	wg.Add(1)
	go myfunc(&wg)
	wg.Wait()

	fmt.Println("Finished executing go program")

}
