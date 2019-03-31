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



func main(){

	allbyte,_:= ioutil.ReadFile("./functionEmp/index.html")
	//fmt.Println(allbyte)
	doc, err := html.Parse(	bytes.NewReader(allbyte))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	forEach(doc)
}


var depth int

func forEach( n * html.Node){

	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n",depth*2,"",n.Data)
		depth++
	}


	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEach(c)
	}


	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n",depth*2,"",n.Data)
	}
}