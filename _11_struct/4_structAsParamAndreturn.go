/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"

type Point struct{ X, Y int }


func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

func main(){

	fmt.Println(Scale(Point{1, 2}, 5)) // "{5 10}"

}


//如果要在函数内部修改结构体成员的话，用指针传入是必须的；
// 因为在Go语言中，所有的函数参数都是值拷贝传入的，函数参数将不再是函数调用时的原始变量。

func changeX(e *Point) {
   e.X = 100
}
//因为结构体通常通过指针处理，可以用下面的写法来创建并初始化一个结构体变量，并返回结构体的地址：
//pp := &Point{1, 2}



//案例2

type Flower struct {
	name, color string
}

func main2() {
	//1、结构体作为参数的用法
	f1 := Flower{"玫瑰", "红"}
	fmt.Printf("f1: %T , %v , %p \n" , f1 , f1 , &f1)
	fmt.Println("----------------------")

	//将结构体对象作为参数
	changeInfo1(f1)
	fmt.Printf("f1: %T , %v , %p \n" , f1 , f1 , &f1)
	fmt.Println("----------------------")

	//	将结构体指针作为参数
	changeInfo2(&f1)
	fmt.Printf("f1: %T , %v , %p \n" , f1 , f1 , &f1)
	fmt.Println("----------------------")

	//2、结构体作为返回值的用法
	//结构体对象作为返回值
	f2 := getFlower1()
	f3 := getFlower1()
	fmt.Println("更改前" , f2 , f3)
	f2.name = "杏花"
	fmt.Println("更改后" , f2 , f3)

	//结构体指针作为返回值
	f4 := getFlower2()
	f5 := getFlower2()
	fmt.Println("更改前" , f4 , f5)
	f4.name = "桃花"
	fmt.Println("更改后" , f4 , f5)
}

//返回结构体对象
func getFlower1() (f Flower){
	f = Flower{"牡丹", "白"}
	fmt.Printf("函数getFlower1内f: %T , %v , %p \n" , f , f , &f)
	return
}

//返回结构体指针
func getFlower2() (f *Flower){
	//f = &Flower{"芙蓉", "红"}
	temp := Flower{"芙蓉", "红"}
	fmt.Printf("函数getFlower2内temp: %T , %v , %p \n" , temp , temp , &temp)
	f = &temp
	fmt.Printf("函数getFlower2内f: %T , %v , %p , %p \n" , f , f , f , &f)
	return
}

//传结构体对象
func changeInfo1(f Flower) {
	f.name = "月季"
	f.color = "粉"
	fmt.Printf("函数changeInfo1内f: %T , %v , %p \n" , f , f , &f)
}

//传结构体指针
func changeInfo2(f *Flower) {
	f.name = "蔷薇"
	f.color = "紫"
	fmt.Printf("函数changeInfo2内f: %T , %v , %p , %p \n" , f , f , f , &f)
}
