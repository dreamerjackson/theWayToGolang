/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main


/*
	type error interface {
		Error() string
	}
	关于errors.New()
	package errors
	func New(text string) error { return &errorString{text} }
	type errorString struct { text string }
	func (e *errorString) Error() string { return e.text }

syscall包中的error
	package syscall
	type Errno uintptr // operating system error code
	var errors = [...]string{
		1: "operation not permitted", // EPERM
		2: "no such file or directory", // ENOENT
	   3: "no such process", // ESRCH
	// ...
	}
	func (e Errno) Error() string {
		if 0 <= int(e) && int(e) < len(errors) {
		   return errors[e]
		}
	return fmt.Sprintf("errno %d", e)
	}
*/
