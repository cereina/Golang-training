package main

import "fmt"

var	a int = 3

func main() {
	b := func() int	 {
		a ++
		return a
	}
	fmt.Println(b())
}
