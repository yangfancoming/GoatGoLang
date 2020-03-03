package main

import "fmt"

func main() {
	add1()
	add2()
	//add1()
}

var myGreeting = map[string]string{"Tim": "A!", "Jenny": "B!"}

// 添加操作
func add1() {
	myGreeting["Harleen"] = "C!"
	fmt.Println(myGreeting)
}

// 获取map长度
func add2() {
	fmt.Println(len(myGreeting))
}
