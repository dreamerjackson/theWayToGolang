




/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */




/*
Scanner is part of the bufio package. It is useful for stepping through files at specific delimiters. Commonly, the newline character is used as the delimiter to break up a file by lines. In a CSV file, commas would be the delimiter. The os.File can be wrapped in a bufio.Scanner just like a buffered reader. We call Scan() to read up to the next delimiter, and then use Text() or Bytes() to get the data that was read.

The delimiter is not just a simple byte or character. There is actually a special function you have to implement that will determine where the next delimiter is, how far forward to advance the pointer, and what data to return. If no custom SplitFunc is provided, it defaults to ScanLines which will split at every newline character. Other split functions included in bufio are ScanRunes, and ScanWords.

// To define your own split function, match this fingerprint
type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)

// Returning (0, nil, nil) will tell the scanner
// to scan again, but with a bigger buffer because
// it wasn't enough data to reach the delimiter
In the next example, a bufio.Scanner is created from the file, and then we scan read the file word by word.

*/

package main

import (
"os"
"log"
"fmt"
"bufio"
)

func main() {
	// Open file and create scanner on top of it
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	// Default scanner is bufio.ScanLines. Lets use ScanWords.
	// Could also use a custom function of SplitFunc type
	scanner.Split(bufio.ScanWords)

	// Scan for next token.
	success := scanner.Scan()
	if success == false {
		// False on error or EOF. Check error
		err = scanner.Err()
		if err == nil {
			log.Println("Scan completed and reached EOF")
		} else {
			log.Fatal(err)
		}
	}

	// Get data from scan with Bytes() or Text()
	fmt.Println("First word found:", scanner.Text())

	// Call scanner.Scan() again to find next token
}