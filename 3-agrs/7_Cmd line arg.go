/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"os"
	"fmt"
	"strings"
)

//命令行参数

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

//方式2
func main2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

//方式3

func main3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

/*
$ go build main.go
$ ./main a b c d f
a b c d f
*/