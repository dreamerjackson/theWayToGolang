
/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)
var g_pp = make([]*[]*byte,10)

//初始化构建内存模型
func init(){
	fmt.Println("初始化开始")
	file, err := os.Open("./lab-QQRecursive tree/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str:= scanner.Text()

		if len(str) <50{//数据已经整理过，最多50位。
			qq:= getQQ(str)
			//fmt.Println(qq)
			if len(qq)==10  && isAllNum(qq){
				//递归树，将模型构建完毕。
				assign(&g_pp,11,qq,str)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("初始化结束")
}



func main() {
	search()
}