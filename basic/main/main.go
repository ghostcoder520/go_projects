package main

import (
	"basic/mypkg"
	"fmt"
)

func main() {

	// demo0()
	// demo1()
	// demo2()
	// demo3()
	// demo4()
	// demo5()
	//demo6() // 闭包演示
	// demo7() // defer语句
	// demo8() // panic/recover
	// demo9() // time包
	//demo10() // 指针
	// demo11() // 结构体、序列化(json包)
	// demo12() // 接口
	demo13() // reflect包
	// demo14() // 文件操作

	mypkg.AbcFunc()
	fmt.Printf("\r\nmain线程结束.")
}
