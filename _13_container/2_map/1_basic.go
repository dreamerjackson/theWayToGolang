/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"

//1、声明map的方式1
var map1 map[string]string


//2、声明map的方式2
var map2 = make(map[string]string)

//3、map中key可以是：int、float、bool、string、数组
//	一定不可以是：切片、函数、map
var m1 map[int]string
var m2 map[float64]string
var m3 map[bool]string
var m4 map[string]string


func main() {
	//4、声明时同时初始化
	var country = map[string]string{
		"China":  "Beijing",
		"Japan":  "Tokyo",
		"India":  "New Delhi",
		"France": "Paris",
		"Italy":  "Rome",
	}
	fmt.Println(country)

	//5、短变量声明初始化方式
	rating := map[string]float64{"c": 5, "Go": 4.5, "Python": 4.5, "C++": 3}
	fmt.Println(rating)

	//2、创建map后再赋值
	countryMap := make(map[string]string)
	countryMap["China"] = "Beijing"
	countryMap["Japan"] = "Tokyo"
	countryMap["India"] = "New Delhi"
	countryMap["France"] = "Paris"
	countryMap["Italy"] = "Rome"

	fmt.Println(countryMap)
}