
As we’ve seen, in concurrent programs it’s often necessary to preempt operations
because of timeouts, cancellation, or failure of another portion of the system. We’ve
looked at the idiom of creating a done channel, which flows through your program
and cancels all blocking concurrent operations. This works well, but it’s also some‐
what limited.
It would be useful if we could communicate extra information alongside the simple
notification to cancel: why the cancellation was occuring, or whether or not our func‐
tion has a deadline by which it needs to complete.
It turns out that the need to wrap a done channel with this information is very com‐
mon in systems of any size, and so the Go authors decided to create a standard pat‐
tern for doing so. It started out as an experiment that lived outside the standard
library, but in Go 1.7, the context package was brought into the standard library,
making this a standard Go idiom to consider when working with concurrent code

## 背景
golang在1.6.2的时候还没有自己的context，在1.7的版本中就把golang.org/x/net/context包被加入到了官方的库中。golang 的 Context包，是专门用来简化对于处理单个请求的多个goroutine之间与请求域的数据、取消信号、截止时间等相关操作，这些操作可能涉及多个 API 调用。
比如有一个网络请求Request，每个Request都需要开启一个goroutine做一些事情，这些goroutine又可能会开启其他的goroutine。这样的话， 我们就可以通过Context，来跟踪这些goroutine，并且通过Context来控制他们的目的，这就是Go语言为我们提供的Context，中文可以称之为“上下文”。
另外一个实际例子是，在Go服务器程序中，每个请求都会有一个goroutine去处理。然而，处理程序往往还需要创建额外的goroutine去访问后端资源，比如数据库、RPC服务等。由于这些goroutine都是在处理同一个请求，所以它们往往需要访问一些共享的资源，比如用户身份信息、认证token、请求截止时间等。而且如果请求超时或者被取消后，所有的goroutine都应该马上退出并且释放相关的资源。这种情况也需要用Context来为我们取消掉所有goroutine
如果要使用可以通过 go get golang.org/x/net/context 命令获取这个包。

## Context 定义
ontext的主要数据结构是一种嵌套的结构或者说是单向的继承关系的结构，比如最初的context是一个小盒子，里面装了一些数据，之后从这个context继承下来的children就像在原本的context中又套上了一个盒子，然后里面装着一些自己的数据。或者说context是一种分层的结构，根据使用场景的不同，每一层context都具备有一些不同的特性，这种层级式的组织也使得context易于扩展，职责清晰。
context 包的核心是 struct Context，声明如下：

```go
type Context interface {

Deadline() (deadline time.Time, ok bool)

Done() <-chan struct{}

Err() error

Value(key interface{}) interface{}

}

```

可以看到Context是一个interface，在golang里面，interface是一个使用非常广泛的结构，它可以接纳任何类型。Context定义很简单，一共4个方法，我们需要能够很好的理解这几个方法


Deadline方法是获取设置的截止时间的意思，第一个返回式是截止时间，到了这个时间点，Context会自动发起取消请求；第二个返回值ok==false时表示没有设置截止时间，如果需要取消的话，需要调用取消函数进行取消。


Done方法返回一个只读的chan，类型为struct{}，我们在goroutine中，如果该方法返回的chan可以读取，则意味着parent context已经发起了取消请求，我们通过Done方法收到这个信号后，就应该做清理操作，然后退出goroutine，释放资源。之后，Err 方法会返回一个错误，告知为什么 Context 被取消。


Err方法返回取消的错误原因，因为什么Context被取消。


Value方法获取该Context上绑定的值，是一个键值对，所以要通过一个Key才可以获取对应的值，这个值一般是线程安全的。

## Context 的实现方法
Context 虽然是个接口，但是并不需要使用方实现，golang内置的context 包，已经帮我们实现了2个方法，一般在代码中，开始上下文的时候都是以这两个作为最顶层的parent context，然后再衍生出子context。这些 Context 对象形成一棵树：当一个 Context 对象被取消时，继承自它的所有 Context 都会被取消。两个实现如下：
```go
var (
	background = new(emptyCtx)

	todo = new(emptyCtx)
)

func Background() Context {

	return background

}

func TODO() Context {

	return todo
}

```
一个是Background，主要用于main函数、初始化以及测试代码中，作为Context这个树结构的最顶层的Context，也就是根Context，它不能被取消。
一个是TODO，如果我们不知道该使用什么Context的时候，可以使用这个，但是实际应用中，暂时还没有使用过这个TODO。
他们两个本质上都是emptyCtx结构体类型，是一个不可取消，没有设置截止时间，没有携带任何值的Context。

