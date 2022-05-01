package main

import "fmt"

func main() {
	fmt.Println("日照香炉生紫烟")
	fmt.Println("遥看瀑布挂前川")
	fmt.Println("飞流直下三千尺")
	goto GAMEOVER

	LASTWORD:
	fmt.Println("薛文朝")
	return 

	GAMEOVER:
	fmt.Println("疑似银河落九天")
	goto LASTWORD
}
