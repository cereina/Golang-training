package function1

import (
	"fmt"
	"github.com/Cereina/Golang-training/01-function/function"
)

func PrintName1() {
	fmt.Println(firstName)
	fmt.Println(lastName)
	e := function.Add()
	fmt.Println(e())
}
