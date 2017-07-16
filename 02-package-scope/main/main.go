package main

import (
	"fmt"
	"github.com/Cereina/Golang-training/02-package-scope/visibility"
)

func main() {
	fmt.Println(visibility.MyName)
	visibility.PrintVar()
}