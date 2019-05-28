/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"

const COUNT int = 4

func main() {
	a := [COUNT]string{"abc", "ABC", "123", "一二三"}
	//查看数组的指针的类型和值
	fmt.Printf("%T , %v \n", &a, &a)

	//定义指针数组
	var ptr [COUNT]*string
	fmt.Printf("%T , %v \n", ptr, ptr)

	for i := 0; i < COUNT; i++ {
		//	将数组中每个元素的地址赋值给指针数组的每个元素
		ptr[i] = &a[i]
	}
	fmt.Printf("%T , %v \n", ptr, ptr)

	fmt.Println(ptr[0])

	//	根据指针数组元素的每个地址获取该地址所指向的元素的真实数值
	//  for i:=0; i<COUNT ;i++  {
	//	   fmt.Println(*ptr[i])
	//   }

	for _,value :=range ptr {
		fmt.Println(*value)
	}
}