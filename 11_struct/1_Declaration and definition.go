/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"

//定义Teacher结构体
type Teacher struct {
	name string
	age  int8
	sex  byte
}

func main() {
	//1、var声明方式实例化结构体，初始化方式为：对象.属性=值
	var t1 Teacher
	fmt.Println(t1)
	fmt.Printf("t1:%T , %v , %q \n", t1, t1, t1)
	//if t1 == nil {
	//	fmt.Println()
	//}
	t1.name = "Steven"
	t1.age = 35
	t1.sex = 1
	fmt.Println(t1)
	fmt.Println("-------------------")

	//2、变量简短声明格式实例化结构体，初始化方式为：对象.属性=值
	t2 := Teacher{}
	t2.name = "David"
	t2.age = 30
	t2.sex = 1
	fmt.Println(t2)
	fmt.Println("-------------------")

	//3、变量简短声明格式实例化结构体，声明时初始化。初始化方式为：属性:值 。属性:值可以同行，也可以换行。（类似map的用法）
	t3 := Teacher{
		name: "Josh",
		age:  28,
		sex:  1,
	}
	t3 = Teacher{name: "Josh2", age: 27, sex: 1}
	fmt.Println(t3)
	fmt.Println("-------------------")

	//4、变量简短声明格式实例化结构体，声明时初始化，不写属性名，按属性顺序只写属性值
	t4 := Teacher{"Ruby", 30, 0}
	fmt.Println(t4)
	fmt.Println("-------------------")

	//5、创建指针类型的结构体
	t5 := new(Teacher)
	fmt.Printf("t5:%T , %v , %p \n", t5, t5, t5)
	//(*t5).name = "Running"
	//(*t5).age = 31
	//(*t5).sex = 0

	//语法简写形式——语法糖
	t5.name = "Running2"
	t5.age = 31
	t5.sex = 0
	fmt.Println(t5)
	fmt.Println("-------------------")
}