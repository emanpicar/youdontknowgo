package core

import "fmt"

func conversion() {
	x := "Bombs Away"

	fmt.Println([]byte(x))

	for i := 48; i <= 90; i++ {
		fmt.Println("converting int32 into string", string(i))
	}
}