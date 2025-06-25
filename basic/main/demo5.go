package main

import "fmt"

func sumFn(x int, y int) int {
	return x + y
}

func subFn(x, y int) int { // 参数类型相同时可只写最后一个
	return x - y
}

func sumFn2(x int, nums ...int) int { // 可变参数, 实际是 切片. 可变参数必须放在参数列表的最后
	var sum = 0
	for _, num := range nums {
		sum += num
	}
	return sum * x
}

func calcFn(x, y int) (int, int) { // 可有多个返回值
	return x + y, x * y
}

func calcFn1(x, y int) (sum, dup int) { // 可给返回值命名并在函数体中使用，最后全部return
	sum = x + y
	dup = x * y
	return
}

func myFn(nums []int) { // 引用传递
	if len(nums) > 0 {
		nums[0] = 999
	}
}

func qSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	i, j, mid := l-1, r+1, l
	for i < j {
		i++
		j--
		for nums[i] < nums[mid] {
			i++
		}
		for nums[j] > nums[mid] {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	qSort(nums, l, j)
	qSort(nums, j+1, r)
}

// 定义函数类型
type twoCalcType func(int, int) int

// 方法作为参数
func myFn1(x, y int, fn twoCalcType) int {
	return 100 + fn(x, y)
}

// 方法作为返回值
func myFn2(s string) twoCalcType {
	switch s {
	case "+":
		return sumFn
	case "-":
		return subFn
	case "*":
		return func(x, y int) int {
			return x * y
		}
	default:
		return nil
	}
}

func demo5() {
	fmt.Println("============= demo 5 ==============")
	{
		arr := []int{1, 2, 3, 4, 5}
		myFn(arr)
		fmt.Println(arr) //[999 2 3 4 5]
		arr1 := []int{4, 6, 2, 5, 7, 4, 7}
		qSort(arr1, 0, len(arr1)-1)
		fmt.Println(arr1)

		// 使用 type关键字 自定义函数类型
		type calcType func(int, int) int
		var c calcType = subFn
		f := subFn
		fmt.Printf("c的类型: %T\n", c) // main.calcType
		fmt.Printf("f的类型: %T\n", f) // func(int, int) int
		fmt.Println(c(10, 2))
		fmt.Println(f(10, 2))

		type myInt int
		var a int = 10
		var b myInt = 10
		fmt.Printf("a的类型: %T\n", a) // int
		fmt.Printf("b的类型: %T\n", b) // main.myInt
		fmt.Println(a + int(b))

		// 方法作为参数和返回值
		fmt.Println(myFn1(1, 2, sumFn)) // 103
		fmt.Println(myFn1(3, 4, func(x, y int) int {
			return x * y
		})) // 112

		f1 := myFn2("+")
		fmt.Println(f1(2, 3)) // 5
		f1 = myFn2("-")
		fmt.Println(f1(2, 3)) // -1
		f1 = myFn2("*")
		fmt.Println(f1(2, 3)) // 6
	}

	{
		// 匿名函数： 可以保存到变量或者直接自执行
		func() {
			fmt.Println("test...")
		}()

		func(x, y int) {
			fmt.Println("hello", x+y)
		}(1, 2)
	}

}
