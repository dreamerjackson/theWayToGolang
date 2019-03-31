/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"

func main(){


	//变量的地址
	a :=10
	fmt.Printf("a变量的地址为：%#X\n",&a)


	//声明
	var p *int

	//空指针
	if p==nil{
		fmt.Println("p为空指针")
	}

	//通过指针获取值
	p = &a
	fmt.Printf("p的类型为%T, p的值为：%v,p指向的int的值为：%v,a的值为：%d\n",p,p,*p,a)


	//通过指针修改值
	*p = 99
	fmt.Printf("p的类型为%T, p的值为：%v,p指向的int的值为：%v,a的值为：%d\n",p,p,*p,a)

}