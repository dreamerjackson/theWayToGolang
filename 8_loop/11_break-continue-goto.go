/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"fmt"
)

func main() {
	//break，continue的区别
	breakContinue()

	//	输出1-50之间所有不包含4的数字（continue实现）
	eludeFour()

	eludeFourGoto()

	//输出1-100的素数（借助goto跳转）
	printPrimeNumber()
}

/*
break，continue的区别：
break语句将无条件跳出并结束当前的循环， 然后执行循环体后的语句;
continue语句是跳过当前的循环， 而开始执行下一次循环。
 */
func breakContinue() {
	fmt.Println("\n1、break终止循环")
	//	1、break终止循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Print(i)
	}
	//	结果是01234

	fmt.Println("\n2、continue跳过某次循环")
	//	2、continue跳过某次循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Print(i)
	}
	//	结果是012346789
}

//输出1-50之间所有不包含4的数字（个位和十位）
func eludeFour() {
	fmt.Println("\n输出1-50之间所有不包含4的数字")
	//定义局部变量
	num := 0
	for num < 50 {
		num++
		if num%10 == 4 || num/10%10 == 4 {
			continue
		}
		fmt.Printf("%d\t", num)
	}
}

//输出1-50之间不包含4的数（goto实现）
func eludeFourGoto() {
	fmt.Println("\n输出1-50之间所有不包含4的数字。goto实现")
	num := 0
LOOP:
	for num < 50 {
		num++
		if num%10 == 4 || num/10%10 == 4 {
			goto LOOP
		}
		fmt.Printf("%d\t", num)
	}
}

//输出1-100的素数（借助goto跳转）
func printPrimeNumber() {
	fmt.Println("\n1-100的素数（借助goto跳转）")
	num := 0
LOOP:
	for num < 100 {
		num++
		for i := 2; i < num; i++ {
			if num%i == 0 {
				goto LOOP
			}
		}
		fmt.Printf("%d\t", num)
	}

}
