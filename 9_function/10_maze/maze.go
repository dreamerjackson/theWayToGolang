/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"fmt"
	"bufio"
	"os"
)

var arr  = [][]int{
	{0,2,0,0,0,0},
	{0,0,2,0,0,0},
	{0,2,0,2,0,0},
	{0,0,2,0,2,0},
	{0,0,2,0,0,0},
	{0,0,0,0,2,0},
}

//打印迷宫
func printMaze(arr [][]int){

	for _,row :=  range arr{
		for _,col:= range row{

			fmt.Printf("%3d",col)

		}
		fmt.Println()

	}
	fmt.Println()

}


//AI走法
func RunAI(x int,y int) bool{

	mazelen:= len(arr)
	//是否越界
	if (x<0 || y<0 || x>=mazelen || y >=mazelen){
		return false
	}
	//是否障碍物
	if arr[x][y] > 0{
		return false
	}

	if x==mazelen-1 && y == mazelen-1{
		arr[x][y] = 1
		return true
	}

	arr[x][y] = 1

	if(RunAI(x,y+1) || RunAI(x+1,y) || RunAI(x,y-1) || RunAI(x-1,y)){
		return true
	}


	arr[x][y] = 0

	return false

}


func main(){

	printMaze(arr)

	 complete := RunAI(0,0)

	 if complete{
		 printMaze(arr)
	 }else{
	 	fmt.Println("无解")
	 }

}


// 键盘走法
func keyboard(){

	  x,y:=0,0
	  arr[x][y] = 1
	  mazelen := len(arr)
		printMaze(arr)
	input := bufio.NewScanner(os.Stdin)


	for  input.Scan(){
		switch input.Text(){
		case "a":
			arr[x][y] = 0
			if y==0 || arr[x][y-1]>1{
				continue
			}

			y -=1
			arr[x][y] = 1
		case "w":
			arr[x][y] = 0
			if x==0 || arr[x-1][y] > 1{
				continue
			}

			x -=1
			arr[x][y] = 1
		case "s":
			arr[x][y] = 0
			if x==mazelen-1 || arr[x+1][y] >2{
				continue
			}

			x+=1
			arr[x][y] = 1
		case "d":
			arr[x][y] = 0
			if y==mazelen-1  || arr[x][y+1] > 1{
				continue
			}

			y+=1
			arr[x][y] = 1
		default:



		}

		printMaze(arr)

	}

}
