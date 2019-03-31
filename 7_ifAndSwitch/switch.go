/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"


//1、switch 语句执行的过程自上而下，直到找到case匹配项，匹配项中无需使用break，因为Go语言中的switch默认给每个case自带break，因此匹配成功后不会向下执行其他的case分支，而是跳出整个switch。
//2、变量 var1 可以是任何类型，而 val1 和 val2 则可以是同类型的任意值。类型不被局限于常量或整数，但必须是相同类型或最终结果为相同类型的表达式。
//3、case后的值不能重复。
//4、可以同时测试多个符合条件的值，也就是说case后可以有多个值，这些值之间使用逗号分割，例如：case val1, val2, val3。
//5、Go语言中switch后的表达式可以省略，那么默认是switch  true。
//6、Go语言中的switch case因为自带break，所以匹配某个case后不会自动向下执行其他case，如需贯通后续的case，可以添加fallthrough（中文含义是：贯穿）， 强制执行后面的case分支
//7、fallthrough必须放在case分支的最后一行。如果它出现在中间的某个地方，编译器就会抛出错误（fallthrough  statement  out  of place，含义是fallthrough不在合适的位置）。
//switch 形式1
func main(){

	score :=56
	switch{

	case score >=90:
		fmt.Printf("优秀")
	case score >=80:
		fmt.Printf("良好")
	case score >=70:
		fmt.Printf("中等")
	case score >=60:
		fmt.Printf("及格")
	default:
		fmt.Printf("不及格")
	}

	operate()

}
//switch 形式2
func  operate(){
	a,b,c := 4,2,0

	operate :="*"

	switch operate{
	case "+":
		c=a+b
	case "-":
		c=a-b
	case "*":
		c=a*b

	case "/":
		c=a/b
	default:
		c = -1
	}

	fmt.Println(c)
}


//switch 形式3
func getDaysByMonth() {
	//	定义局部变量：年、月、日
	year := 2008
	month := 12
	days := 0

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		days = 31
	case 4, 6, 9, 11:
		days = 30
	case 2:
		//判断闰年
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			days = 29
		} else {
			days = 28
		}
	default:
		days = -1
	}
	fmt.Printf("%d年%d月的天数为：%d", year , month , days)
}





//switch 形式4
func eval() {
	num1, num2, result := 12, 4, 0
	operation := "+"

	switch operation {
	case "+":
		result = num1 + num2
		//fallthrough
	case "-":
		result = num1 - num2
		//fallthrough
	case "*":
		result = num1 * num2
		//fallthrough
	case "/":
		result = num1 / num2
		//fallthrough
	case "%":
		result = num1 % num2
	default:
		result = -1
	}
	fmt.Println(result)
}