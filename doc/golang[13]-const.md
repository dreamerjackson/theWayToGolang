# golang快速入门[8.4]-常量与隐式类型转换

## 前文
* [golang快速入门[1]-go语言导论](https://zhuanlan.zhihu.com/p/107658283)
* [golang快速入门[2.1]-go语言开发环境配置-windows](https://zhuanlan.zhihu.com/p/107659334)
* [golang快速入门[2.2]-go语言开发环境配置-macOS](https://zhuanlan.zhihu.com/p/107661202)
* [golang快速入门[2.3]-go语言开发环境配置-linux](https://zhuanlan.zhihu.com/p/107662649)
* [golang快速入门[3]-go语言helloworld](https://zhuanlan.zhihu.com/p/107664129)
* [golang快速入门[4]-go语言如何编译为机器码](https://zhuanlan.zhihu.com/p/107665043)
* [golang快速入门[5.1]-go语言是如何运行的-链接器](https://zhuanlan.zhihu.com/p/107665658)
* [golang快速入门[5.2]-go语言是如何运行的-内存概述](https://zhuanlan.zhihu.com/p/107807229)
* [golang快速入门[5.3]-go语言是如何运行的-内存分配](https://zhuanlan.zhihu.com/p/108598942)
* [golang快速入门[6.1]-集成开发环境-goland详解](https://zhuanlan.zhihu.com/p/109564120)
* [golang快速入门[6.2]-集成开发环境-emacs详解](https://zhuanlan.zhihu.com/p/110003756)
* [golang快速入门[7.1]-项目与依赖管理-gopath](https://zhuanlan.zhihu.com/p/110526009)
* [golang快速入门[7.2]-北冥神功—go module绝技](https://zhuanlan.zhihu.com/p/111409419)
* [golang快速入门[8.1]-变量类型、声明赋值、作用域声明周期与变量内存分配](https://zhuanlan.zhihu.com/p/112513336)
* [golang快速入门[8.2]-自动类型推断的秘密](https://zhuanlan.zhihu.com/p/115085755)
* [golang快速入门[8.3]-深入理解浮点数](https://zhuanlan.zhihu.com/p/115085755)

## 前言
* 在前文中我们学习了go语言中的自动类型推断以及浮点数的细节
* 我们将在本文中深入介绍go语言中另一个比较特殊的类型---const常量

## const常量
* Go语言最独特的功能之一就是该语言如何实现常量。Go语言规范中的常量规则是Go特有的。其在编译器级别提供了Go需求的灵活性，以使我们编写的代码更具可读性和直观性，同时仍保持类型安全的语言。

## 常量声明
* 可以在Go中使用类型或省略类型来声明常量，如下所示
```
const untypedInteger = 12345
const untypedFloatingPoint = 3.141592
const typedInteger int           = 12345
const typedFloatingPoint float64 = 3.141592
```
* 其中等式左边的常量可以叫做`命名常量`。等式右边的常量可以叫做`未命名常量`，拥有未定义的类型。
* 当有足够的内存来存储整数值时，可以始终精确地表示整数。由于规范要求整数常量的精度至少为256位，因此可以肯定地说整数常量在数学上是精确的。
* 为了获得数学上精确的浮点数，编译器可以采用不同的策略和选项。go语言规范未说明编译器必须如何执行此操作，它仅指定一组需要满足的强制性要求。
* 以下是当今不同的编译器用来实现数学上精确的浮点数的两种策略：
    + 一种策略是将所有浮点数表示为分数，并对这些分数使用有理算术。这些浮点数永远不会损失任何精度。
    + 另一种策略是使用高精度浮点数，其精度如此之高，以至于它们对于所有实际目的都是精确的。当我们使用具有数百位精度的浮点数时，精确值和近似值之间的差异实际上不存在。

## 常量的生存周期
* 常量只会在编译期间存在，因此其不会存储在内存之中，也就不可以被寻址.因此下面的代码是错误的。
```
    const k = 5
    address := &k
```

## 自动类型转换
* 如下所示,常量进行自动类型推断。我们在之前自动类型推断文章中其实就详细介绍了未定义的常量如何进行转换的。大致是在转换为具体的类型之前，会使用一种高精度的结构去存储
* 详细的介绍参见: [golang快速入门[8.2]-自动类型推断的秘密](https://zhuanlan.zhihu.com/p/115085755)
```
var myInt =123
```
## 隐式整数类型转换
* 在Go中，变量之间没有隐式类型转换。但是，编译器可以进行变量和常量之间的隐式类型转换。
* 例如：如下示例中，我们将常量123的整数类型隐式转换为int类型的值。由于常量的形式不使用小数点或指数，因此常量采用整数类型。只要不需要截断，就可以将整数类型的常量隐式转换为有符号和无符号整数变量。

```
var myInt int = 123
```
* 如果常量使用与整数类型兼容的形式，则也可以将浮点类型的常量隐式转换为整数变量：
```
var myInt int = 123.0
```
* 但是下面的转换却是不可行的
```
var myInt int = 123.1
```
## 隐式浮点类型转换
* 如下所示，编译器将在浮点类型的常量0.333到类型为float64的变量之间执行隐式转换。由于常量的形式使用小数点，因此该常量采用浮点类型。
```
var myFloat float64 = 0.333
```
* 编译器还可以在整数类型的常量到float64类型的变量之间执行隐式转换：
```
var myFloat float64 = 1
```

## 常量算术中的隐形转换
* 在其他常量和变量之间执行常量算术是我们在程序中经常要做的事情。它遵循规范中二进制运算符的规则。该规则规定，除非操作涉及`位运算`或`未类型化`的常量，否则操作数类型必须相同。
* 让我们看一个将两个常数相乘的例子：
```
var answer = 3 * 0.333
```
* 在go语言规范中，有一个关于常量表达式的规则专用于此操作：
* 除了移位操作，如果算数操作的操作数是不同类型的无类型常量，则结果类型的优先级为： 整数(int)<符文数(rune)<浮点数(float)<复数(Imag)
* 根据此规则，上面两个常数之间相乘的结果将是一个浮点数。因为浮点数要比整数优先级高
* 因此我们也能够理解:下面的例子结果为浮点数：
```
const third = 1 / 3.0
```
* 而下面的例子我们将在整数类型的两个常量之间进行除法。除法的结果将是一个整数类型的新常量。由于将3除以1的值表示小于1的数字，因此该除法的结果为整数常数0
```
const zero = 1 / 3
```

## 常量与具体类型的变量之间的算数转换规则
* 常量与具体类型的变量之间的算数，会使用已有的具体类型
* 例如下例中常量p结果为float64类型,常量2会转换为和b的类型相同
```
const b float64 = 1
const p = b * 2
```

* 下面的例子出错，因为2.3不能转化为b的类型int
```
const b int = 1
const p = b * 2.3
```

## 自定义类型的转换
* 下面再来看一个更复杂的例子,即用户自定义的类型
```
type Numbers int8
const One Numbers = 1
const Two         = 2 * One
```
* 在这里，我们声明了一个新类型，称为Numbers，其基本类型为int8。然后，我们以Numbers类型声明常量One，并分配整数类型的常量1。接下来，我们声明常量2，该常量通过将常量2和Numbers类型的常量One相乘而提升为Numbers类型。
* 让我们看一下标准库中的一个实际示例。time包声明了常量集:
```
type Duration int64

const (
    Nanosecond Duration = 1
    Microsecond         = 1000 * Nanosecond
    Millisecond         = 1000 * Microsecond
    Second              = 1000 * Millisecond
)
```

* 由于编译器将对常量执行隐式转换，因此我们可以在Go中像下面一样编写代码
```
package main

import (
    "fmt"
    "time"
)

const fiveSeconds = 5 * time.Second

func main() {
    now := time.Now()
    lessFiveNanoseconds := now.Add(-5)
    lessFiveSeconds := now.Add(-fiveSeconds)

    fmt.Printf("Now     : %v\n", now)
    fmt.Printf("Nano    : %v\n", lessFiveNanoseconds)
    fmt.Printf("Seconds : %v\n", lessFiveSeconds)
}

Output:
Now     : 2014-03-27 13:30:49.111038384 -0400 EDT
Nano    : 2014-03-27 13:30:49.111038379 -0400 EDT
Seconds : 2014-03-27 13:30:44.111038384 -0400 EDT
```
* 让我们看一下Time的Add方法的定义
```
func (t Time) Add(d Duration) Time
```

* Add方法接受一个类型为Duration的参数。让我们仔细看看程序中对Add的方法调用
```
var lessFiveNanoseconds = now.Add(-5)
var lessFiveMinutes = now.Add(-fiveSeconds)
```
* 编译器会将常量-5隐式转换为Duration类型的变量，以允许方法调用发生。同时,由于常量算术规则，常量FiveSeconds已经是Duration类型
* 但是如果指定了常量的类型，则调用不会成功,例如下面的例子
```
var difference int = -5
var lessFiveNano = now.Add(difference)

Compiler Error:
./const.go:16: cannot use difference (type int) as type time.Duration in function argument
```
* 如上所示,一旦我们使用具体类型整数值作为Add方法调用的参数，我们就会收到编译器错误。编译器将不允许在类型变量之间进行隐式类型转换。为了编译该代码，我们需要执行显式的类型转换
```
Add(time.Duration(difference))
```
* 常量是我们无需使用显式类型转换即可编写代码的唯一机制

## 编译时代码
* 对于涉及到常量的算术规则，统一在defaultlit2函数中进行了处理。首先判断操作符左边的有无类型,有的话就将操作符右边的类型转成左边的类型.
* 如果操作符左边为无类型,右边为有类型,则将左边的类型转换为右边的类型.
* 如果操作符左右都无具体类型,根据优先级决定类型的转换
```
// go/src/cmd/compile/internal/gc
func defaultlit2(l *Node, r *Node, force bool) (*Node, *Node) {
	if l.Type == nil || r.Type == nil {
		return l, r
	}
	if !l.Type.IsUntyped() {
		r = convlit(r, l.Type)
		return l, r
	}

	if !r.Type.IsUntyped() {
		l = convlit(l, r.Type)
		return l, r
	}

	if !force {
		return l, r
	}

	if l.Type.IsBoolean() {
		l = convlit(l, types.Types[TBOOL])
		r = convlit(r, types.Types[TBOOL])
	}

	lkind := idealkind(l)
	rkind := idealkind(r)
	if lkind == CTCPLX || rkind == CTCPLX {
		l = convlit(l, types.Types[TCOMPLEX128])
		r = convlit(r, types.Types[TCOMPLEX128])
		return l, r
	}

	if lkind == CTFLT || rkind == CTFLT {
		l = convlit(l, types.Types[TFLOAT64])
		r = convlit(r, types.Types[TFLOAT64])
		return l, r
	}

	if lkind == CTRUNE || rkind == CTRUNE {
		l = convlit(l, types.Runetype)
		r = convlit(r, types.Runetype)
		return l, r
	}

	l = convlit(l, types.Types[TINT])
	r = convlit(r, types.Types[TINT])

	return l, r
}

```

## 总结
* 在本文中介绍了常量的内涵、规则、常量的生存周期以及各种下情形下的隐式类型转换。
* 常量,Go语言最独特的功能之一，其是编译时期的独特功能。具体来说,隐式类型转换的规则为:有类型常量优先于无类型常量,当两个无类型常量算术运算时,结果类型的优先级为： 整数(int)<符文数(rune)<浮点数(float)<复数(Imag)
* see you~
## 参考资料
* [项目链接](https://github.com/dreamerjackson/theWayToGolang)
* [作者知乎](https://www.zhihu.com/people/ke-ai-de-xiao-tu-ji-71)
* [blog](https://dreamerjonson.com/)