
/*
Sometimes you want to time your channels operations: keep trying to do something,
and if you can’t do it in time just drop the ball.

To do so you can either use context or time, both are fine.
Context might be more idiomatic, time is a little bit more efficient, but they are almost identical:
*/

//TYPE 为任意类型
```go
func ToChanTimedContext(ctx context.Context, d time.Duration, message Type, c chan<- Type) (written bool) {
	ctx, cancel := context.WithTimeout(ctx, d)
	defer cancel()
	select {
	case c <- message:
		return true
	case <-ctx.Done():
		return false
	}
}

func ToChanTimedTimer(d time.Duration, message Type, c chan<- Type) (written bool) {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case c <- message:
		return true
	case <-t.C:
		return false
	}
}
```
Since performance is not really relevant here (after all we are waiting) the only difference I found is that the solution using context performs more allocations(also because the one with the timer can be further optimized to recycle timers).

Beware that re-using timers is tricky, so keep in mind that it might not be worth the risk to just save 10 allocs/op.

## First come first served
Sometimes you want to write the same message to many channels, writing to whichever is available first, but never writing the same message twice on the same channel.

To do this there are two ways: you can mask the channels with local variables, and disable the select cases accordingly, or use goroutines and waits.
```go
func FirstComeFirstServedSelect(message Type, a, b chan<- Type) {
	for i := 0; i < 2; i++ {
		select {
		case a <- message:
			a = nil
		case b <- message:
			b = nil
		}
	}
}

func FirstComeFirstServedGoroutines(message Type, a, b chan<- Type) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() { a <- message; wg.Done() }()
	go func() { b <- message; wg.Done() }()
	wg.Wait()
}
```
Please note that in this case performance might matter, and at the time of writing the solution that spawns goroutines takes almost 4 times more than the one with select.

If the amount of channels is not known at compile time, the first solution becomes trickier, but it is still possible, while the second one stays basically unchanged.

NOTE: if your program has many moving parts of unknown size, it might be worth revising your design, as it is very likely possible to simplify it.

If your code survives you review and still has unbound moving parts, here are the two solutions to support that:
```go
func FirstComeFirstServedGoroutinesVariadic(message Type, chs ...chan<- Type) {
	var wg sync.WaitGroup
	wg.Add(len(chs))
	for _, c := range chs {
		c := c
		go func() { c <- message; wg.Done() }()
	}
	wg.Wait()
}

func FirstComeFirstServedSelectVariadic(message Type, chs ...chan<- Type) {
	cases := make([]reflect.SelectCase, len(chs))
	for i, ch := range chs {
		cases[i] = reflect.SelectCase{
			Dir:  reflect.SelectSend,
			Chan: reflect.ValueOf(ch),
			Send: reflect.ValueOf(message),
		}
	}
	for i := 0; i < len(chs); i++ {
		chosen, _, _ := reflect.Select(cases)
		cases[chosen].Chan = reflect.ValueOf(nil)
	}
}
```
Needless to say: the solution using reflection is almost two orders of magnitude slower than the one with goroutines and unreadable, so please don’t use it.

Put it together
In case you want to both try a several sends for a while and abort if it’s taking too long here are two solutions: one with time+select and one with context+go. The first one is better if the amount of channels is known at compile time, while the other one should be used when it is not.
```go
func ToChansTimedTimerSelect(d time.Duration, message Type, a, b chan Type) (written int) {
	t := time.NewTimer(d)
	for i := 0; i < 2; i++ {
		select {
		case a <- message:
			a = nil
		case b <- message:
			b = nil
		case <-t.C:
			return i
		}
	}
	t.Stop()
	return 2
}

func ToChansTimedContextGoroutines(ctx context.Context, d time.Duration, message Type, ch ...chan Type) (written int) {
	ctx, cancel := context.WithTimeout(ctx, d)
	defer cancel()
	var (
		wr int32
		wg sync.WaitGroup
	)
	wg.Add(len(ch))
	for _, c := range ch {
		c := c
		go func() {
			defer wg.Done()
			select {
			case c <- message:
				atomic.AddInt32(&wr, 1)
			case <-ctx.Done():
			}
		}()
	}
	wg.Wait()
	return int(wr)
}

```

## 参考
https://blogtitle.github.io/go-advanced-concurrency-patterns-part-2-timers/