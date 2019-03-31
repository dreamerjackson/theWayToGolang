/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"

//计算1-100之间的和：
func sum(){

	result:=0

	for i:=0;i<=100;i++{
		fmt.Printf("result:%d   i:%d\n",result,i)
		result +=i
	}

	fmt.Println(result)
}


//计算1 - 100之间所有的奇数的和
func sum2(){

	result:=0

	for i:=0;i<=100;i++{
		if(i %2 ==0){
			result +=i
		}
	}
}