/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"fmt"
	"net/http"
	"time"
)

//复杂的写法，循环不断的访问网址
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


	for{
		go checkLink(<-c,c)
	}
}

func checkLink(link string,c chan string) {
	time.Sleep(2*time.Second)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c<-link
		return
	}

	fmt.Println(link, "is up!")
	c<-link
}
