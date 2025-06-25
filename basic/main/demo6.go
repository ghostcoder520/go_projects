package main

import "fmt"

// golang闭包
/*
闭包：有权访问另一个函数作用域的函数
闭包的创建方式：在一个函数fnA内部创建另一个函数fnB，通过fnB访问fnA的局部变量
闭包的作用：通过闭包可以实现让变量既常驻内存又不会像全局变量一样污染全局

闭包作用域内的局部变量资源不会被立即回收，因此不要过度使用闭包


闭包可以理解成"定义在个函数内部的函数"，在本质上，
	闭包是将函数内部和函数外部连接起来的桥梁。或者说是函数和其引用环境的组合体.
*/
func adder() func(int) int {
	fmt.Println("闭包演示：")
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func demo6() {
	fmt.Println("============= demo 6 ==============")
	fn := adder() // 执行adder()并返回了一个特定方法
	// 此时 adder中的局部变量sum 相对于返回的特定方法而言是可被访问的，且常驻内存
	fmt.Println(fn(1)) // 1
	fmt.Println(fn(1)) // 2
	fmt.Println(fn(1)) // 3
}
