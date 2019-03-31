
/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

/*
The ioutil package provides two functions: TempDir() and TempFile(). It is the callers responsibility to delete the temporary items when done. The only benefit these functions provide is that you can pass it an empty string for the directory, and it will automatically create the item in the system's default temporary folder (/tmp on Linux). Since os.TempDir() function that will return the defauly system temporary directory.
*/

package main

import (
"os"
"io/ioutil"
"log"
"fmt"
)

func main() {
	// Create a temp dir in the system default temp folder
	tempDirPath, err := ioutil.TempDir("", "myTempDir")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Temp dir created:", tempDirPath)

	// Create a file in new temp directory
	tempFile, err := ioutil.TempFile(tempDirPath, "myTempFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Temp file created:", tempFile.Name())

	// ... do something with temp file/dir ...

	// Close file
	err = tempFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Delete the resources we created
	err = os.Remove(tempFile.Name())
	if err != nil {
		log.Fatal(err)
	}
	err = os.Remove(tempDirPath)
	if err != nil {
		log.Fatal(err)
	}
}