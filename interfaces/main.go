package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//eb := englishBot{}
	//sb := spanishBot{}

	//printGreeting(eb)
	//printGreeting(sb)

	//s := square{sideLength: 5}
	//t := triangle{height: 5, base: 5}

	//printArea(s)
	//printArea(t)

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
