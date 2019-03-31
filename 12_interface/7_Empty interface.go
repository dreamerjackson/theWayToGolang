/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"fmt"
)

//空接口  可以存储任何类型的值
type A interface {
}

type Cat struct {
	name string
	age  int
}

type Person struct {
	name string
	sex  string
}

func main() {
	var a1 A = Cat{"Mimi", 1}
	var a2 A = Person{"Steven", "男"}
	var a3 A = "Learn golang with me!"
	var a4 A = 100
	var a5 A = 3.14

	fmt.Println("------------------")

	//1、fmt.println参数就是空接口
	fmt.Println("println的参数就是空接口，可以是任何数据类型", 100, 3.14, Cat{"旺旺", 2})

	//2、定义map。value是任何数据类型
	map1 := make(map[string]interface{})
	map1["name"] = "Daniel"
	map1["age"] = 13
	map1["height"] = 1.71
	fmt.Println(map1)
	fmt.Println("------------------")

	//	3、定义一个切片，其中存储任意数据类型
	slice1 := make([]interface{}, 0, 10)
	slice1 = append(slice1, a1, a2, a3, a4, a5)
	fmt.Println(slice1)

	transInterface(slice1)

	//var cat1 A = Cat{"MiaoMiao" , 3}
	//fmt.Println(cat1.name , cat1.age)

}
//接口对象转型
//接口对象.(type)，配合switch...case语句
func transInterface(s []interface{}) {
	for i := range s {
		fmt.Println("第", i+1 , "个数据：")
		switch t := s[i].(type) {
		case Cat:
			fmt.Printf("\t Cat对象，name属性：%s,age属性：%d \n" , t.name , t.age)
		case Person:
			fmt.Printf("\t Person对象，name属性：%s,sex属性：%s \n" , t.name , t.sex)
		case string:
			fmt.Println("\t string类型" , t)
		case int:
			fmt.Println("\t int类型" , t)
		case float64:
			fmt.Println("\t float64类型" , t)
		}
	}
}

