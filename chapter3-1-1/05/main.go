package main

import "fmt"

// 映射声明
func main() {
	test1()
	test2()
	test3()
}

// 使用 make 声明映射  创建一个映射，键的类型是 string，值的类型是 int
func test1() {
	dict := make(map[string]int)
	fmt.Println(dict)
}

// 创建一个映射，键和值的类型都是 string使用两个键值对初始化映射
func test2() {
	dict := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	fmt.Println(dict)
}

// 尝试给null map 赋值后报错：  panic: assignment to entry in nil map
func test3() {
	var myGreeting map[string]string
	fmt.Println(myGreeting)
	myGreeting["Tim"] = "A"
	myGreeting["Jenny"] = "B."
	fmt.Println(myGreeting == nil)
}
