package main

import "fmt"
const (number1 int = iota
	   number2 )

const(
	_ = iota
	number4 = iota * 5
	number5 = iota * 5
)

const (
	_ = iota
	number6 = 1 << (iota * 10)
)


func main()  {
	fmt.Println(number1)
	fmt.Println(number2)
	/*fmt.Println(number3)*/
	fmt.Println(number4)
	fmt.Println(number5)
	fmt.Println(number6)
	fmt.Printf("%b\t", number6 )
	fmt.Printf("%d\t", number6 )
}
