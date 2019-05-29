package main

import (
	"fmt"
)

func main() {
	// const norm1 = 2
	// const norm3 = 4
	// const norm2 string = "asd";
	// var area int
	// area = norm1 * norm3
	// fmt.Printf("面积是 : %d", area)

	// const(
	// 	read = 1
	// 	face = 2
	// 	less = 3
	// )

	// const (
	// 	a = 1
	// 	b = 2
	// 	c = 3
	// )
	// fmt.Println(a,b,c)

		const(
			ali = iota
			tencent = iota
			other = "我的"
			baidu = iota
		)
		fmt.Println(ali, tencent, other, baidu)

	// const (
	// 	a = iota   //0
	// 	b          //1
	// 	c          //2
	// 	d = "ha"   //独立值，iota += 1
	// 	e          //"ha"   iota += 1
	// 	f = 100    //iota +=1
	// 	g          //100  iota +=1
	// 	h = iota   //7,恢复计数
	// 	i          //8
	// )
	// fmt.Println(a,b,c,d,e,f,g,h,i)


}
