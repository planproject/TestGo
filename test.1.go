package main

import (
	"fmt"
)

func main() {
	// fmt.Println("satan")
	// fmt.Println("fly with me")
	// fmt.Println("zahuishier");fmt.Println("ok");
	// fmt.Println("fade" + "sad")

	// var str string = "fake"
	// fmt.Println(str)

	// var num1, num2 int = 1, 2
	// fmt.Println(num1, num2)

	// 声明一个变量并初始化
	// var a = "RUNOOB"
	// fmt.Println(a)

	// // 没有初始化就为零值
	// var b int
	// fmt.Println(b)

	// // bool 零值为 false
	// var c bool
	// fmt.Println(c)
	// var s string
	// fmt.Printf("%q %v\n", s, c)
	//%v 为默认模式

	//:= 声明新的变量，如左已声明则会编译错误
	// faka, lala := 1, 2	//多变量声明
	// fmt.Println(faka, lala)


	//以下因式分解关键字的写法多用于声明全局变量
	var (
		name1 string
		name2 int
	)
	name1 = "asd"
	name2 = 123
	fmt.Println(name1, name2)

	//这种不带声明格式的只能在函数体中出现
	// g, h := 123, "hello"

	//不可以再次对于相同名称的变量使用初始化声明
	//局部定义了一个变量却没使用，也会报编译错误
	//全局变量是允许声明但不使用,同一类型的多个变量可以声明在同一行
	// var a, b, c int
	//语言常量 TODO


}
