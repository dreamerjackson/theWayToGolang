/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"testing"
	"time"
)

func DoWork(
	done <-chan interface{},
	pulseInterval time.Duration,
	nums ...int,
) (<-chan interface{}, <-chan int) {
	heartbeat := make(chan interface{}, 1)
	intStream := make(chan int)
	go func() {
		defer close(heartbeat)
		defer close(intStream)

		time.Sleep(2 * time.Second)

		pulse := time.Tick(pulseInterval)
	numLoop: // <2>
		for _, n := range nums {
			for { // <1>
				select {
				case <-done:
					return
				case <-pulse:
					select {
					case heartbeat <- struct{}{}:
					default:
					}
				case intStream <- n:
					continue numLoop // <3>
				}
			}
		}
	}()

	return heartbeat, intStream
}

func TestDoWork_GeneratesAllNumbers(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{0, 1, 2, 3, 5}
	const timeout = 2 * time.Second
	heartbeat, results := DoWork(done, timeout/2, intSlice...)

	<-heartbeat // <4>

	i := 0
	for {
		select {
		case r, ok := <-results:
			if ok == false {
				return
			} else if expected := intSlice[i]; r != expected {
				t.Errorf("index %v: expected %v, but received %v,", i, expected, r)
			}
			i++
		case <-heartbeat: // <5>
		case <-time.After(timeout):
			t.Fatal("test timed out")
		}
	}
}
/*
Because of the heartbeat, we can safely write our test without timeouts. The only risk
we run is of one of our iterations taking an inordinate amount of time. If that’s
important to us, we can utilize the safer interval-based heartbeats and achieve perfect
safety

You’ve probably noticed that this version of the test is much less clear. The logic of
what we’re testing is a bit muddled. For this reason—if you’re reasonably sure the gor‐
outine’s loop won’t stop executing once it’s started—I recommend only blocking on
the first heartbeat and then falling into a simple range statement. You can write sepa‐
rate tests that specifically test for failing to close channels, loop iterations taking too
long, and any other timing-related issues.
Heartbeats aren’t strictly necessary when writing concurrent code, but this section
demonstrates their utility. For any long-running goroutines, or goroutines that need
to be tested, I highly recommend this pattern.
*/