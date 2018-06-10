package core

import (
	"fmt"
)

func controlFlow() {
	forLoop()
}

func forLoop() {
	for i := 0; i < 10; i++ {
		fmt.Println("i := 0; i < 10; i++ ==", i)
	}

	c := 0
	for c < 10 {
		fmt.Println("c < 10 ==", c)
		c++
	}

	b := 0
	for {
		b++

		if (b % 2) == 0 {
			continue
		}

		if b >= 10 { break }

		fmt.Println("no condition loop ==", b)
	}
}