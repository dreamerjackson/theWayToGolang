/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"log"
	"os"
)

func main() {
	err := os.Remove("test.txt")
	if err != nil {
		log.Fatal(err)
	}
}