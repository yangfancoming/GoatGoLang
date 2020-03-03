package main

import "fmt"

func main() {
	var x [6]int
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[3])
	x[3] = 777
	fmt.Println(x[3])
}
