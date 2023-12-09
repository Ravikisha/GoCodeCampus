package main

import "fmt"

func fu() *int {
	v := 1
	return &v
}

func fu2(v int) *int {
	return &v
}

func main() {
	var x, y int
	var f, g *int
	f = &x
	g = &y
	fmt.Println("x:", x, "y:", y, "f:", f, "g:", g)

	*f = 1
	*g = 2

	fmt.Println("x:", x, "y:", y, "f:", f, "g:", g)

	fmt.Println(fu())
	fmt.Println(fu())

	fmt.Println(fu() == fu())
	fmt.Println(fu2(1) == fu2(1))

	fmt.Println(incr(f))
	fmt.Println(incr(f))

	fmt.Println("f:", *f)

}

func incr(p *int) int {
	*p++
	return *p
}
