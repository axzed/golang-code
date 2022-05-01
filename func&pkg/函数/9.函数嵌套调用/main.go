package main

import "fmt"

// 发送邮箱信息
func SendMsg() {

}

// 完成用户信息校验
func CheckInfo(userName string, userPwd string, userEmail string) {

}

func Register() {
	// 接收用户在网页中填写的信息

	// 对用户信息校验
	CheckInfo("xuewenchao", "123", "xuewenchao@qq.com")

	// 完成用户信息保存

	// 发送电子邮件
	SendMsg()
	fmt.Println("用户注册成功")
}

// 不定参数函数嵌套调用
func Test1(args ...int) {
	for _, data := range args {
		fmt.Println(data)
	}
}

func Test(args ...int) {
	// 全部将数据传递给了Test1
	// Test1(args...)
	
	// 将编号为2(包含2)以后的所有数据传递给了Test1,编号从0开始
	// Test1(args[2:]...)

	// 将编号0至编号2(不包含2)之前的数据进行传递
	Test1(args[:2]...)
}

func main() {
	// Register()
	Test(3, 5, 9, 13)
}
