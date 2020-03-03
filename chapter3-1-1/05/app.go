package main

import "fmt"

func main() {
	test11()
	test12()
	//test13()
}

// make 正常方式
func test11() {
	var myGreeting = make(map[string]string)
	common(myGreeting)
}

// make 简写方式
func test12() {
	myGreeting := make(map[string]string)
	common(myGreeting)
}

func common(myGreeting map[string]string) {
	myGreeting["Tim"] = "A"
	myGreeting["Jenny"] = "B."
	fmt.Println(myGreeting)
}
