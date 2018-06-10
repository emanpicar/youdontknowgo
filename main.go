package main

import "fmt"

import (
	_ "os"
	goCore "github.com/emanpicar/youdontknowgo/core"
)

func main() {
	fmt.Println("Hello 世界")

	goCore.Initialize()
}
