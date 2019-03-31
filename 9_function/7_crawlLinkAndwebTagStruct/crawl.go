/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"golang.org/x/net/html"
	"fmt"
	"io/ioutil"
	"bytes"
	"os"
)
//golang.org/x/net/html 将网页转换为类似js中的DOM对象
func main(){

	allbyte,_:= ioutil.ReadFile("./functionEmp/index.html")
	//fmt.Println(allbyte)
	doc, err := html.Parse(	bytes.NewReader(allbyte))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	//outline(nil,doc)

	for _,link:= range visit(nil,doc){

		fmt.Println(link)
	}
}

//打印tag
func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}


//打印网页中的链接
func visit(links []string,n*html.Node) []string{

	if n.Type == html.ElementNode && n.Data =="a"{
		for _,a:=range n.Attr{
			if a.Key == "href"{
				links = append(links,a.Val)
			}

		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		 links = visit(links, c)
	}

	return links

}