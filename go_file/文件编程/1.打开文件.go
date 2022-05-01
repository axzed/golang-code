package main

import (
	"fmt"
	"os"
)

func main1() {
	path := "C:\\Users\\薛文朝\\Desktop\\GolangCode\\src\\go_file\\文件编程"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file)
	defer file.Close() // 关闭文件
}
