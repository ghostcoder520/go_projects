package main

import (
	"errors"
	"fmt"
)

func demo8() {
	fmt.Println("============= demo 8 ==============")
	fn1()
	fn2(1, 0)
	fmt.Println("demo8继续执行...")
	fn3()
	fmt.Println("demo8结束")
}

func fn1() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error:", err)
		}
	}()
	panic("手动抛出的异常") // 程序遇到panic会停止执行. 为了使程序能继续执行，需要使用recover()捕获错误
}

func fn2(a, b int) int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error:", err)
		}
	}()
	return a / b
}
func readFile(fileName string) error {
	if fileName == "main.go" {
		return nil
	}
	return errors.New("读取文件失败")
}

func fn3() {
	defer func() {
		e := recover()
		if e != nil {
			fmt.Println("给管理员发送邮件:", "[", e, "]")
		}
	}()
	err := readFile("abc.go")
	if err != nil {
		panic(err)
	}
}

/*
============= demo 8 ==============
error: 手动抛出的异常
error: runtime error: integer divide by zero
demo8继续执行...
给管理员发送邮件: [ 读取文件失败 ]
demo8结束
*/
