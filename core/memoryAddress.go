package core

import "fmt"

func memoryAddress() {
	justName := "Eman"
	anotherName := &justName

	fmt.Println(anotherName)
	fmt.Println(*anotherName)

	changeName(&justName)
	fmt.Println("after changeName ==>", &justName, justName)
}

func changeName(newJustName *string) {
	fmt.Println("changeName ==>", newJustName, *newJustName)
	*newJustName = "Picar"
}