package main

import "fmt"

import (
	goCore "github.com/emanpicar/youdontknowgo/core"
	_ "os"
)

func main() {
	fmt.Println("Hello 世界")

	goCore.Initialize()
}
