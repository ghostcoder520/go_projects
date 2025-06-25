package main

import "fmt"

// 定义接口
type Usber interface {
	Start()
	Stop()
}

type Phone struct {
	Name string
}

func (p Phone) Start() {
	fmt.Println("手机", p.Name, "启动")
}

func (p Phone) Stop() {
	fmt.Println("手机", p.Name, "关机")
}

type Camera struct {
}

func (c Camera) Start() {
	fmt.Println("相机开启")
}

func (c Camera) Stop() {
	fmt.Println("相机关机")
}

type Computer struct {
}

type NameSetter interface {
	SetName(string)
}

type NameGetter interface {
	GetName() string
}

// 接口嵌套
type NameGetterAndSetter interface {
	NameGetter
	NameSetter
}

type Dog struct {
	Name string
}

func (d *Dog) SetName(name string) {
	d.Name = name
}

func (d Dog) GetName() string {
	return d.Name
}

// 接口多态 作为函数参数
func (c Computer) ConnectDevice(usb Usber) {
	fmt.Println("电脑开始连接新设备..")
	usb.Start()
	usb.Stop()
}

func show(a interface{}) {
	fmt.Printf("值:%v, 类型:%T\n", a, a)
}
func demo12() {
	fmt.Println("============= demo 12 ==============")
	p := Phone{
		Name: "Huawei Mate70",
	}
	c := Camera{}

	// golang中结构体实现接口的形式：

	var a Usber
	a = p // Phone实现了Usber接口（Phone必须实现接口Usber中的所有方法）

	a.Start()
	a.Stop()

	a = c // Camera实现了Usber接口
	a.Start()
	a.Stop()

	fmt.Println()

	computer := Computer{}
	phone := Phone{
		Name: "Xiaomi 12X",
	}
	camera := Camera{}

	computer.ConnectDevice(phone)
	computer.ConnectDevice(camera)

	dog1 := &Dog{
		Name: "小黑狗",
	}

	// 结构体实现多个接口
	var nameGetter NameGetter = dog1                   // dog1实现 NameGetter 接口
	var nameSetter NameSetter = dog1                   // dog1实现 NameSetter 接口
	var nameGetterAndSetter NameGetterAndSetter = dog1 // dog1实现 NameGetterAndSetter 接口

	fmt.Println(nameGetter.GetName()) // 小黑狗
	nameSetter.SetName("小白狗")
	fmt.Println(nameGetter.GetName())          // 小白狗
	fmt.Println(nameGetterAndSetter.GetName()) // 小白狗

	// 空接口
	{

		// 空接口当做任何类型使用

		var a interface{}

		a = 20
		fmt.Printf("值:%v, 类型:%T\n", a, a) //值:20, 类型:int
		a = "golang"
		fmt.Printf("值:%v, 类型:%T\n", a, a) //值:golang, 类型:string
		a = true
		fmt.Printf("值:%v, 类型:%T\n", a, a) //值:true, 类型:bool

		show(100)
		show("abcd")
		show([]int{1, 2, 3, 4, 5}) //值:[1 2 3 4 5], 类型:[]int

		mp := make(map[int]interface{})                     // 类型: map[int]interface{}
		slice := []interface{}{10, "你好", true, -67.34, 'a'} // 类型: []interface{}

		mp[1] = "cccc"
		mp[0] = 233
		mp[5] = -16.78

		fmt.Printf("%T\n", mp)
		fmt.Println(mp)    // map[0:233 1:cccc 5:-16.78]
		fmt.Println(slice) // [10 你好 true -67.34 97]

	}

	{
		// 接口的类型断言 interface{}.(type)  类似Java 的 instance of
		var a interface{}
		a = "你好golang"
		v, ok := a.(string)

		if ok {
			fmt.Printf("b是string类型, 值为:%v\n", v)
		} else {
			fmt.Printf("类型断言失败")
		}

		// interface.(type) 判断一个接口的实际类型，这个写法只能用在switch语句中
		switch a.(type) {
		case int:
			fmt.Println("int类型")
		case string:
			fmt.Println("string类型")
		case Camera:
			fmt.Println("Camera类型")
		default:
			fmt.Println("其他类型")
		}

	}

	{
		// 空接口深入
		/*
			空接口传递的数组和切片无法使用索引操作，传递的结构体无法访问结构体任何的成员和方法。
			如果需要通过空接口访问上述数据类型，则需要先使用空接口的类型断言获取value，再对value变量进行具体操作
		*/
		var a interface{}

		d1 := Dog{"哈哈"}
		a = d1
		fmt.Print(a, " -> ")
		d1.Name = "呵呵"
		fmt.Println(a) // {哈哈} -> {哈哈}

		// 如果直接使用 a.Name 或 a.GetName() 则报错！
		// 正确使用空接口访问结构体成员的方法：
		v1, _ := a.(Dog)
		fmt.Println(v1.Name)

		arr1 := [3]int{0, 0, 0}
		a = arr1
		fmt.Print(a, " -> ")
		arr1[0] = 1
		fmt.Println(a)        // [0 0 0] -> [0 0 0]
		fmt.Printf("%T\n", a) // [3]int

		// 如果直接使用a[0] 则报错!
		// 正确使用空接口访问数组元素的方法：
		v2, _ := a.([3]int)
		fmt.Println(v2[0])

		slice1 := []int{0, 0, 0}
		a = slice1
		fmt.Print(a, " -> ")
		slice1[0] = 1
		fmt.Println(a)
		fmt.Printf("%T\n", a)

		// 如果直接使用a[0] 则报错!
		// 正确使用空接口访问切片元素的方法：
		v3, _ := a.([]int)
		fmt.Println(v3[0])

	}

	{
		// 对于结构体值类型接收者的接口方法，结构体值类型和结构体指针类型都可以赋值给接口变量
		var p1 = Phone{
			Name: "小米手机",
		}
		var p2 = &Phone{
			Name: "苹果手机",
		}
		var usb Usber = p1
		usb.Start()
		usb = p2
		usb.Start()

		// 对于结构体指针类型接收者的接口方法，只有结构体指针类型可以赋值给接口变量(即上述接口usb只能接收p2，不能接收p1)
	}

}
