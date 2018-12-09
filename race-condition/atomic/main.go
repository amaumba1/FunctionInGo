package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	// declare our variable counter to int64
	var counter int64 // anytime you see int64 in program in Go you should think package atomic

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			// increment our counter using below func
			// func AddInt64(addr *int64, delta int64) (new int64)
			// from atomic.addInt64 which takes an address a pointer to an int64 (so an address where an int64 are stored , and then some delta which is int64)
			atomic.AddInt64(&counter, 1)
			// func LoadInt64(addr *int64) (val int64)
			// to print a counter from package atomic i have to use Load64 to read it and Load64 takes an address a pointer to an int64 where the address is stored  and it returns int64
			fmt.Println("Counter\t", atomic.LoadInt64(&counter)) // give the address to our counter
			runtime.Gosched()                                    // just to switch my go routines
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait() // to all go routine to be finish
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("That's the count:", counter) // the counter here is not being read inside of another goroutine at this point.
}