```go
type emptyCtx int

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {

	return
}

func (*emptyCtx) Done() <-chan struct{} {

	return nil
}

func (*emptyCtx) Err() error {

	return nil
}

func (*emptyCtx) Value(key interface{}) interface{} {

	return nil
}

```

## Context 的 继承
有了如上的根Context，那么是如何衍生更多的子Context的呢？这就要靠context包为我们提供的With系列的函数了。

```
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)

func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)

func WithValue(parent Context, key, val interface{}) Context


```


通过这些函数，就创建了一颗Context树，树的每个节点都可以有任意多个子节点，节点层级可以有任意多个。
WithCancel函数，传递一个父Context作为参数，返回子Context，以及一个取消函数用来取消Context。
WithDeadline函数，和WithCancel差不多，它会多传递一个截止时间参数，意味着到了这个时间点，会自动取消Context，当然我们也可以不等到这个时候，可以提前通过取消函数进行取消。
WithTimeout和WithDeadline基本上一样，这个表示是超时自动取消，是多少时间后自动取消Context的意思。
WithValue函数和取消Context无关，它是为了生成一个绑定了一个键值对数据的Context，这个绑定的数据可以通过Context.Value方法访问到，这是我们实际用经常要用到的技巧，一般我们想要通过上下文来传递数据时，可以通过这个方法，如我们需要tarce追踪系统调用栈的时候。
With 系列函数详解
WithCancel
context.WithCancel生成了一个withCancel的实例以及一个cancelFuc，这个函数就是用来关闭ctxWithCancel中的 Done channel 函数。
下面来分析下源码实现，首先看看初始化，如下：

```
func newCancelCtx(parent Context) cancelCtx {
	return cancelCtx{
		Context: parent,
		done:    make(chan struct{}),
	}
}

func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
	c := newCancelCtx(parent)
	propagateCancel(parent, &c)
	return &c, func() { c.cancel(true, Canceled) }
}

```
newCancelCtx返回一个初始化的cancelCtx，cancelCtx结构体继承了Context，实现了canceler方法：
```go
//*cancelCtx 和 *timerCtx 都实现了canceler接口，实现该接口的类型都可以被直接canceled
type canceler interface {
    cancel(removeFromParent bool, err error)
    Done() <-chan struct{}
}


type cancelCtx struct {
    Context
    done chan struct{} // closed by the first cancel call.
    mu       sync.Mutex
    children map[canceler]bool // set to nil by the first cancel call
    err      error             // 当其被cancel时将会把err设置为非nil
}

func (c *cancelCtx) Done() <-chan struct{} {
    return c.done
}

func (c *cancelCtx) Err() error {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.err
}

func (c *cancelCtx) String() string {
    return fmt.Sprintf("%v.WithCancel", c.Context)
}

//核心是关闭c.done
//同时会设置c.err = err, c.children = nil
//依次遍历c.children，每个child分别cancel
//如果设置了removeFromParent，则将c从其parent的children中删除
func (c *cancelCtx) cancel(removeFromParent bool, err error) {
    if err == nil {
        panic("context: internal error: missing cancel error")
    }
    c.mu.Lock()
    if c.err != nil {
        c.mu.Unlock()
        return // already canceled
    }
    c.err = err
    close(c.done)
    for child := range c.children {
        // NOTE: acquiring the child's lock while holding parent's lock.
        child.cancel(false, err)
    }
    c.children = nil
    c.mu.Unlock()

    if removeFromParent {
        removeChild(c.Context, c) // 从此处可以看到 cancelCtx的Context项是一个类似于parent的概念
    }
}

```
可以看到，所有的children都存在一个map中；Done方法会返回其中的done channel， 而另外的cancel方法会关闭Done channel并且逐层向下遍历，关闭children的channel，并且将当前canceler从parent中移除。
WithCancel初始化一个cancelCtx的同时，还执行了propagateCancel方法，最后返回一个cancel function。
propagateCancel 方法定义如下：

