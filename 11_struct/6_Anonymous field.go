/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"
//没有变量名
type User struct {
	//name string
	//sex byte
	//age int8
	//height float64
	//weight float64
	string
	byte
	int8
	float64

}

func main() {
	//	实例化结构体
	user:= User{"Steven" , 'm' , 35 , 177.5}
	fmt.Println(user)
	//如果想依次输出姓名、年龄、身高、性别
	fmt.Printf("姓名：%s \n" , user.string)
	fmt.Printf("身高：%.2f \n" , user.float64)
	fmt.Printf("性别：%c \n" , user.byte)
	fmt.Printf("年龄：%d \n" , user.int8)
}
