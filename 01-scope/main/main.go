package main

import "fmt"

func main() {
	a := 3
	{
		var b int = 5
		fmt.Println(a)
		fmt.Println(b)

	}

}
/*b is declared inside an inner scope than it can't never be called from the outer scope
a is declared inside the outer scope then it can be easily called from the inner scope
 */
