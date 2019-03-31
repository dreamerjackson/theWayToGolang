/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"time"
	"fmt"
)
//协程中的竞争锁
func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
			}
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println(a)
}


/*
$ go run -race 3_race.go
==================
WARNING: DATA RACE
Read at 0x00c0000c4000 by main goroutine:
  main.main()
      /Users/jackson/go/src/theWayToGolang/15_goroutineAndChannel/2_Goroutine-emp/3_race/3_race.go:18 +0xfb

Previous write at 0x00c0000c4000 by goroutine 6:
  main.main.func1()
      /Users/jackson/go/src/theWayToGolang/15_goroutineAndChannel/2_Goroutine-emp/3_race/3_race.go:13 +0x64

Goroutine 6 (running) created at:
  main.main()
      /Users/jackson/go/src/theWayToGolang/15_goroutineAndChannel/2_Goroutine-emp/3_race/3_race.go:11 +0xc3
==================
[9448584 9258518 9561628 9219457 9155083 9542995 9603923 9391295 16362051 16777799]
Found 1 data race(s)
exit status 66

*/


