/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"fmt"
)

var lines = 9

func main() {
	//	打印矩形
	printRectangle()

	//打印左下直角三角形
	printRightTriangleLB()

	//打印左上直角三角形
	printRightTriangleLT()

	//	打印右下直角三角形
	printRightTriangleRB()

	//	打印右上直角三角形
	printRightTriangleRT()

	//	打印等腰三角形
	printEqualTriangle()

	//	打印九九乘法表
	multiple99()
}

//1、打印矩形
func printRectangle() {
	fmt.Println("\n打印矩形")
	for i := 1; i <= lines; i++ {
		for j := 1; j <= lines; j++ {
			fmt.Print("❤ ")
		}
		fmt.Println()
	}
}

//2、打印左下直角三角形
func printRightTriangleLB() {
	fmt.Println("\n打印左下直角三角形")
	for i := 1; i <= lines; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("❤ ")
		}
		fmt.Println()
	}
}

//3、打印左上直角三角形
func printRightTriangleLT() {
	fmt.Println("\n打印左上直角三角形")
	for i := 1; i <= lines; i++ {
		for j := lines; j >= i; j-- {
			fmt.Print("❤ ")
		}
		fmt.Println()
	}
}

//4、打印右下直角三角形
func printRightTriangleRB() {
	fmt.Println("\n打印右下直角三角形")
	for i := 1; i <= lines; i++ {
		//打印空格
		for m := lines; m >= i; m-- {
			fmt.Print("  ")
		}
		//打印三角形
		for j := 1; j <= i; j++ {
			fmt.Print("❤ ")
		}
		fmt.Println()
	}
}

//5、打印右上直角三角形
func printRightTriangleRT() {
	fmt.Println("\n打印右上直角三角形")
	for i := 1; i <= lines; i++ {
		//打印空格
		for m := 1; m <= i; m++ {
			fmt.Print("  ")
		}
		//打印三角形
		for j := lines; j >= i; j-- {
			fmt.Print("❤ ")
		}
		fmt.Println()
	}
}

//6、打印等腰三角形
func printEqualTriangle() {
	fmt.Println("\n打印等腰三角形")
	for i := 1; i <= lines; i++ {
		//打印空格
		for m := lines; m >= i; m-- {
			fmt.Print("  ")
		}

		//	打印三角形
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("❤ ")
		}
		fmt.Println()
	}
}

//7、打印九九乘法表
func multiple99() {
	fmt.Println("\n打印九九乘法表")
	for i := 1; i <= lines; i++ {//i控制行，乘法表中的第二个数字
		for j := 1; j <= i; j++ {//j控制列，乘法表中的第一个数字
			fmt.Printf("%d*%d=%2d ", j, i, i*j)
		}
		fmt.Println()
	}
}
