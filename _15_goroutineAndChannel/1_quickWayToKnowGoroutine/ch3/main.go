/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"fmt"
	"net/http"
)

//goroutine
func main() {
	links := []string{
		"http://www.baidu.com",
		"http://www.jd.com/‎",
		"https://www.taobao.com/",
		"https://www.163.com/",
		"https://www.sohu.com/",
	}

	c:=make(chan string)
	for _, link := range links {
		go checkLink(link,c)
	}
	fmt.Println(<-c)
}

func checkLink(link string,c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c<-"might be down"
		return
	}

	fmt.Println(link, "is up!")
	c<-"is up"
}

