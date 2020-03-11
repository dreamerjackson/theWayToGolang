


## Go语言关键字


Go语言中类似if和switch的关键字有25个；关键字不能用于自定义名字，只能在特定语法结构中使用。
break      default       func     interface   select
case       defer         go       map         struct
chan       else          goto     package     switch
const      fallthrough   if       range       type
continue   for           import   return      var
此外，还有大约30多个预定义的名字，比如int和true等，主要对应内建的常量、类型和函数。
內建常量: true false iota nil

內建類型: int int8 int16 int32 int64
          uint uint8 uint16 uint32 uint64 uintptr
          float32 float64 complex128 complex64
          bool byte rune string error

內建函數: make len cap new append copy close delete
          complex real imag
          panic recover


## Go语言中的byte和rune
Go语言中byte和rune实质上就是uint8和int32类型。byte用来强调数据是raw data，而不是数字；而rune用来表示Unicode的code point。参考规范：

uint8       the set of all unsigned  8-bit integers (0 to 255)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)

byte        alias for uint8
rune        alias for int32

## Go语言的数值类型
* Go语言的数值类型包括几种不同大小的整形数、浮点数和复数。每种数值类型都决定了对应的大小范围和是否支持正负符号。让我们先从整型数类型开始介绍。
* Go语言同时提供了有符号和无符号类型的整数运算，有int8、int16、int32和int64四种截然不同大小的有符号整型数类型，分别对应8、16、32、64bit大小的有符号整型数，与此对应的是uint8、uint16、uint32和uint64四种无符号整型数类型。
* 还有两种一般对应特定CPU平台机器字大小的有符号和无符号整数int和uint。其中int是应用最广泛的数值类型。这两种类型都有同样的大小，32或64bit，但是我们不能对此做任何的假设；因为不同的编译器卽使在相同的硬件平台上可能产生不同的大小。
* Unicode字符rune类型是和int32等价的类型，通常用于表示一个Unicode码点。这两个名称可以互换使用。同样byte也是uint8类型的等价类型，byte类型一般用于强调数值是一个原始的数据而不是一个小的整数。
* 最后，还有一种无符号的整数类型uintptr，没有指定具体的bit大小但是足以容纳指针。uintptr类型只有在底层编程是才需要，特别是Go语言和C语言函数库或操作系统接口相交互的地方。在介绍指针时，会详细介绍它。