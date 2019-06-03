/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	doWork := func(
		done <-chan interface{},
		pulseInterval time.Duration,
	) (<-chan interface{}, <-chan time.Time) {
		heartbeat := make(chan interface{}) // <1>
		results := make(chan time.Time)
		go func() {
			defer close(heartbeat)
			defer close(results)

			pulse := time.Tick(pulseInterval)       // <2>
			workGen := time.Tick(2 * pulseInterval) // <3>

			sendPulse := func() {
				select {
				case heartbeat <- struct{}{}:
				default: // <4>
				}
			}
			sendResult := func(r time.Time) {
				for {
					select {
					case <-done:
						return
					case <-pulse: // <5>
						sendPulse()
					case results <- r:
						return
					}
				}
			}

			for {
				select {
				case <-done:
					return
				case <-pulse: // <5>
					sendPulse()
				case r := <-workGen:
					sendResult(r)
				}
			}
		}()
		return heartbeat, results
	}
	done := make(chan interface{})
	time.AfterFunc(10*time.Second, func() { close(done) }) // <1>

	const timeout = 2 * time.Second               // <2>
	heartbeat, results := doWork(done, timeout/2) // <3>
	for {
		select {
		case _, ok := <-heartbeat: // <4>
			if ok == false {
				return
			}
			fmt.Println("pulse")
		case r, ok := <-results: // <5>
			if ok == false {
				return
			}
			fmt.Printf("results %v\n", r.Second())
		case <-time.After(timeout): // <6>
			return
		}
	}
}

/*
<1>Here we set up a channel to send heartbeats on. We return this out of doWork.

<2>Here we set the heartbeat to pulse at the pulseInterval we were given. Every
pulseInterval there will be something to read on this channel.

<3>This is just another ticker used to simulate work coming in. We choose a dura‐
tion greater than the pulseInterval so that we can see some heartbeats coming
out of the goroutine.

<4>Note that we include a default clause. We must always guard against the fact
that no one may be listening to our heartbeat. The results emitted from the gor‐
outine are critical, but the pulses are not.

<5>Just like with done channels, anytime you perform a send or receive, you also
need to include a case for the heartbeat’s pulse.




<1>We set up the standard done channel and close it after 10 seconds. This gives our
goroutine time to do some work.


<2>Here we set our timeout period. We’ll use this to couple our heartbeat interval to
our timeout.
<3>We pass in timeout/2 here. This gives our heartbeat an extra tick to respond so
that our timeout isn’t too sensitive.
<4>Here we select on the heartbeat. When there are no results, we are at least guar‐
anteed a message from the heartbeat channel every timeout/2. If we don’t
receive it, we know there’s something wrong with the goroutine itself.
<5>Here we select from the results channel; nothing fancy going on here.
<6>Here we time out if we haven’t received either a heartbeat or a new result.
*/