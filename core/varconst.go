package core

import "fmt"

var(
	factoredExplicitVariable string = "explicit assignment"
	factoredImplicitVariable2 = "implicit assignment"
)

var explicitVariable string = "explicit/strict assignment creates typed variable"
var implicitVariable = "implicit assignment creates untyped variable"

const(
	zero int = iota
	one
	two
	three
	hello1 = "Hello"
	hello2
	hello3
)

const explicitConstant string = "explicit/strict assignment creates typed variable"
const implicitConstant = "implicit assignment creates untyped variable"

func varconst() {
	fmt.Println(factoredExplicitVariable)
	fmt.Println(factoredImplicitVariable2)

	fmt.Println(explicitVariable)
	fmt.Println(implicitVariable)

	fmt.Println(explicitConstant)
	fmt.Println(implicitConstant)

	fmt.Println(zero, one, two, three)
	fmt.Println(hello1, hello2, hello3)

	anotherImplicit := 5.55
	fmt.Println(anotherImplicit)
	fmt.Printf("%T \n", anotherImplicit)
}