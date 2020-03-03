package main

import "fmt"

func main() {
	test11()
	test13()
}

// 映射使用
func test11() {
	var myGreeting = make(map[string]string)
	myGreeting["Tim"] = "A"
	myGreeting["Jenny"] = "B."
	fmt.Println(myGreeting)
}

// panic: assignment to entry in nil map
func test13() {
	var myGreeting map[string]string
	fmt.Println(myGreeting)
	myGreeting["Tim"] = "A"
	myGreeting["Jenny"] = "B."
	fmt.Println(myGreeting == nil)
}
