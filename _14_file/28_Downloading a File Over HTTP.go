/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"os"
	"io"
	"log"
	"net/http"
)

func main() {
	// Create output file
	newFile, err := os.Create("devdungeon.html")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	// HTTP GET request devdungeon.com
	url := "http://www.devdungeon.com/archive"
	response, err := http.Get(url)
	defer response.Body.Close()

	// Write bytes from HTTP response to file.
	// response.Body satisfies the reader interface.
	// newFile satisfies the writer interface.
	// That allows us to use io.Copy which accepts
	// any type that implements reader and writer interface
	numBytesWritten, err := io.Copy(newFile, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Downloaded %d byte file.\n", numBytesWritten)
}