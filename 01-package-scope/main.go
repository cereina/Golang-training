package main

import "fmt"

var x int = 42
var a string = "Mazen"

func main() {
	fmt.Println(x)
	foo()
	max := max (7)
	fmt.Println("Maximum = ", max)
}

func foo() {
	fmt.Println(x)
	fmt.Println(a)
}
func max(x int) int {
	return 42 + x
}

