package main

import "fmt"

var x int = 42
var a string = "Mazen"

func main() {
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println(x)
	fmt.Println(a)
}

