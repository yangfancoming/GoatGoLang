package main

import "fmt"

func main() {
	//test1()
	//test2()
	//test3()
	test4()
}

//  声明一个包含 5 个元素的整型数组  并设置为零值
func test1() {
	var myAyyay [5]int
	fmt.Println(myAyyay)
	fmt.Println(len(myAyyay))
	fmt.Println(myAyyay[3])
	myAyyay[3] = 777
	fmt.Println(myAyyay[3])
}

// 用具体值初始化每个元素 (使用数组字面量声明数组)
func test2() {
	myAyyay := [5]int{10, 20, 30, 40, 50}
	fmt.Println(myAyyay)
}

// 容量由初始化值的数量决定
func test3() {
	myAyyay := [...]int{10, 20, 30, 40, 50}
	fmt.Println(len(myAyyay))
}

// 用具体值初始化索引为 1 和 2 的元素 其余元素保持零值
func test4() {
	myAyyay := [5]int{1: 10, 2: 20}
	fmt.Println(myAyyay)
}
