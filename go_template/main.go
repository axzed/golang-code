package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}
	u1 := User{
		Name:   "薛文朝",
		Gender: "男",
		Age:    18,
	}
	m1 := map[string]interface{}{
		"name":   "二号机",
		"gender": "女",
		"age":    29,
	}
	// 渲染模板
	t.Execute(w, map[string]interface{}{
		"u1": u1,
		"m1": m1,
	})

}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
