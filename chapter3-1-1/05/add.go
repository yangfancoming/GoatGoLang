package main

import "fmt"

var myGreeting = map[string]string{"Tim": "T!", "Jenny": "J!"}

func main() {
	fmt.Println(myGreeting)
	//add1()
	//add2()
	//update()
	//delete1()
	delete2()
}

// 添加操作
func add1() {
	myGreeting["Harleen"] = "H!"
	fmt.Println(myGreeting)
}

// 获取map长度
func add2() {
	fmt.Println(len(myGreeting))
}

// 修改操作
func update() {
	myGreeting["Jenny"] = "F!"
	fmt.Println(myGreeting)
}

// 删除操作
func delete1() {
	delete(myGreeting, "Tim")
	fmt.Println(myGreeting)
}

// 删除操作： 对不存在的key进行删除操作 不会报错
func delete2() {
	delete(myGreeting, "goat")
	fmt.Println(myGreeting)
}
