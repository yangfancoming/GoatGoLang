package main

import "fmt"

// 数组声明
func main() {
	//test1()
	//test2()
	test3()
}

// 使用长度声明一个字符串切片     其长度和容量都是 5 个元素
func test1() {
	mySlice := make([]string, 5)
	fmt.Println(len(mySlice))
}

// 使用长度和容量声明整型切片  其长度为 3 个元素，容量为 5 个元素
func test2() {
	mySlice := make([]int, 3, 5)
	fmt.Println(len(mySlice))
}

// 记住，如果在[]运算符里指定了一个值，那么创建的就是数组而不是切片。只有不指定值的时候，才会创建切片
// 声明数组和声明切片的不同
func test3() {
	// 创建有 3 个元素的整型数组
	myArray := [3]int{10, 20, 30}
	fmt.Println(len(myArray))
	// 创建长度和容量都是 3 的整型切片
	mySlice := []int{10, 20, 30}
	fmt.Println(len(mySlice))
}
