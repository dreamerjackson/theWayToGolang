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
	// Test write permissions. It is possible the file
	// does not exist and that will return a different
	// error that can be checked with os.IsNotExist(err)
	file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error: Write permission denied.")
		}
	}
	file.Close()

	// Test read permissions
	file, err = os.OpenFile("test.txt", os.O_RDONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error: Read permission denied.")
		}
	}
	file.Close()
}

//https://www.linux.com/learn/understanding-linux-file-permissions