```go
// propagateCancel arranges for child to be canceled when parent is.
func propagateCancel(parent Context, child canceler) {
	if parent.Done() == nil {
		return // parent is never canceled
	}
	if p, ok := parentCancelCtx(parent); ok {
		p.mu.Lock()
		if p.err != nil {
			// parent has already been canceled
			child.cancel(false, p.err)
		} else {
			if p.children == nil {
				p.children = make(map[canceler]struct{})
			}
			p.children[child] = struct{}{}
		}
		p.mu.Unlock()
	} else {
		go func() {
			select {
			case <-parent.Done():
				child.cancel(false, parent.Err())
			case <-child.Done():
			}
		}()
	}
}

```

propagateCancel 的含义就是传递cancel，从当前传入的parent开始（包括该parent），向上查找最近的一个可以被cancel的parent， 如果找到的parent已经被cancel，则将方才传入的child树给cancel掉，否则，将child节点直接连接为找到的parent的children中（Context字段不变，即向上的父亲指针不变，但是向下的孩子指针变直接了）； 如果没有找到最近的可以被cancel的parent，即其上都不可被cancel，则启动一个goroutine等待传入的parent终止，则cancel传入的child树，或者等待传入的child终结。
WithDeadLine
在withCancel的基础上进行的扩展，如果时间到了之后就进行cancel的操作，具体的操作流程基本上与withCancel一致，只不过控制cancel函数调用的时机是有一个timeout的channel所控制的。

## Context 使用原则 和 技巧

不要把Context放在结构体中，要以参数的方式传递，parent Context一般为Background
应该要把Context作为第一个参数传递给入口请求和出口请求链路上的每一个函数，放在第一位，变量名建议都统一，如ctx。
给一个函数方法传递Context的时候，不要传递nil，否则在tarce追踪的时候，就会断了连接
Context的Value相关方法应该传递必须的数据，不要什么数据都使用这个传递
Context是线程安全的，可以放心的在多个goroutine中传递
可以把一个 Context 对象传递给任意个数的 gorotuine，对它执行 取消 操作时，所有 goroutine 都会接收到取消信号。

## Context的常用方法实例
调用Context Done方法取消
```go
func Stream(ctx context.Context, out chan<- Value) error {

	for {
		v, err := DoSomething(ctx)

		if err != nil {
			return err
		}
		select {
		case <-ctx.Done():

			return ctx.Err()
		case out <- v:
		}
	}
}


```

2、通过 context.WithValue 来传值

```go
func main() {
	ctx, cancel := context.WithCancel(context.Background())

	valueCtx := context.WithValue(ctx, key, "add value")

	go watch(valueCtx)
	time.Sleep(10 * time.Second)
	cancel()

	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			//get value
			fmt.Println(ctx.Value(key), "is cancel")

			return
		default:
			//get value
			fmt.Println(ctx.Value(key), "int goroutine")

			time.Sleep(2 * time.Second)
		}
	}
}
```

3、超时取消 context.WithTimeout

```go
package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/net/context"
)

var (
	wg sync.WaitGroup
)

func work(ctx context.Context) error {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("Doing some work ", i)

		// we received the signal of cancelation in this channel
		case <-ctx.Done():
			fmt.Println("Cancel the context ", i)
			return ctx.Err()
		}
	}
	return nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	fmt.Println("Hey, I'm going to do some work")

	wg.Add(1)
	go work(ctx)
	wg.Wait()

	fmt.Println("Finished. I'm going home")
}



```

4、截止时间 取消 context.WithDeadline
```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(1 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancelation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("oversleep")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
```

5、
```go
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx,"【监控1】")
	go watch(ctx,"【监控2】")
	go watch(ctx,"【监控3】")

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name,"监控退出，停止了...")
			return
		default:
			fmt.Println(name,"goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}

```