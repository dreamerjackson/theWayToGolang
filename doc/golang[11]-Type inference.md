# golang快速入门[8.2]-自动类型推断的秘密

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

## 前言
* 在上文中我们学习了变量的各种概念和go语言中的类型系统
* 我们将在本文中学习到:
  + 什么是自动类型推断
  + 为什么需要自动类型推断
  + go语言中自动类型推断的特点与陷阱
  + go语言在编译时是如何进行自动类型推断的

## 类型推断(Type inference)
* 类型推断是编程语言在编译时自动解释表达式数据类型的能力，通常在函数式编程的语言（例如Haskell）中存在,类型推断的优势主要在于可以省略类型，这使编程任务更加容易。
* 明确的指出变量的类型在编程语言中很常见，编译器在多大程度上可以做到这一点，因语言而异。例如，某些编译器可以通过这种方式推断出值：变量，函数参数和返回值。
* go语言作为静态类型语言表示在编译时就需要知道变量的类型

## 类型推断的优势
* 使编译器支持诸如类型推断之类的东西的实际意义可以分为两个主要部分。首先，如果使用得当，它可以使代码更易读，例如，可以将如下C ++代码：
```
vector<int> v;
vector<int>::iterator itr = v.iterator();
```
变为：
```
vector<int> v;
auto itr = v.iterator();
```
* 尽管在这里获得的收益似乎微不足道，但是如果类型更加复杂，则类型推断的价值变得显而易见。在许多情况下，这将使我们减少代码中的冗余信息。
* 类型推断还用于其他功能,`Haskell`语言可以编写为：
```
succ x = x + 1
```
* 上面的函数中，不管变量X是什么类型,加1并返回结果。
* 尽管如此,显式的指出类型仍然有效，因为编译器可以更轻松地了解代码实际应执行的操作，不太可能犯任何错误。

## go语言中的类型推断
如上所述，类型推断的能力每个语言是不相同的，在go语言中根据开发人员的说法，他们的目标是减少在静态类型语言中发现的混乱情况。他们认为许多像Java或C++这样的语言中的类型系统过于繁琐，并且他们更喜欢动态类型语言中的方法。
* 因此，在设计Go时，他们从这些语言中借鉴了一些想法。 这些想法之一是对变量使用简单的类型推断，给人以编写动态类型代码的感觉，同时仍然使用静态类型的好处
* 如前所述，类型推断可以涵盖参数和返回值之类的内容，但是Go中没有
* 在实践中，可以通过在声明新变量或常量时简单地忽略类型信息，或使用`:=`表示法（表示没有新的缩写）来触发Go中的类型推断
* 在Go中，以下三个语句是等效的：
```
var a int = 10
var a = 10
a := 10
```

* Go的类型推断在处理包含标识符的推断方面是半完成的。 本质上，编译器将不允许对从`标识符`引用的值进行类型转换，举几个例子：
* 下面这段代码正常运行，并且a的类型为float64
```
a := 1 + 1.1
```
* 下面的代码仍然正确,a会被推断为浮点数,`1`会变为浮点数与a的值相加
```
a := 1.1
b := 1 + a
```
* 但是，下面代码将会错误，即a的值已被推断为整数，而1.1为浮点数,但是不能将a强制转换为浮点数,相加失败。编译器报错:constant 1.1 truncated to integer
```
a := 1
b := a + 1.1
```

* 下面的类型会犯相同的错误，编译器提示:,invalid operation: a + b (mismatched types int and float64)
```
a := 1
b := 1.1
c := a + b
```

