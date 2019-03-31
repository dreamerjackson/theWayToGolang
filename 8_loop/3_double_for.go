/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"

//循环嵌套
func main(){
	var sum int
	for i:=0;i<5;i++{
		for j:=0;j<3;j++{
			sum = i*j
		}
	}
	fmt.Println(sum)
}