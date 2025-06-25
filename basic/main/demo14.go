package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readByByteStream() {
	file, err := os.Open("./abc.txt") // 绝对路径或者相对路径
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	// 使用os.Open() 读取文件
	var strSlice []byte
	tempSlice := make([]byte, 32)

	for {
		byteCount, err := file.Read(tempSlice)
		if err == io.EOF {
			fmt.Printf("读取完毕, 读取到的全部内容为: \n%s\n", string(strSlice))
			break
		}
		if err != nil {
			fmt.Println("读取失败")
		}
		fmt.Printf("此次读取到了%v个字节, 内容为:\n%s\n", byteCount, string(tempSlice[:byteCount]))
		strSlice = append(strSlice, tempSlice[:byteCount]...)
	}
}

func readByBufio() {
	file, err := os.Open("./abc.txt") // 绝对路径或者相对路径
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	//使用 bufio 读取文件
	reader := bufio.NewReader(file)
	var fileStr string

	for {
		str, err := reader.ReadString('\n') //参数为读取字符的分隔符,这里表示一次读取一行
		if err == io.EOF {
			fileStr += str // 注意：使用bufio读取数据时，读取到EOF后，最后一次读取的变量中可能仍要内容，不要忽略掉
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fileStr += str
	}

	fmt.Printf("bufio读取完成 内容为: \n%s\n", fileStr)
}

func readByIoutil() {
	// 通过ioutil读取文件，最为简单，已经封装好了打开关闭文件的方法，只需一行代码即可读取
	// 适合读取小文件
	byteSlice, err := ioutil.ReadFile("./abc.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ioutil读取完毕, 内容为:")
	fmt.Println(string(byteSlice))

}

func wirteByByte() {

	file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	// for i := 0; i < 10; i++ {
	// 	file.WriteString("字符串数据" + strconv.Itoa(i) + "\r\n")
	// }

	str := "直接写入的字符串数据1234567890qwertyuiop"
	file.Write([]byte(str)) // string转换成切片写入

	fmt.Println("byte写入完成, 文件已关闭.")
}

func wirteByBufio() {
	file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	defer file.Close() //

	if err != nil {
		fmt.Println(err)
		return
	}

	write := bufio.NewWriter(file)
	write.WriteString("你好golang") //将string写入缓存
	write.Flush()                 // 将缓存数据写入文件

	fmt.Println("bufio写入完成, 文件已关闭.")
}

func wirteByIoutil() {
	// 文件打开关闭已经封装好了
	str := "hello golang!"
	err := ioutil.WriteFile("./test.txt", []byte(str), 0666) // 默认是清楚重写模式
	if err != nil {
		fmt.Println("write file failed! err :", err)
		return
	}
	fmt.Println("ioutil写入完成, 文件已关闭.")

}

func demo14() {

	// 文件操作

	// 三种读取方式
	// readByByteStream()
	// readByBufio()
	// readByIoutil()

	// 三种写入方法
	// wirteByByte()
	// wirteByBufio()
	// wirteByIoutil()

	// 可以通过 ioutil或者流的文件读写方式实现文件复制

	// 目录操作
	//创建目录
	// 如果目录已存在则不做任何操作
	// if err := os.Mkdir("./123", 0666); err != nil {
	// 	fmt.Println(err)
	// }

	// 创建多级目录
	// if err := os.MkdirAll("./dir1/dir2/dir3", 0666); err != nil {
	// 	fmt.Println(err)
	// }

	// os.Remove(文件或者目录) 删除文件或者目录
	// os.RemoveAll(目录) 递归删除文件或目录
	os.Remove("./test.txt")
	os.RemoveAll("./dir1")

	// os.Rename(文件或者目录1, 文件或者目录2) 文件重命名, 只能同盘操作

}
