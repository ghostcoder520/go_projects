package main

import (
	"fmt"
	"strconv"
)

func demo1() {

	fmt.Println("============= demo 1 ==============")
	// 其他类型 --> string
	var a1 int8 = 10
	var a2 int16 = 20
	var a3 float32 = 1.5
	var a4 bool = true
	c := int16(a1) + a2 + int16(a3) // 必须显式类型转换
	fmt.Printf("%v--%T\n", c, c)

	str1 := fmt.Sprintf("%d-%d-%.2f-%t", a1, a2, a3, a4) // %t 表示bool类型输出
	fmt.Printf("%v--%T\n", str1, str1)

	// strconv包实现string与其他类型互转
	var b1 int = 20
	str2 := strconv.FormatInt(int64(b1), 10)
	fmt.Printf("%v--%T\n", str2, str2)

	var b2 float32 = 12.12345
	str3 := strconv.FormatFloat(float64(b2), 'f', 4, 32)
	fmt.Printf("%v--%T\n", str3, str3)

	str4 := "123456"
	sum, _ := strconv.ParseInt(str4, 10, 32)
	fmt.Printf("%v--%T\n", sum, sum)

	str5 := "123456.666"
	sum1, _ := strconv.ParseFloat(str5, 32)
	fmt.Printf("%v--%T\n", sum1, sum1)

}

/*
============= demo 1 ==============
31--int16
10-20-1.50-true--string
20--string
12.1235--string
123456--int64
123456.6640625--float64
abc work..
*/
