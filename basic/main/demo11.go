package main

import (
	"encoding/json"
	"fmt"
)

// 自定义结构体
type Person struct {
	Name    string `json:"name"` // json tag 自定义序列化时的key
	Age     int    `json:"age"`
	Sex     string `json:"gender"`
	Hobbies []string
	Scores  map[string]float64
}

// 自定义类型中的自定义方法
func (p Person) PrintInfo() {
	fmt.Printf("Person姓名:%v 年龄: %v\n", p.Name, p.Age)
}

func (p *Person) SetInfo(name string, age int) {
	p.Name = name
	p.Age = age
}

type myFloat float64

func (m myFloat) PrintInfo() {
	fmt.Println("我是自定义类型myFloat中的自定义方法")
}

type Student struct {
	Person     // 通过结构体嵌套实现【继承】，同时作为匿名字段
	Department string
	Number     string
}

type User struct {
	Username string
	Password string
	Address
}

type Address struct {
	City  string
	Phone string
}

func (s Student) Study() {
	fmt.Printf("%v 在学习.\n", s.Name)
}

func demo11() {
	fmt.Println("============= demo 11 ==============")

	{ // 结构体的实例化两种方式: 返回结构体类型或结构体类型指针
		var person1 Person
		person1.Name = "张三"
		person1.Age = 20
		person1.Sex = "男"

		person2 := new(Person)

		person3 := &Person{}

		person4 := Person{
			Name: "小明",
			Age:  20,
			Sex:  "男",
		}

		person5 := &Person{
			Name: "小赵",
		}

		person6 := &Person{
			"小赵",
			20,
			"男",
			[]string{"唱", "跳", "rap"},
			map[string]float64{
				"c++":  100,
				"java": 59.9,
			},
		}

		fmt.Printf("person1类型: %T\n", person1) //类型: main.Person
		fmt.Printf("person2类型: %T\n", person2) //类型: *main.Person
		fmt.Printf("person3类型: %T\n", person3) //类型: *main.Person
		fmt.Printf("person4类型: %T\n", person4) //类型: main.Person
		fmt.Printf("person5类型: %T\n", person5) //类型: *main.Person
		fmt.Printf("person6类型: %T\n", person6) //类型: *main.Person

		fmt.Printf("person1: %v\n", person1)  // {张三 20 男 [] map[]}
		fmt.Printf("person2: %v\n", person2)  // &{ 0  [] map[]}
		fmt.Printf("person3: %#v\n", person3) // person3: &main.Person{Name:"", Age:0, Sex:"", Hobbies:[]string(nil), Scores:map[string]float64(nil)}
		fmt.Printf("person1: %#v\n", person1) // main.Person{Name:"张三", Age:20, Sex:"男"}
		fmt.Printf("person6: %#v\n", person6) // &main.Person{Name:"小赵", Age:20, Sex:"男", Hobbies:[]string{"唱", "跳", "rap"}, Scores:map[string]float64{"c++":100, "java":59.9}}

		p := Person{}
		p.Name = "张三"
		p.Age = 18
		p.Hobbies = make([]string, 3, 6)
		p.Hobbies[0] = "羽毛球"
		p.Hobbies[1] = "乒乓球"
		p.Scores = make(map[string]float64)
		p.Scores["语文"] = 88.88
		p.Scores["数学"] = 77.77

		fmt.Printf("%#v\n", p)

		fmt.Println()

	}
	{
		p1 := Person{
			Name: "哈哈",
			Age:  20,
			Sex:  "男",
		}

		// 结构体是值类型

		p2 := p1

		fmt.Printf("%#v\n", p1)
		fmt.Printf("%#v\n", p2)

		p1.PrintInfo()
		p1.SetInfo("呵呵", 34)
		p1.PrintInfo()

		var a myFloat = 1.4
		a.PrintInfo()
	}
	{
		var u User
		u.Username = "zyh"
		u.Password = "123456"
		u.Address.City = "beijing"
		u.Address.Phone = "13077775555"

		u.City = "shanghai"
		// 访问字段时，先在父结构体中查找该字段，找不到再在嵌套子结构体中查找该字段
		// 此时如果有多个子结构体包含相同的字段名，则会报错

		fmt.Printf("%#v\n", u) // main.User{Username:"zyh", Password:"123456", Address:main.Address{City:"shanghai", Phone:"13077775555"}}

		fmt.Println(u.Address.Phone) // 13077775555
		fmt.Println(u.Phone)         // 13077775555

		student := Student{
			Department: "信通",
			Person: Person{
				Name: "小A",
				Age:  24,
			},
		}

		student.PrintInfo()
		student.Study()

		p1 := Person{
			Name: "张三",
			Age:  24,
		}

		// 序列化与反序列化
		jsonByte, _ := json.Marshal(p1) // 使用 json.Marshal() 序列化
		jsonStr := string(jsonByte)
		fmt.Println(jsonStr) // {"Name":"张三","Age":24,"Sex":"","Hobbies":null,"Scores":null}
		p2 := Person{}
		err := json.Unmarshal([]byte(jsonStr), &p2)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%#v\n", p2)

	}

}

/*
============= demo 11 ==============
person1类型: main.Person
person2类型: *main.Person
person3类型: *main.Person
person4类型: main.Person
person5类型: *main.Person
person6类型: *main.Person
{张三 20 男 [] map[]}
main.Person{Name:"张三", Age:20, Sex:"男", Hobbies:[]string(nil), Scores:map[string]float64(nil)}
&main.Person{Name:"小赵", Age:20, Sex:"男", Hobbies:[]string{"唱", "跳", "rap"}, Scores:map[string]float64{"c++":100, "java":59.9}}
main.Person{Name:"张三", Age:18, Sex:"", Hobbies:[]string{"羽毛球", "乒乓球", ""}, Scores:map[string]float64{"数学":77.77, "语文":88.88}}

main.Person{Name:"哈哈", Age:20, Sex:"男", Hobbies:[]string(nil), Scores:map[string]float64(nil)}
main.Person{Name:"哈哈", Age:20, Sex:"男", Hobbies:[]string(nil), Scores:map[string]float64(nil)}
Person姓名:哈哈 年龄: 20
Person姓名:呵呵 年龄: 34
我是自定义类型myFloat中的自定义方法
main.User{Username:"zyh", Password:"123456", Address:main.Address{City:"shanghai", Phone:"13077775555"}}
13077775555
13077775555
Person姓名:小A 年龄: 24
小A 在学习.
{"name":"张三","age":24,"gender":"","Hobbies":null,"Scores":null}
main.Person{Name:"张三", Age:24, Sex:"", Hobbies:[]string(nil), Scores:map[string]float64(nil)}
*/
