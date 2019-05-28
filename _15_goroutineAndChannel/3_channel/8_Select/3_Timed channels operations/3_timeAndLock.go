/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "time"
//下面的程序会发生死锁现象 因为t.reset 与t.stop()不能并发
func toChanTimed(t *time.Timer, ch chan int) {
	t.Reset(1 * time.Second)

	defer func() {

		if !t.Stop() {
			<-t.C
		}
	}()


	select {
	case ch <- 42:
	case <-t.C:
		return
	}

}

func main(){

	t := time.NewTimer(time.Second)

	var ch chan int = make(chan int)

	toChanTimed(t,ch)
}

/*
package main

import "time"

func toChanTimed(t *time.Timer, ch chan int) {
	t.Reset(1 * time.Second)



	select {
	case ch <- 42:
	case <-t.C:
		return
	}

	if !t.Stop() {
		<-t.C
	}
}

func main(){

	t := time.NewTimer(time.Second)

	var ch chan int = make(chan int)

	toChanTimed(t,ch)
}



*/