package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 循环读取文件
func main1() {
	path := "C:\\Users\\薛文朝\\Desktop\\GolangCode\\src\\go_file\\文件编程\\1.txt"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // 关闭文件
	reader := bufio.NewReader(file)
	// 循环按行读取
	for {
		str, err := reader.ReadString('\n')
		fmt.Print(str)
		if err == io.EOF {
			break
		}
	}
}

// 一次性读取文件
func main() {
	path := "C:\\Users\\薛文朝\\Desktop\\GolangCode\\src\\go_file\\文件编程\\1.txt"
	// 使用ioutil一次性读取文件,其中content为[]byte,记得转为string
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}
