package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 42
	var j string
	j = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", j, j)
}