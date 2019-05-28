

/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

/*
A typical file is just a pointer to a place on the hard disk called an inode. A hard link creates a new pointer to the same place. A file will only be deleted from disk after all links are removed. Hard links only work on the same file system. A hard link is what you might consider a 'normal' link.

A symbolic link, or soft link, is a little different, it does not point directly to a place on the disk. Symlinks only reference other files by name. They can point to files on different filesystems. Not all systems support symlinks.
*/

package main

import (
"os"
"log"
"fmt"
)

func main() {
	// Create a hard link
	// You will have two file names that point to the same contents
	// Changing the contents of one will change the other
	// Deleting/renaming one will not affect the other
	err := os.Link("original.txt", "original_also.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("creating sym")
	// Create a symlink
	err = os.Symlink("original.txt", "original_sym.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Lstat will return file info, but if it is actually
	// a symlink, it will return info about the symlink.
	// It will not follow the link and give information
	// about the real file
	// Symlinks do not work in Windows
	fileInfo, err := os.Lstat("original_sym.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Link info: %+v", fileInfo)

	// Change ownership of a symlink only
	// and not the file it points to
	err = os.Lchown("original_sym.txt", os.Getuid(), os.Getgid())
	if err != nil {
		log.Fatal(err)
	}
}