package function

func Add() func() int {
	number1 = 5+1
	return func() int {
		return number1
	}

}