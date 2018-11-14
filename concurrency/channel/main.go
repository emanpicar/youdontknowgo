package main

import (
	"log"
	"sync"
)

func main() {
	// chan01()
	// chanBuff()
	// chanDirectional()
	// chanRange()
	// chanSelect()
	// chanCommaOk()
	chanFanIn()
}

func chan01() {
	c := make(chan int)

	go func() {
		c <- 42
		log.Println("test hey hey")
	}()

	log.Println(<-c)
}

func chanBuff() {
	c := make(chan int, 1)
	c <- 96
	// c <- 96
	log.Println(<-c)
}

func chanDirectional() {
	c := make(chan int, 1)

	go chanReceive(c)
	chanSend(c)

	log.Println("===============")
}

func chanReceive(c chan<- int) {
	c <- 69
}

func chanSend(c <-chan int) {
	log.Println(<-c)
}

func chanRange() {
	c := make(chan int)

	go func() {
		for i := 0; i < 20; i++ {
			c <- i
		}

		close(c)
	}()

	for i := range c {
		log.Println(i)
	}

	log.Println("=====")
	log.Println(c)
	log.Println(len(c))
}

func chanSelect() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 20; i++ {
			if (i % 2) == 0 {
				even <- i
			} else {
				odd <- i
			}
		}

		quit <- 999

		close(odd)
		close(even)
		close(quit)
	}()

	for {
		select {
		case v := <-even:
			log.Println("## even", v)
		case v := <-odd:
			log.Println("## odd", v)
		case v := <-quit:
			log.Println("## quit", v)
			return
		}
	}
}

func chanCommaOk() {
	// The boolean variable ok returned by a receive operator indicates whether the received value was sent on the channel (true)
	// or is a zero value returned because the channel is closed and empty (false).
	c := make(chan int)

	go func() {
		// c <- 5

		close(c)
	}()

	v, ok := <-c
	log.Println(v, ok)
	// v, ok = <-c
	// log.Println(v, ok)
}

func chanFanIn() {
	even := make(chan int)
	odd := make(chan int)
	result := make(chan int)

	go fanOutSend(even, odd)
	go fanInReceive(even, odd, result)

	for r := range result {
		// explain race conditions
		log.Println(r)
	}
}

func fanOutSend(even, odd chan<- int) {
	for i := 0; i < 20; i++ {
		if (i % 2) == 0 {
			even <- i
		} else {
			odd <- i
		}
	}

	close(even)
	close(odd)
}

func fanInReceive(even, odd <-chan int, result chan<- int) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	// multiple channel Fanning In in one channel
	go func() {
		for neve := range even {
			result <- neve
		}

		wg.Done()
	}()

	go func() {
		for ddo := range odd {
			result <- ddo
		}

		wg.Done()
	}()

	wg.Wait()
	close(result)
}
