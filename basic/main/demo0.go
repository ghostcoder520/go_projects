package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func getUserinfo() (string, int) {
	return "xiaoming", 24
}

func demo0() {
	fmt.Println("============= demo 0 ==============")
	fmt.Println("中文输出测试", "AAA")
	var a float32 = 10.56
	b := a + 10
	fmt.Println(b)

	// 匿名变量，不占用命名空间，不分配内存
	var username, _ = getUserinfo()
	fmt.Println(username)

	// 常量直接 = 赋值
	const PI = 3.14

	const (
		N1 = 1
		N2
		N3
	)

	fmt.Println("iota:-----------------------------")
	const (
		m1 = 1 << iota
		m2
		_
		_
		m3 = 100
		m4
		m5, m6 = iota, iota
		m7     = iota
	)

	a = 1
	b = 2

	fmt.Println(PI, m1, m2, m3, m4, m5, m6, m7)
	fmt.Println("数据类型:-----------------------------")

	var num int = 10   // int8 int16 int32 int64, 默认int/float具体位数根据操作系统位数决定
	var unum uint = 10 // int8 int16 int32 int64
	// %v 原样输出
	fmt.Printf("num = %v, type: %T, size = %d\n", num, num, unsafe.Sizeof(num))
	fmt.Printf("unum = %v, type: %T, size = %d\n", unum, unum, unsafe.Sizeof(unum))

	var f1 = 3.14e2 // float64

	fmt.Printf("f1 = %v, type: %T, size = %d\n", f1, f1, unsafe.Sizeof(f1))

	fmt.Println("字符串string:-----------------------------")

	// str1 := "this is \nstr"
	// str2 := "C:\\Go\\bin"
	str3 := `this is \nstr
    this is \nstr
    this is \nstr
    this is \nstr`
	fmt.Println(str3)

	// string和切片的转换
	str1 := "111-222-333"
	arr := strings.Split(str1, "-")
	fmt.Println(arr)

	str2 := strings.Join(arr, "*")
	fmt.Println(str2)

	ch1 := 'a'
	ch2 := '国'                                            // utf-8编码
	fmt.Printf("%c--%v--%T\n", ch1, ch1, ch1)             // a--97--int32
	fmt.Printf("%c--%v--%T\n", ch2, ch2, ch2)             // 国--22269--int32
	fmt.Printf("%c--%v--%T\n", str3[0], str3[0], str3[0]) // t--116--uint8

	// 修改字符串--只能通过切片修改
	s1 := "大big"
	runeStr := []rune(s1)
	runeStr[0] = '小'
	fmt.Println(string(runeStr)) // 小big

	s2 := "big"
	byteStr := []byte(s2)
	byteStr[0] = 'p'
	fmt.Println(string(byteStr)) // pig
}
