package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for i, arg := range os.Args[1:] { // You can use _ instead of i if you don't need it | Range returns index and value
		fmt.Println("Index: ", i, " Value: ", arg)
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
