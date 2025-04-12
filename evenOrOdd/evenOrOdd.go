package main

import "fmt"

type evenOrOdd []int

func newSlice() evenOrOdd {
	slice := evenOrOdd{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	return slice
}

func (e evenOrOdd) print() {
	for _, num := range e {
		if num%2 == 0 {
			fmt.Println(num, " is even")
		} else {
			fmt.Println(num, " is odd")
		}
	}
}
