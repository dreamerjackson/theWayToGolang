## interface接口
接口类型是对其它类型行为的抽象和概括；因为接口类型不会和特定的实现细节绑定在一起，通过这种抽象的方式我们可以让我们的函数更加灵活和更具有适应能力。
很多面向对象的语言都有相似的接口概念，但Go语言中接口类型的独特之处在于它是满足隐式实现的。
也就是说，我们没有必要对于给定的具体类型定义所有满足的接口类型；简单地拥有一些必需的方法就足够了。
这种设计可以让你创建一个新的接口类型满足已经存在的具体类型却不会去改变这些类型的定义；当我们使用的类型来自于不受我们控制的包时这种设计尤其有用。
在本章，我们会开始看到接口类型和值的一些基本技巧。顺着这种方式我们将学习几个来自标准库的重要接口。很多Go程序中都尽可能多的去使用标准库中的接口。

目前为止，我们看到的类型都是具体的类型。
一个具体的类型可以准确的描述它所代表的值并且展示出对类型本身的一些操作方式就像数字类型的算术操作，切片类型的索引、附加和取范围操作。
具体的类型还可以通过它的方法提供额外的行为操作。总的来说，当你拿到一个具体的类型时你就知道它的本身是什么和你可以用它来做什么。
在Go语言中还存在着另外一种类型：接口类型。接口类型是一种抽象的类型。
它不会暴露出它所代表的对象的内部值的结构和这个对象支持的基础操作的集合；它们只会展示出它们自己的方法。
也就是说当你有看到一个接口类型的值时，你不知道它是什么，唯一知道的就是可以通过它的方法来做什么。
在本书中，我们一直使用两个相似的函数来进行字符串的格式化：fmt.Printf它会把结果写到标准输出和fmt.Sprintf它会把结果以字符串的形式返回。
得益于使用接口，我们不必可悲的因为返回结果在使用方式上的一些浅显不同就必需把格式化这个最困难的过程复制一份。
实际上，这两个函数都使用了另一个函数fmt.Fprintf来进行封装。fmt.Fprintf这个函数对它的计算结果会被怎么使用是完全不知道的。
```go
package fmt
func Fprintf(w io.Writer, format string, args ...interface{}) (int, error)
func Printf(format string, args ...interface{}) (int, error) {
    return Fprintf(os.Stdout, format, args...)
}
func Sprintf(format string, args ...interface{}) string {
    var buf bytes.Buffer
    Fprintf(&buf, format, args...)
    return buf.String()
}
```
Fprintf的前缀F表示文件(File)也表明格式化输出结果应该被写入第一个参数提供的文件中。
在Printf函数中的第一个参数os.Stdout是*os.File类型；在Sprintf函数中的第一个参数&buf是一个指向可以写入字节的内存缓冲区，
然而它并不是一个文件类型尽管它在某种意义上和文件类型相似。
卽使Fprintf函数中的第一个参数也不是一个文件类型。它是io.Writer类型这是一个接口类型定义如下：
```go
package io
// Writer is the interface that wraps the basic Write method.
type Writer interface {
    // Write writes len(p) bytes from p to the underlying data stream.
    // It returns the number of bytes written from p (0 <= n <= len(p))
    // and any error encountered that caused the write to stop early.
    // Write must return a non-nil error if it returns n < len(p).
    // Write must not modify the slice data, even temporarily.
    //
    // Implementations must not retain p.
    Write(p []byte) (n int, err error)
}
```
io.Writer类型定义了函数Fprintf和这个函数调用者之间的约定。
一方面这个约定需要调用者提供具体类型的值就像*os.File和*bytes.Buffer，这些类型都有一个特定签名和行为的Write的函数。
另一方面这个约定保证了Fprintf接受任何满足io.Writer接口的值都可以工作。Fprintf函数可能没有假定写入的是一个文件或是一段内存，而是写入一个可以调用Write函数的值。
因为fmt.Fprintf函数没有对具体操作的值做任何假设而是仅仅通过io.Writer接口的约定来保证行为，所以第一个参数可以安全地传入一个任何具体类型的值只需要满足io.Writer接口。
一个类型可以自由的使用另一个满足相同接口的类型来进行替换被称作可替换性(LSP里氏替换)。这是一个面向对象的特征。
让我们通过一个新的类型来进行校验，下面*ByteCounter类型里的Write方法，仅仅在丢失写向它的字节前统计它们的长度。(在这个+=赋值语句中，让len(p)的类型和*c的类型匹配的转换是必须的。)
```go
type ByteCounter int
func (c *ByteCounter) Write(p []byte) (int, error) {
    *c += ByteCounter(len(p)) // convert int to ByteCounter
    return len(p), nil
}
因为*ByteCounter满足io.Writer的约定，我们可以把它传入Fprintf函数中；Fprintf函数执行字符串格式化的过程不会去关注ByteCounter正确的累加结果的长度。
var c ByteCounter
c.Write([]byte("hello"))
fmt.Println(c) // "5", = len("hello")
c = 0 // reset the counter
var name = "Dolly"
fmt.Fprintf(&c, "hello, %s", name)
fmt.Println(c) // "12", = len("hello, Dolly")
```
除了io.Writer这个接口类型，还有另一个对fmt包很重要的接口类型。Fprintf和Fprintln函数向类型提供了一种控制它们值输出的途径。

```go
package fmt

// The String method is used to print values passed
// as an operand to any format that accepts a string
// or to an unformatted printer such as Print.
type Stringer interface {
    String() string
}
```

## 参考资料：
How to avoid Go gotchas：https://divan.dev/posts/avoid_gotchas/?from=singlemessage&isappinstalled=0
Go Data Structures: Interfaces：https://research.swtch.com/interfaces