/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "time"

//乒乓球模型
func main(){

	var Ball int

	table:= make(chan int)

	go player(table)
	go player(table)
/*
	for i:=0;i<100;i++{
		go player(table)
	}
*/
	table<-Ball

	time.Sleep(1*time.Second)
	<-table

}
func player(table chan int) {
	for{
		ball:=<-table
		ball++
		time.Sleep(100*time.Millisecond)
		table<-ball
	}
}
