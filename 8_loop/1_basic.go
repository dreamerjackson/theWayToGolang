/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"

func loop1(){
	//形式1
	for i:=0;i<10;i++{
		fmt.Println(i)
	}


}

func loop2(){
	//形式二
	i:=0
	for ;i<10;i++{
		fmt.Println(i)
	}

}

func loop3(){
	//第三种形式
	i:=0
	for ;;i++{

		if(i>20){
			break
		}
		fmt.Println(i)
	}
}

func loop4(){
	//第四种形式
	i:=0
	for ; ; {
		if i>20{
			break
		}
		i++
		fmt.Println(i)
	}
}

func loop5(){
	//第5种形式
	i:=0
	for i<20{
		i++
		fmt.Println(i)

	}
}

func loop6(){
	//第6种形式
	i:=0
	for{
		if i <20{
			i++
			fmt.Println(i)
		}else{
			break
		}


	}
}