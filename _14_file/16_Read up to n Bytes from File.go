


/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */


/*
The os.File type provides a couple basic functions. The io, ioutil, and bufio packages provided additional functions for working with files.
*/

package main

import (
	"os"
	"log"
)

func main() {
	// Open file for reading
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read up to len(b) bytes from the File
	// Zero bytes written means end of file
	// End of file returns error type io.EOF
	byteSlice := make([]byte, 16)
	bytesRead, err := file.Read(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of bytes read: %d\n", bytesRead)
	log.Printf("Data read: %s\n", byteSlice)
}