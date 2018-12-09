package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Begin CPU:", runtime.NumCPU())
	fmt.Println("Begin Goroutine:", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("Hello Mau")
		wg.Done()

	}()
	go func() {
		fmt.Println("It's your birthday!")
		wg.Done()
	}()

	fmt.Println("Mid CPU:", runtime.NumCPU())
	fmt.Println("Mid Goroutine:", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("about to exit")
	fmt.Println("End CPU:", runtime.NumCPU())
	fmt.Println("End Goroutine:", runtime.NumGoroutine())
}
