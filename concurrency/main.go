package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var mu sync.Mutex
var result int64

func main() {
	log.Println("NumCPU: ========", runtime.NumCPU())
	num := 100
	wg.Add(num)

	for i := 0; i < num; i++ {
		// explain race conditions
		go mutexCounter(fmt.Sprintf("e%v", i), i)
		log.Println("NumGoroutine: ==", runtime.NumGoroutine())
	}

	wg.Wait()
	log.Println("result:", result)
}

func normalCounter(ex string, i int) {
	newNum := result
	log.Println(ex)
	newNum++
	result = newNum
}

func mutexCounter(ex string, i int) {
	mu.Lock()
	// log.Println("NumGoroutine: ==", runtime.NumGoroutine())
	newNum := result
	log.Println(ex)
	runtime.Gosched()
	newNum++
	result = newNum
	mu.Unlock()
	wg.Done()
}

func atomicCounter(ex string, i int) {
	// log.Println("NumGoroutine: ==", runtime.NumGoroutine())
	atomic.AddInt64(&result, 1)
	runtime.Gosched()
	log.Println(ex, "=======", atomic.LoadInt64(&result))
	wg.Done()
}
