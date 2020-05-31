package main

import "fmt"

// https://gobyexample.com/values
func main() {
	fmt.Println("a" + "b")           // 字符串拼接
	fmt.Println("1+1=", 1+1)         // 加法
	fmt.Println("7.0/3.0=", 7.0/3.0) // 除法
	fmt.Println(true && false)       // 与
	fmt.Println(true || false)       // 或
	fmt.Println(!true)               // 非
}

// ============== output ==============
// ab
// 1+1= 2
// 7.0/3.0= 2.3333333333333335
// false
// true
// false
