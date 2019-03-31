/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"

//if 判断学生成绩
func scoreTest(){

	var  score =80

	if score >=90{
		fmt.Printf("优秀")
	}

	if score >=80 && score <90{
		fmt.Printf("良好")
	}
	if score >=70 && score <80{
		fmt.Printf("中等")
	}

	if score >=60 && score <70 {
		fmt.Printf("及格")
	}

	if score <60{
		fmt.Printf("不及格")
	}
}

//if else 语句判断
//判断学生成绩
func scoreTest3(score int) {
	if score >= 90 {
		fmt.Printf("优秀")
	} else if score >= 80 {
		fmt.Printf("良好")
	} else if score >= 70 {
		fmt.Printf("中等")
	} else if score >= 60 {
		fmt.Printf("及格")
	} else {
		fmt.Printf("不及格")
	}
}