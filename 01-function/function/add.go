package function

func Add() func() int {
	Number1 = 5+1
	return func() int {
		return Number1
	}

}