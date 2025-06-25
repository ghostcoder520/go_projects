package main

import (
	"fmt"
	"sort"
)

func demo3() {
	{
		fmt.Println("============= demo 3 ==============")
		var arr1 [3]int
		var arr2 = [3]int{2, 4, 6}
		var arr3 = [...]int{2, 4, 6, 8, 10}   // 自行推断数组长度
		var arr4 = [...]int{0: 1, 1: 2, 4: 5} // k:v 按下标初始化
		var strArr [3]string
		fmt.Printf("arr1: %T, strArr: %T\n", arr1, strArr)
		fmt.Println(arr1)           // [0 0 0]
		fmt.Println(strArr)         // []
		fmt.Println(len(strArr[0])) // 0
		fmt.Println(arr2)           // [2 4 6]
		fmt.Println(arr3)           // [2 4 6 8 10]
		fmt.Println(arr4)           // [1 2 0 0 5]

		arr5 := [...][2]string{
			{"a00", "a01"},
			{"a10", "a11"},
			{"a20", "a21"}, // 多维数组必须使用多层大括号，每一行都有逗号
		}

		for _, v := range arr5 {
			for _, r := range v {
				fmt.Printf("%v-", r)
			}
		}
		fmt.Println()

	}

	{ // golang中，基本数据类型和**数组** 都是值类型
		// 数组是值类型
		arr1 := [3]int{1, 2, 3}
		arr2 := arr1
		arr2[0] = 11
		fmt.Println(arr1) //[1 2 3]
		fmt.Println(arr2) //[11 2 3]
		// 切片是引用类型
		arr3 := []int{1, 2, 3} // 切片
		arr4 := arr3
		arr4[0] = 11
		fmt.Println(arr3) //[11 2 3]
		fmt.Println(arr4) //[11 2 3]
	}

	// 切片 声明切片的方式类似于数组
	{
		var arr1 []int
		var arr2 = []int{1, 2, 3, 4}
		fmt.Println(arr1)        //[]
		fmt.Println(arr1 == nil) //true
		fmt.Println(arr2 == nil) //false

		var strSlice = []string{"php", "java", "nodejs", "golang"}
		fmt.Println(strSlice)
		// 切片遍历跟数组相同

		// 基于数组定义切片
		a := [8]int{0, 1, 2, 3, 4, 5, 6, 7}
		b1 := a[:]
		fmt.Printf("%v - %T\n", b1, b1) // b1的类型是int切片：[]int
		b2 := a[1:2]                    // 左闭右开
		fmt.Printf("%v - %T\n", b2, b2) // [20]
		b3 := a[1:]
		b4 := a[:2]
		fmt.Printf("%v - %T\n", b3, b3) //[20 30 40 0]
		fmt.Printf("%v - %T\n", b4, b4) //[10 20]
		b4[0] = 88
		fmt.Println(b1)
		fmt.Println(b4)
		fmt.Println(a)

		// 基于切片定义新的切片
		// 切片的长度和容量
		// 切片的容量 = 从切片的第一个元素下标开始到其底层数组末尾下标的长度
		fmt.Printf("长度:%d, 容量:%d\n", len(b1), cap(b1)) //长度:8, 容量:8

		c1 := b1[3:5]
		c2 := b1[3:]
		c3 := b1[:5]
		c3[0] = 999
		fmt.Println(b1)
		fmt.Println(c3)

		fmt.Printf("长度:%d, 容量:%d\n", len(c1), cap(c1)) //长度:2, 容量:5
		fmt.Printf("长度:%d, 容量:%d\n", len(c2), cap(c2)) //长度:5, 容量:5
		fmt.Printf("长度:%d, 容量:%d\n", len(c3), cap(c3)) //长度:5, 容量:8

		// make()创建切片
		var sliceA = make([]int, 4, 10)
		fmt.Println(sliceA)                                    // [0 0 0 0]
		fmt.Printf("长度:%d, 容量:%d\n", len(sliceA), cap(sliceA)) // 长度:4, 容量:10

		// 切片元素的修改: 可以像修改数组元素一样使用下标索引修改

		sliceB := []string{"php", "java", "go"}
		sliceB[2] = "golang"

	}

	{
		// 空切片
		var sliceA []int

		// append添加元素
		fmt.Printf("%v 长度:%d, 容量:%d\n", sliceA, len(sliceA), cap(sliceA)) //长度:0, 容量:0
		sliceA = append(sliceA, 10)
		fmt.Printf("%v 长度:%d, 容量:%d\n", sliceA, len(sliceA), cap(sliceA)) //长度:1, 容量:1
		sliceA = append(sliceA, 20, 30, 40)
		fmt.Printf("%v 长度:%d, 容量:%d\n", sliceA, len(sliceA), cap(sliceA)) //长度:4, 容量:4

		// append合并切片
		slice1 := []string{"php", "java"}
		slice2 := []string{"go", "cpp"}

		slice1 = append(slice1, slice2...)
		fmt.Println(slice1) //[php java go cpp]

		// 切片的扩容策略
		var slice3 []int
		for i := 1; i <= 10; i++ {
			slice3 = append(slice3, i)
			// fmt.Printf("%v 长度:%d, 容量:%d\n", slice3, len(slice3), cap(slice3))
		}

		// copy() 深拷贝切片
		slice4 := make([]int, 10, 10)

		copy(slice4, slice3)

		slice4[0] = 111
		fmt.Println(slice3)
		fmt.Println(slice4)

		// sort
		intList := []int{7, 6, 5, 4, 3, 2, 1}
		floatList := []float64{19.55, 18.5, -25, -0.88, 45.88}
		stringList := []string{"apple", "app", "abc", "b", "bcd"}

		sort.Ints(intList)
		sort.Float64s(floatList)
		sort.Strings(stringList)

		fmt.Println(intList)
		fmt.Println(floatList)
		fmt.Println(stringList)

		sort.Sort(sort.Reverse(sort.IntSlice(intList)))
		sort.Sort(sort.Reverse(sort.Float64Slice(floatList)))
		sort.Sort(sort.Reverse(sort.StringSlice(stringList)))

		fmt.Println(intList)
		fmt.Println(floatList)
		fmt.Println(stringList)
	}

}
