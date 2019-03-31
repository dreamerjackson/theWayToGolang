/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"io/ioutil"
	"log"
)

func main() {
	err := ioutil.WriteFile("test.txt", []byte("Hi\n"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
