/*
	Variables
	- In Go, variables are explicitly declared and used by the compiler to e.g. check type-correctness of function calls.


*/

package main

import "fmt"

func main() {
	// var declares 1 or more variables.
	var a = "initial"
	fmt.Println(a)

	// You can declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go will infer the type of initialized variables.
	var d = true
	fmt.Println(d)

	// Variables declared without a corresponding initialization are zero-valued.
	var e int
	fmt.Println(e)

	// The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case.
	f := "apple"
	fmt.Println(f)
}