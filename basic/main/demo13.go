package main

import (
	"fmt"
	"reflect"
)

type myInt int
type myStruct struct {
	Name string
	Age  int
}

func reflectTypeFn(x interface{}) {
	// 获取类型对象
	v := reflect.TypeOf(x)
	fmt.Printf("类型:%v, 类型名称:%v, 类型大种类:%v\n", v, v.Name(), v.Kind())
}

func reflectValueFn(x interface{}) {
	// 获取值对象
	v := reflect.ValueOf(x)
	kind := v.Kind()
	switch kind {
	case reflect.Int:
		fmt.Printf("%v\n", v.Int())
	case reflect.Float64:
		fmt.Printf("%v\n", v.Float())
	default:
		fmt.Printf("未判断的类型\n")
	}
}

func mySetValue(x interface{}) {
	v, _ := x.(*int)
	*v = 1
}

func reflectSetValueByPointer(x interface{}) {
	v := reflect.ValueOf(x)

	fmt.Println(v.Kind())        //ptr
	fmt.Println(v.Elem().Kind()) //int

	if v.Elem().Kind() == reflect.Int {
		v.Elem().SetInt(2)
	} else if v.Elem().Kind() == reflect.Float64 {
		v.Elem().SetFloat(2.2)
	} else if v.Elem().Kind() == reflect.String {
		v.Elem().SetString("hello golang")
	} else {
		fmt.Println("未判断的类型")
	}

}

func reflectStructField(s interface{}) {

	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传参错误 非结构体无法使用结构体反射!")
		return
	}

	//  Type.NumField 获取字段个数
	fieldCount := t.NumField()
	fmt.Printf("结构体有 %v 个属性\n", fieldCount)

	//  Type.Field() 获取第i个字段的元信息
	field0 := t.Field(0)              // 返回 reflect.StructField的结构体
	fmt.Println(field0)               // {Person  main.Person  0 [0] true}
	fmt.Println("字段名称:", field0.Name) // Person
	fmt.Println("字段类型:", field0.Type) // main.Person
	fmt.Println("字段标签(key为json):", field0.Tag.Get("json"))
	fmt.Println("--------- 全部字段 ---------")
	for i := 0; i < fieldCount; i++ {
		f := t.Field(i)
		fmt.Printf("字段名称:%v 字段值:%v 字段类型:%v 字段Tag:%v\n", f.Name, v.Field(i), f.Type, f.Tag.Get("json"))
	}
	fmt.Println("---------------------------")

	//  Type.FieldByName() 根据名称获取字段
	field1, ok := t.FieldByName("Department")
	if ok {
		fmt.Println("字段名称:", field1.Name) // Department
		fmt.Println("字段类型:", field1.Type) // string
		fmt.Println("字段标签(key为json):", field1.Tag.Get("json"))
	}

	//  Value获取结构体属性对应的值
	fmt.Println(v.FieldByName("Department")) // 信通
	fmt.Println(v.FieldByName("Person"))     // {小明 24  [] map[]}

}

func reflectStructMethod(s interface{}) {

	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传参错误 非结构体无法使用结构体反射!")
		return
	}

	// 通过获取Type变量 反射获取结构体的方法信息
	// 反射只能获取到公有方法

	// 根据索引获取方法时，索引顺序是方法名的ASCII顺序
	method0 := t.Method(0)    // 返回relfect.Method 类型
	fmt.Println(method0.Name) // PrintInfo
	fmt.Println(method0.Type) // func(main.Student)

	method1, ok := t.MethodByName("Study")
	if ok {
		fmt.Println(method1.Name) // Study
		fmt.Println(method1.Type) // func(main.Student)
	}

	// 通过获取Value变量 反射调用方法
	v.Method(0).Call(nil) // nil表示不传入任何参数
	v.MethodByName("Study").Call(nil)

	var params []reflect.Value
	params = append(params, reflect.ValueOf("小红"))
	params = append(params, reflect.ValueOf(22))
	v.MethodByName("SetInfo").Call(params) // 反射调用方法时，传入参数切片

	v.MethodByName("PrintInfo").Call(nil)
}

func reflectChangeStruct(s interface{}) {

	t := reflect.TypeOf(s)
	if t.Kind() != reflect.Pointer || t.Elem().Kind() != reflect.Struct {
		fmt.Println("传参错误 非结构体指针无法通过反射修改字段值!")
		return
	}

	v := reflect.ValueOf(s)

	// 反射修改结构体字段值

	v.Elem().FieldByName("Department").SetString("计算机")
	v.Elem().FieldByName("Person").FieldByName("Name").SetString("小王")
	v.Elem().FieldByName("Person").FieldByName("Age").SetInt(33)

	v.MethodByName("PrintInfo").Call(nil)

}

func demo13() {
	fmt.Println("============= demo 13 反射==============")
	reflectTypeFn(10)
	reflectTypeFn(3.14)
	reflectTypeFn("hello")
	reflectTypeFn(true)

	var a myInt = 10
	var b = myStruct{
		Name: "张三",
		Age:  20,
	}
	var c = 20
	var d = 1.5
	var slice = []int{1, 2, 3}
	var arr = [3]int{1, 2, 3}

	reflectTypeFn(a)     //类型:main.myInt, 类型名称:myInt, 类型大种类:int
	reflectTypeFn(b)     //类型:main.myStruct, 类型名称:myStruct, 类型大种类:struct
	reflectTypeFn(&c)    //类型:*int, 类型名称:, 类型大种类:ptr
	reflectTypeFn(&d)    //类型:*float64, 类型名称:, 类型大种类:ptr
	reflectTypeFn(slice) //类型:[]int, 类型名称:, 类型大种类:slice
	reflectTypeFn(arr)   //类型:[3]int, 类型名称:, 类型大种类:array

	reflectValueFn(1.1)
	reflectValueFn(1)

	num := 0
	mySetValue(&num)
	fmt.Println(num)

	reflectSetValueByPointer(&num)
	fmt.Println(num)

	{
		// 结构体反射
		stu1 := Student{
			Person: Person{
				Name: "小明",
				Age:  24,
			},
			Department: "信通",
		}

		reflectStructField(stu1)
		reflectStructMethod(&stu1)
		reflectChangeStruct(&stu1)

	}
}
