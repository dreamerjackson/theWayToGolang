/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"


//蛇形矩阵
func main(){

	var arr [8][8]int
	N:= 8
	data:=1
	for i:=0;i<N;i++{
		for j:=0;j<N;j++{
			fmt.Printf("%4d",arr[i][j])
		}
		fmt.Println()
	}

	i,j,k:=0,0,0
	for ;k<(8+1)/2 ;k++{

		for j< N-k{
			arr[i][j] = data
			data++
			j++
		}
		j--
		i++

		for i< N-k{
			arr[i][j] = data
			data++
			i++
		}
		i--
		j--

		for j > k-1{
			arr[i][j] = data
			j--;
			data++
		}
		j++
		i--

		for i>k{

			arr[i][j] = data
			i--
			data++
		}

		i++
		j++


	}



	for i:=0;i<N;i++{
		for j:=0;j<N;j++{
			fmt.Printf("%4d",arr[i][j])
		}
		fmt.Println()
	}
}