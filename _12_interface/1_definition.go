/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"

type Phone interface {
	call()
}

type AndroidPhone struct {
}

type IPhone struct {
}

func (a AndroidPhone) call() {
	fmt.Println("我是安卓手机，可以打电话了")
}

func (i IPhone) call() {
	fmt.Println("我是苹果手机，可以打电话了")
}

func main() {
	//	定义接口类型的变量 只要实现了接口中的函数，就可以被接口所接受，类似于多态
	var phone Phone
	phone = new(AndroidPhone)
	phone = AndroidPhone{}
	fmt.Printf("%T , %v , %p \n" , phone , phone , &phone)
	phone.call()

	phone = new(IPhone)
	phone = IPhone{}
	fmt.Printf("%T , %v , %p \n" , phone , phone , &phone)
	phone.call()
}