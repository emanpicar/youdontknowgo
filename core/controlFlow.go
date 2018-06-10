package core

import (
	"fmt"
)

func controlFlow() {
	forLoop()
	boolExpressions()
	switchStatements()
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

func boolExpressions() {
	if !true {
		fmt.Println("! Operation")
	} else if true && false {
		fmt.Println("&& Operation")
	} else if false || false {
		fmt.Println("|| Operation")
	} else {
		fmt.Println("boolExpressions")
	}
}

func switchStatements() {
	// break not needed
	switch "Baseball" {
	case "Billiard", "Bowling":
		fmt.Println("Billiard")
	case "Tennis":
		fmt.Println("Tennis")
	case "Baseball":
		fmt.Println("Baseball")
		fallthrough
	case "Badminton":
		fmt.Println("Badminton")
		fallthrough
	case "Chess":
		fmt.Println("Chess")
	case "Soccer":
		fmt.Println("Soccer")
	default:
		fmt.Println("Aye, have you no Sports?")
	}

	type CustomType struct {
		fname string
		lname string
	}

	var x interface{} = 2.5
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case CustomType:
		fmt.Println("CustomType")
	default:
		fmt.Printf("%T", x)
	}
}