## 详细的实现说明
* 在之前的这篇文章中[go语言如何编译为机器码](https://zhuanlan.zhihu.com/p/107665043)，我们介绍了编译器执行的过程：词法分析 => 语法分析 => 类型检查 => 中间代码 => 代码优化 => 生成机器码
* 编译阶段的代码位于`go/src/cmd/compile`文件中
#### 词法分析阶段
* 具体来说,在词法分析阶段，会将赋值右边的常量解析为一个未定义的类型,类型有如下几种：其中ImagLit代表复数，IntLit代表整数...
```
//go/src/cmd/compile/internal/syntax
const (
 IntLit LitKind = iota
 FloatLit
 ImagLit
 RuneLit
 StringLit
)
```




* go语言源代码采用UTF-8的编码方式,在进行词法分析时当遇到需要赋值的常量操作时，会逐个的读取后面常量的UTF-8字符。字符串的首字符为`"`,数字的首字母为'0'-'9'。实现函数位于：
```
// go/src/cmd/compile/internal/syntax

func (s *scanner) next() {
...
switch c {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		s.number(c)
	case '"':
		s.stdString()
	case '`':
		s.rawString()
    ...
```

* 因此对于整数、小数等常量的识别就显得非常简单。具体来说,一个整数就是全是"0"-"9"的数字。一个浮点数就是字符中有"."号的数字，字符串就是首字符为`"`
* 下面列出的函数为小数和整数语法分析的具体实现:
```
// go/src/cmd/compile/internal/syntax
func (s *scanner) number(c rune) {
	s.startLit()

	base := 10        // number base
	prefix := rune(0) // one of 0 (decimal), '0' (0-octal), 'x', 'o', or 'b'
	digsep := 0       // bit 0: digit present, bit 1: '_' present
	invalid := -1     // index of invalid digit in literal, or < 0

	// integer part
	var ds int
	if c != '.' {
		s.kind = IntLit
		if c == '0' {
			c = s.getr()
			switch lower(c) {
			case 'x':
				c = s.getr()
				base, prefix = 16, 'x'
			case 'o':
				c = s.getr()
				base, prefix = 8, 'o'
			case 'b':
				c = s.getr()
				base, prefix = 2, 'b'
			default:
				base, prefix = 8, '0'
				digsep = 1 // leading 0
			}
		}
		c, ds = s.digits(c, base, &invalid)
		digsep |= ds
	}

	// fractional part
	if c == '.' {
		s.kind = FloatLit
		if prefix == 'o' || prefix == 'b' {
			s.error("invalid radix point in " + litname(prefix))
		}
		c, ds = s.digits(s.getr(), base, &invalid)
		digsep |= ds
	}
...
```

* 我们以赋值操作`a := 333`为例, 当完成词法分析时, 此赋值语句用`AssignStmt`表示。
```
	AssignStmt struct {
		Op       Operator // 0 means no operation
		Lhs, Rhs Expr     // Rhs == ImplicitOne means Lhs++ (Op == Add) or Lhs-- (Op == Sub)
		simpleStmt
	}
```
* 其中`Op`代表操作符,在这里是赋值操作,Lhs与Rhs分别代表左右两个表达式,左边代表了`变量a`,右边代表了整数`2`，此时右边整数的类型为`intLit`
#### 抽象语法树阶段
* 接着生成在抽象语法树AST时, 会将赋值`AssignStmt`变为一个`Node`,`node`结构体是对于抽象语法树中节点的抽象。
```
type Node struct {
	// Tree structure.
	// Generic recursive walks should follow these fields.
	Left  *Node
	Right *Node
	Ninit Nodes
	Nbody Nodes
	List  Nodes
	Rlist Nodes
    E   interface{} // Opt or Val, see methods below
    ...
```
* 仍然是左节点代表了左边的`变量a`,右边代表了整数`2`。
* 此时在E接口中，如果为整数会存储mpint类型,mpint存储整数常量
* 具体的代码如下,如果为IntLit类型，转换为Mpint类型,其他类型类似。
* 但是注意，此时左边的节点还是没有任何类型的。
```
// go/src/cmd/compile/internal/gc
func (p *noder) basicLit(lit *syntax.BasicLit) Val {
	// TODO: Don't try to convert if we had syntax errors (conversions may fail).
	//       Use dummy values so we can continue to compile. Eventually, use a
	//       form of "unknown" literals that are ignored during type-checking so
	//       we can continue type-checking w/o spurious follow-up errors.
	switch s := lit.Value; lit.Kind {
	case syntax.IntLit:
		checkLangCompat(lit)
		x := new(Mpint)
		x.SetString(s)
		return Val{U: x}

	case syntax.FloatLit:
		checkLangCompat(lit)
		x := newMpflt()
		x.SetString(s)
		return Val{U: x}
```
我们可以看到AST阶段整数存储通过math/big.int进行高精度存储。
```
// Mpint represents an integer constant.
type Mpint struct {
	Val  big.Int
	Ovf  bool // set if Val overflowed compiler limit (sticky)
	Rune bool // set if syntax indicates default type rune
}
```

* 最后在抽象语法树进行类型检查的阶段，会完成最终的赋值操作。将右边常量的类型赋值给左边变量的类型。
* 具体的函数位于`typecheckas`，将右边的类型赋值给左边
```
func typecheckas(n *Node) {
...
if n.Left.Name != nil && n.Left.Name.Defn == n && n.Left.Name.Param.Ntype == nil {
		n.Right = defaultlit(n.Right, nil)
		n.Left.Type = n.Right.Type
	}
}
...
```

* mpint对应的为`CTINT`标识。最终左边的变量存储的类型为`types.Types[TINT]`
```
func (v Val) Ctype() Ctype {
 switch x := v.U.(type) {
 default:
  Fatalf("unexpected Ctype for %T", v.U)
  panic("unreachable")
 case nil:
  return 0
 case *NilVal:
  return CTNIL
 case bool:
  return CTBOOL
 case *Mpint:
  if x.Rune {
   return CTRUNE
  }
  return CTINT
 case *Mpflt:
  return CTFLT
 case *Mpcplx:
  return CTCPLX
 case string:
  return CTSTR
 }
}
```
types.Types是一个数组，存储了不同标识对应的go语言中的实际类型。
```
var Types [NTYPE]*Type
```

`Type`是go语言中类型的存储结构,`types.Types[TINT]`最终代表的类型为`int`类型
```
// A Type represents a Go type.
type Type struct {
	Extra interface{}

	// Width is the width of this Type in bytes.
	Width int64 // valid if Align > 0

	methods    Fields
	allMethods Fields

	Nod  *Node // canonical OTYPE node
	Orig *Type // original type (type literal or predefined type)

	// Cache of composite types, with this type being the element type.
	Cache struct {
		ptr   *Type // *T, or nil
		slice *Type // []T, or nil
	}

	Sym    *Sym  // symbol containing name, for named types
	Vargen int32 // unique name for OTYPE/ONAME

	Etype EType // kind of type
	Align uint8 // the required alignment of this type, in bytes (0 means Width and Align have not yet been computed)

	flags bitset8
}

```



* 最后,我们可以用下面的代码来验证类型，输出结果为:int
```
a :=  2222
fmt.Printf("%T",a)
```

## 总结
* 在本文中,我们介绍了自动类型推断的内涵以及其意义。同时,我们用例子指出了go语言中自动类型推断的特点。
* 最后,我们用`a:=2`为例,介绍了go语言在编译时是如何进行自动类型推断的。
* 具体来说，go语言在编译时对于数字的处理首先采用了math包中进行了高精度的处理,接着会转换为go语言中的标准类型,int或float64.在本文中没有对字符串等做详细介绍，留给以后的文章。
* see you~




## 参考资料
* [项目链接](https://github.com/dreamerjackson/theWayToGolang)
* [作者知乎](https://www.zhihu.com/people/ke-ai-de-xiao-tu-ji-71)
* [blog](https://dreamerjonson.com/)
* [Type inference](https://en.wikipedia.org/wiki/Type_inference)
* [Rob Pike:Less is exponentially more](https://commandcenter.blogspot.com/2012/06/)
* [Type inference for go](http://fileadmin.cs.lth.se/cs/Education/EDAN70/CompilerProjects/2015/Reports/GigovicMalmros.pdf)
