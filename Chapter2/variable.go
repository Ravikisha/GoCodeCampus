package main

import "fmt"

func main() {
	var i int
	j := 1
	// var k  // error: missing variable type or initialization
	var f float64 = 1.0
	var b, s = true, "yes"
	fmt.Printf("%v %v %f %v %v\n", i, j, f, b, s)
	// fmt.Printf("%v\n", k)

}
