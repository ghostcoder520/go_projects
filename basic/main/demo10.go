package main

import "fmt"

func demo10() {
	fmt.Println("============= demo 10 ==============")
	//指针

	var a int = 10
	var p = &a
	fmt.Printf("a的值:%v a的地址:%p\n", a, &a)
	fmt.Printf("p的值:%v p的类型:%T p的地址:%p\n", p, p, &p) // *int 类型的指针变量
	fmt.Println(*p)

	var p1 *int
	fmt.Println(p1 == nil) //true

	// new(Type) 函数为指定类型分配内存并返回对应地址，地址对应值为零值
	p2 := new(int)
	fmt.Printf("p2的值:%v  p2的类型:%T  *p2取内容:%v\n", p2, p2, *p2) //p2的值:0xc00000a240 p2的类型:*int 取内容:0

	s1 := []int{1, 1, 1}
	p3 := &s1
	fmt.Printf("p3的值:%v  p3的类型:%T  *p3取内容:%v\n", p3, p3, *p3) //p3的值:&[1 1 1]  p3的类型:*[]int  *p3取内容:[1 1 1]
	
}

/*
============= demo 10 ==============
a的值:10 a的地址:0xc00000a0e8
p的值:0xc00000a0e8 p的类型:*int p的地址:0xc000062058
10
true
p2的值:0xc00000a110  p2的类型:*int  *p2取内容:0
p3的值:&[1 1 1]  p3的类型:*[]int  *p3取内容:[1 1 1]
*/
