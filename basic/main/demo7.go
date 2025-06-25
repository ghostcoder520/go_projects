package main

import "fmt"

func myFn71() {
	fmt.Print("开始 ")
	defer func() {
		fmt.Print("aaa ")
	}()
	fmt.Printf("结束")
}

func myFn72() int { // 返回0
	a := 0
	defer func() {
		a++
	}()
	return a
}

func myFn73() (a int) { // 返回6
	defer func() {
		a++
	}()
	return 5
}

func myFn74(index string, a, b int) int {
	fmt.Println(index, a, b, a+b)
	return a + b
}

func demo7() {
	fmt.Println("============= demo 7 ==============")
	// 通过defer延迟执行语句
	myFn71() // 开始 结束 aaa
	fmt.Println()
	fmt.Println(myFn72()) // 0
	fmt.Println(myFn73()) // 6

	{
		x := 1
		y := 2
		defer myFn74("AA", x, myFn74("A", x, y))
		x = 10
		defer myFn74("BB", x, myFn74("B", x, y))
		y = 20

		/*
			defer函数的注册顺序：注册时确定参数，参数中有函数则先执行之
				defer myFn74("AA", x, myFn74("A", x, y))
				defer myFn74("BB", x, myFn74("B", x, y))

			defer函数的执行顺序:
				defer myFn74("BB", x, myFn74("B", x, y))
				defer myFn74("AA", x, myFn74("A", x, y))

			全部函数的实际执行顺序：
				1. myFn74("A", x, y)  A 1 2 返回 3
				2. myFn74("B", x, y)  B 10 2 返回 12
				3. myFn74("BB", x, myFn74("B", x, y))  BB 10 12 22
				4. myFn74("AA", x, myFn74("A", x, y))  AA 1 3 4
		*/
	}

}

/*
============= demo 7 ==============
开始 结束aaa
0
6
A 1 2 3
B 10 2 12
BB 10 12 22
AA 1 3 4
*/
