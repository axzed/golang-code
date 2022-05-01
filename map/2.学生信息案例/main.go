package main

import "fmt"

/*
课堂案例: 演示一个key-value的value是map的案例
例如: 要存放三个学生的信息,每个学生有name和sex信息
思路: map[string]map[string]string
*/
func main() {
	studentMap := make(map[string]map[string]string, 3)

	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["address"] = "长安区"

	studentMap["stu02"] = make(map[string]string)
	studentMap["stu02"]["name"] = "mary"
	studentMap["stu02"]["sex"] = "女"
	studentMap["stu02"]["address"] = "雁塔区"

	fmt.Println(studentMap)
}
