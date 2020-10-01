package main

import "fmt"

func main() {

	v := 10

	p1 := &v
	p2 := &p1

	fmt.Println(p2)
	fmt.Println(*p2)
	fmt.Println(**p2)

	fmt.Println(p1)
	fmt.Println(*p1)
	// fmt.Println(**p1) invalid indirect compilation error

	fmt.Println(v)
	*p1 = 20
	fmt.Println(v)
	**p2 = 30
	fmt.Println(v)

	fmt.Println("Hello World")
}
