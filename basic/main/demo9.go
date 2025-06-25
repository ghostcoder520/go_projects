package main

import (
	"fmt"
	"time"
)

func demo9() {
	fmt.Println("============= demo 9 ==============")

	// 使用time包 获取当前时间的Time对象
	timeObj := time.Now()

	fmt.Println(timeObj) //2024-12-07 10:44:46.7516623 +0800 CST m=+0.003617401

	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()

	fmt.Printf("%02d--%02d--%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	// 格式化日期
	/*
		时间类型有一个自带的方法Format进行格式化，
		需要注意的是Go语言中格式化时间模板不是常见的Y-m-d H:M:S
		而是使用Go的诞生时间2006年1月2号15点4分(即
		年 2006
		月 01
		日 02
		时 03 (写成15表示24小时制)
		分 04
		秒 05)
	*/
	fmt.Println(timeObj.Format("2006-01-02 03:04:05")) // 返回的是string
	fmt.Println(timeObj.Format("2006-01-02 15:04:05"))
	fmt.Println(timeObj.Format("2006/01/02 03:04:05"))

	// 从日期字符串获取日期对象
	str := "2024-12-07 11:15:27"
	layout := "2006-01-02 15:04:05"
	timeObj1, _ := time.ParseInLocation(layout, str, time.Local)
	fmt.Println(timeObj1)

	// 从日期对象获取时间戳
	timeObj = time.Now()
	unixTime := timeObj.Unix()
	unixNanoTime := timeObj.UnixNano()
	fmt.Println("当前时间戳(ms): ", unixTime)     // 1733541146
	fmt.Println("当前时间戳(ns): ", unixNanoTime) // 1733541146252530000

	// time.Unix(int64, int64) 将时间戳转换成日期对象
	timeObj = time.Unix(int64(1733541146), 0)
	fmt.Println(timeObj.Format("2006/01/02 15:04:05"))

	/*
		time包中定义的时间间隔类型的常量如下:
		const(
			NanosecondDuration = 1
			Microsecond = 1000 * Nanosecond
			Millisecond = 1000 * Microsecond
			Second = 1000 * Millisecond
			Minute = 60 * Second
			Hour = 60 * Minute
		)
	*/

	// 时间操作函数
	timeObj = time.Now()
	fmt.Println(timeObj)
	timeObj = timeObj.Add(time.Hour)
	fmt.Println(timeObj)

	//golang 定时器
	ticker := time.NewTicker(time.Second) // 开启一个1秒间隔的定时器

	n := 3
	for v := range ticker.C {
		if n == 0 {
			ticker.Stop() // 终止内存中的定时器执行
			break
		}
		fmt.Println(v)
		n--
	}

	fmt.Println("aaa0")
	time.Sleep(time.Millisecond * 500)
	fmt.Println("aaa1")
	time.Sleep(time.Millisecond * 500)
	fmt.Println("aaa2")
	time.Sleep(time.Millisecond * 500)
	fmt.Println("aaa3")

}
