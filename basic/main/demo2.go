package main

import "fmt"

func demo2() {
	fmt.Println("============= demo 2 ==============")
	// fmt.Printf("--------------运算符-------------\n")
	a, b := 1, 2
	a, b = b, a
	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Println(-8 >> 1)   // -4
	fmt.Println(-8 << 1)   // -16
	fmt.Println(-8<<1 + 1) // -15

	if age := 24; age > 18 { // 这里定义的age作用域在if{}里
		fmt.Printf("%d 是成年人\n", age)
	}

	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	i := 1
	for ; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	i = 1
	for i <= 10 {
		fmt.Print(i, " ")
		i++
	}

	fmt.Println()

	// for range 键值循环
	var str = "你好go"
	for k, v := range str {
		fmt.Println("key = ", k, "value = ", v)
	}

	var arr = []string{"php", "java", "node", "golang"}
	for i := 0; i < len(arr); i++ {
		break
	}
	for _, v := range arr {
		fmt.Println(v)
	}

label:
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				break label // break label实现多重跳出
			}
			fmt.Printf("%d-%d\n", i, j)
		}
	}
	fmt.Println()

label1:
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				continue label1 // 这里相当于直接写一个 break
			}
			fmt.Printf("%d-%d\n", i, j)
		}
	}
	fmt.Println()

	// switch 写法1 判断变量
	ext := ".html"
	switch ext {
	case ".html":
		fmt.Println("text/html")
	case ".css", ".js":
		fmt.Println("text/css or js")
	default:
		fmt.Println("找不到对应类型")
	}

	// switch 写法2 分支使用表达式 此时switch语句后面不需要再加判断变量
	age := 24
	switch {
	case age < 18:
		fmt.Println("未成年")
	case age >= 18:
		fmt.Println("成年")
	}

}

/*
============= demo 2 ==============
a = 2, b = 1
-4
-16
-15
24 是成年人
1 2 3 4 5 6 7 8 9 10
1 2 3 4 5 6 7 8 9 10
1 2 3 4 5 6 7 8 9 10
key =  0 value =  20320
key =  3 value =  22909
key =  6 value =  103
key =  7 value =  111
php
java
node
golang
0-0
0-1
0-2

0-0
0-1
0-2
1-0
1-1
1-2
2-0
2-1
2-2

text/html
成年
*/
