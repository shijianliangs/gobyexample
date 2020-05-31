package main

import "fmt"

// https://gobyexample.com/variables
func main() {
	var a = "hello"
	fmt.Println(a)
	var b, c int = 1, 2
	fmt.Println(b, c)
	var d = false
	fmt.Println(d)
	var e int // int变量的默认值是0
	fmt.Println(e)
	f := "world" // 可以不使用var，使用:=的形式进行定义变量
	fmt.Println(f)
}

// ============== output ==============
// hello
// 1 2
// false
// 0
// world
