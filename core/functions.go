package core

import "fmt"

func functions() {
	noReturn()
	fmt.Println(singleReturn())
	fmt.Println(returnNaming())

	x, y := multipleReturn()
	fmt.Println(x, y, "multipleReturn func")

	ff, ll := multipleReturnNaming()
	fmt.Println(ff, ll, "multipleReturnNaming func")

	sports := []string{"Soccer", "Baseball", "Chess", "Study"}
	variadicFunc(sports...)
	variadicFunc("Tennis", "Billiard", "Futsal", "Rugby")

	funcExpressions()

	closureExec := closureFunc(5)
	fmt.Println(closureExec(10))

	callbackFunc(func() {
		fmt.Println("This iz callback")
	})

	fmt.Println(recursionFunc(4))

	deferFunc()
}

func noReturn() {
	fmt.Println("NoReturn func")
}

func singleReturn() string {
	return "SingleReturn func"
}

func returnNaming() (x string) {
	x = "returnNaming func"

	return
}

func multipleReturn() (int, int) {
	return 5, 10
}

func multipleReturnNaming() (fname, lname string) {
	fname = "Eman"
	lname = "Picar"

	return
}

func variadicFunc(sports ...string) {
	fmt.Println(sports)
	fmt.Printf("%T\n", sports)
}

func funcExpressions() {
	giveMeApple := func() string {
		return "Apples"
	}

	fmt.Println(giveMeApple())
	fmt.Printf("%T\n", giveMeApple)
}

func closureFunc(a int) func(b int) int {
	return func(b int) int {
		return a * b
	}
}

func callbackFunc(callback func()) {
	callback()
}

func recursionFunc(x int) int {
	if x == 0 {
		return 1
	}

	return x * recursionFunc(x - 1)
}

func deferFunc() {
	defer func() {
		fmt.Println("Counter Terrorist Win")
	}()

	func() {
		fmt.Println("Server will restart in 5 seconds")
	}()

	func() {
		fmt.Println("Go go go!! Need Backup!! Do not run we are your friends!!")
	}()
}