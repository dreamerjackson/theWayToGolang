/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"



//Appennd   Delete and Copy
func main() {
	fmt.Println("1、--------------")
	//var numbers []int
	numbers := make([]int, 0, 20)


	//append一个元素
	numbers = append(numbers, 0)
	printSlice("numbers:", numbers) //[0]

	//append多个元素
	numbers = append(numbers, 1, 2, 3, 4, 5, 6, 7)
	printSlice("numbers:", numbers) //[0 1 2 3 4 5 6 7]

	//append添加切片
	s1 := []int{100, 200, 300, 400, 500, 600, 700}
	numbers = append(numbers, s1...) //[0 1 2 3 4 5 6 7 100 200 300 400 500 600 700]
	printSlice("numbers:", numbers)

	fmt.Println("2、--------------")
	//	切片删除
	//	删除第一个元素
	numbers = numbers[1:]
	printSlice("numbers:", numbers) //[ 1 2 3 4 5 6 7 100 200 300 400 500 600 700]

	//	删除最后一个
	numbers = numbers[:len(numbers)-1]
	printSlice("numbers:", numbers) //[ 1 2 3 4 5 6 7 100 200 300 400 500 600]

	//删除中间一个元素
	a := int(len(numbers) / 2)
	fmt.Println("中间下标：", a)
	numbers = append(numbers[:a], numbers[a+1:]...)
	printSlice("numbers:", numbers) //[1 2 3 4 5 6 100 200 300 400 500 600]




	fmt.Println("3、--------------")
	//创建目标切片
	numbers1 := make([]int, len(numbers), cap(numbers)*2)

	//	将numbers的元素拷贝到numbers1中
	//毫无疑问，创建新的目标切片就会有新的指向的数组。数组的copy是对于不同的数组的值的拷贝
	count := copy(numbers1, numbers)
	fmt.Println("拷贝的个数：", count)
	printSlice("numbers1:", numbers1)

	//拷贝的两个切片是否有关联
	numbers[0] = 99
	numbers1[len(numbers1)-1] = 100

	printSlice("numbers", numbers)
	printSlice("numbers1", numbers1)

}

func printSlice(name string, x []int) {
	fmt.Print(name, "\t")
	fmt.Printf("地址：%p \t len=%d \t cap=%d \t value=%v \n", x, len(x), cap(x), x)
}


