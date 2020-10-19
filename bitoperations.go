package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 3
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a &^ b)
	c := 8
	fmt.Println(c << 3)
	fmt.Println(c >> 3)
	n := 3.14
	n = 13.7e72
	n = 2.1E14
	fmt.Printf("%v, %T", n, n)
	s := "this is a string"
	d := []byte(s)
	fmt.Printf("%v, %T\n", d, d)
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)
}