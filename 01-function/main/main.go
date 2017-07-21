package main

import (
	"fmt"
	"github.com/Cereina/Golang-training/01-function/function"
)

var	a int = 3

func add() func() int  {
	c := 0
	return func() int {
		c++
		return c
	}

}

func main() {
	b := func() int	 {
		a ++
		return a
	}
	fmt.Println(b())

	d := add()
	fmt.Println(d())

	e := function.Add()
	fmt.Println(e())
}
