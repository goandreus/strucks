package main

import (
	"fmt"
)

type persona struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := persona{
		first: "Eduar",
		last:  "Tua",
		age:   31,
	}

	fmt.Println(p1)
}
