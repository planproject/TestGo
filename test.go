package main

import (
	"fmt"
)

func main() {
	for i:=1;i<=3;i++ {
		if i>=2 {
			fmt.Print("当前的i值为：")
			fmt.Print(i)
			fmt.Print("\n")
			continue
		}
		fmt.Print("慕课网\n")
	}

}
