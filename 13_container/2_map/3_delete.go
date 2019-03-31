/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"

func main() {
	//1、声明并初始化一个map
	map1 := map[string]string {
		"element":"div",
		"width" :"100px",
		"height":"200px",
		"border":"solid",
		"background":"none",
	}

	//2、根据key删除map中的某个元素
	fmt.Println("删除前：",map1)
	if _,ok := map1["background"]; ok {
		delete(map1 , "background")
	}
	fmt.Println("删除后：",map1)

	//3、清空map
	//map1 = map[string]string{}
	map1 = make(map[string]string)
	fmt.Println("清空后：",map1)
}