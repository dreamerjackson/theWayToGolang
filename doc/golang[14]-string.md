# golang快速入门[9.1]-深入字符串的存储、编译与运行

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
* [golang快速入门[8.3]-深入理解浮点数](https://zhuanlan.zhihu.com/p/115888814)
* [golang快速入门[8.4]-常量与隐式类型转换](https://zhuanlan.zhihu.com/p/118316486)

## 前言
* 在常量和自动类型推断的文章中,我们介绍过整数、浮点数在词法解析阶段的过程。简单的说,整数是全为数字的常量,浮点数是带了`小数点`的常量。字符串也一样,字符串常量声明有两种方式:
```
var a string = `hello world`
var b string = "hello world"
```

* 词法解析阶段,挨个的读取Uft-8字符, 当发现了`单撇号`或者是`双引号`时,说明其是一个字符串。解析函数如下
```
func (s *scanner) next() {
    ...
	c := s.getr()
	for c == ' ' || c == '\t' || c == '\n' && !nlsemi || c == '\r' {
		c = s.getr()
	}
	// token start
	s.line, s.col = s.source.line0, s.source.col0
	if isLetter(c) || c >= utf8.RuneSelf && s.isIdentRune(c, true) {
		s.ident()
		return
	}

	switch c {
	case '"':
		s.stdString()

	case '`':
		s.rawString()

	...
	}
```
* 解析时`单撇号` 会调用rawString,`双引号` 会调用stdString,两者略微有所不同
* `单撇号`比较简单，始终要寻找下一个配对的`单撇号`

```
func (s *scanner) rawString() {
	s.startLit()

	for {
		r := s.getr()
		if r == '`' {
			break
		}
		if r < 0 {
			s.errh(s.line, s.col, "string not terminated")
			break
		}
	}
	// We leave CRs in the string since they are part of the
	// literal (even though they are not part of the literal
	// value).

	s.nlsemi = true
	s.lit = string(s.stopLit())
	s.kind = StringLit
	s.tok = _Literal
}
```
* `双引号`有所不同,其调用stdString函数。
```
func (s *scanner) stdString() {
	s.startLit()

	for {
		r := s.getr()
		if r == '"' {
			break
		}
		if r == '\\' {
			s.escape('"')
			continue
		}
		if r == '\n' {
			s.ungetr() // assume newline is not part of literal
			s.error("newline in string")
			break
		}
		if r < 0 {
			s.errh(s.line, s.col, "string not terminated")
			break
		}
	}

	s.nlsemi = true
	s.lit = string(s.stopLit())
	s.kind = StringLit
	s.tok = _Literal
}
```

* 当出现另一个`双引号`则直接退出,当出现了字符`\`,代表会对后面的字符进行转义。
* `双引号`不能出现如下的换行符,会报错。
```
str := " 微信:
1131052403 "
```

* 无论是标准字符串还是原始字符串最终都会被标记成 StringLit 类型的 Token 并传递到编译的下一个阶段
* s.lit = string(s.stopLit()) 将解析到的字节转换为字符串,例如"hello" 最后会被解析为""hello""
```
// go/src/cmd/compile/internal/gc
func (p *noder) basicLit(lit *syntax.BasicLit) Val {
	case syntax.StringLit:
		if len(s) > 0 && s[0] == '`' {
			// strip carriage returns from raw string
			s = strings.Replace(s, "\r", "", -1)
		}
		// Ignore errors because package syntax already reported them.
		u, _ := strconv.Unquote(s)
		return Val{U: u}

	default:
		panic("unhandled BasicLit kind")
	}
}
```
* 无论是 import 语句中包的路径、结构体中的字段标签还是表达式中的字符串都会使用` strings.Replace `方法将原生字符串中最后的换行符删除并对字符串 Token 进行 Unquote（`strconv.Unquote(s)`），也就是去掉字符串两边的引号等无关干扰，还原其本来的面目。例如将""hello"" 转换为 "hello"

## 字符串拼接
* op操作为:OADDSTR
* 常量中的字符串函数会在语法分析阶段调用sum函数进行拼接。例如对于`"hello"+"world"`,会在noder.sum函数中完成拼接。
```
/usr/local/go/src/cmd/compile/internal/gc/noder.go
func (p *noder) sum(x syntax.Expr) *Node {
	for i := len(adds) - 1; i >= 0; i-- {
		add := adds[i]

		r := p.expr(add.Y)
		if Isconst(r, CTSTR) && r.Sym == nil {
			if nstr != nil {
				// Collapse r into nstr instead of adding to n.
				chunks = append(chunks, r.Val().U.(string))
				continue
			}

			nstr = r
			chunks = append(chunks, nstr.Val().U.(string))
		} else {
			if len(chunks) > 1 {
				nstr.SetVal(Val{U: strings.Join(chunks, "")})
			}
			nstr = nil
			chunks = chunks[:0]
		}
		n = p.nod(add, OADD, n, r)
	}
	if len(chunks) > 1 {
		nstr.SetVal(Val{U: strings.Join(chunks, "")})
	}

	return n
}
```
* 但是如果是变量之间的拼接,例如对于如下代码,其拼接操作是在运行时完成的。
```
	var a = "hello"
	str :=  a + "xxs"
```
* 在语法分析阶段会做一些准备工作。例如在类型检查阶段`typecheck1`函数进行赋值和字符串拼接语义。
* 在walkexpr函数中,还会进行准备工作,决定使用运行时的哪一个拼接函数。
```
go/src/cmd/compile/internal/gc/walk.go
func walkexpr(n *Node, init *Nodes) *Node {
	case OADDSTR:
		n = addstr(n, init)
}
```

* walkexpr函数中调用函数`addstr(n, init)`
* 当拼接数量小于等于5个时，会调用运行时concatstring1~concatstring5之中的函数
* 当字符串的数量大于5个时,调用运行时concatstrings函数,并且字符串通过`切片`传入
```
func addstr(n *Node, init *Nodes) *Node {
	// build list of string arguments
	args := []*Node{buf}
	for _, n2 := range n.List.Slice() {
		args = append(args, conv(n2, types.Types[TSTRING]))
	}

	var fn string
	if c <= 5 {
		// small numbers of strings use direct runtime helpers.
		// note: orderexpr knows this cutoff too.
		fn = fmt.Sprintf("concatstring%d", c)
	} else {
		// large numbers of strings are passed to the runtime as a slice.
		fn = "concatstrings"

		t := types.NewSlice(types.Types[TSTRING])
		slice := nod(OCOMPLIT, nil, typenod(t))
		if prealloc[n] != nil {
			prealloc[slice] = prealloc[n]
		}
		slice.List.Set(args[1:]) // skip buf arg
		args = []*Node{buf, slice}
		slice.Esc = EscNone
	}

	cat := syslook(fn)
	r := nod(OCALL, cat, nil)
	r.List.Set(args)
	r = typecheck(r, ctxExpr)
	r = walkexpr(r, init)
	r.Type = n.Type

	return r
}
```
* 运行时字符串string的表示结构为
```
type StringHeader struct {
	Data uintptr
	Len  int
}
```
* 运行时具体的拼接代码如下，其实无论使用 concatstring{2,3,4,5} 中的哪一个，最终都会调用 runtime.concatstrings，该函数会先对传入的切片参数进行遍历，先过滤空字符串并计算拼接后字符串的长度。
```
/usr/local/go/src/runtime/string.go
func concatstrings(buf *tmpBuf, a []string) string {
	idx := 0
	l := 0
	count := 0
	for i, x := range a {
		n := len(x)
		if n == 0 {
			continue
		}
		if l+n < l {
			throw("string concatenation too long")
		}
		l += n
		count++
		idx = i
	}
	if count == 0 {
		return ""
	}

	// If there is just one string and either it is not on the stack
	// or our result does not escape the calling frame (buf != nil),
	// then we can return that string directly.
	if count == 1 && (buf != nil || !stringDataOnStack(a[idx])) {
		return a[idx]
	}
	s, b := rawstringtmp(buf, l)
	for _, x := range a {
		copy(b, x)
		b = b[len(x):]
	}
	return s
}

func concatstring2(buf *tmpBuf, a [2]string) string {
	return concatstrings(buf, a[:])
}

func concatstring3(buf *tmpBuf, a [3]string) string {
	return concatstrings(buf, a[:])
}

func concatstring4(buf *tmpBuf, a [4]string) string {
	return concatstrings(buf, a[:])
}

func concatstring5(buf *tmpBuf, a [5]string) string {
	return concatstrings(buf, a[:])
}
```

* 这里要注意，如果拼接后的字符串大小 小于32字节时,会有一个临时的缓存供其使用。如果拼接后的字符串大小 `大于` 32字节时,会请求分配内存。
* 拼接的过程就是开辟一个足够大的内存空间，并将多个字符串存入其中的过程。期间会涉及到内存的`Copy`拷贝
```
func rawstringtmp(buf *tmpBuf, l int) (s string, b []byte) {
	if buf != nil && l <= len(buf) {
		b = buf[:l]
		s = slicebytetostringtmp(b)
	} else {
		s, b = rawstring(l)
	}
	return
}
```

## 字符串与字节数组的转换
* 字节数组与字符串相互转换的形式如下:
```
	a := "微信:1131052403"
	b := []byte(a)
	c := string(b)
```

* 需要注意的是,字节数组与字符串的相互转换并不是无损的简单的一个指针的差别。而是涉及到了拷贝！因此相对而言,其仍然是消耗资源的。
* 如下为字节数组转换为字符串
```
func slicebytetostring(buf *tmpBuf, b []byte) (str string) {
	...
	var p unsafe.Pointer
	if buf != nil && len(b) <= len(buf) {
		p = unsafe.Pointer(buf)
	} else {
		p = mallocgc(uintptr(len(b)), nil, false)
	}
	stringStructOf(&str).str = p
	stringStructOf(&str).len = len(b)
	memmove(p, (*(*slice)(unsafe.Pointer(&b))).array, uintptr(len(b)))
	return
}

```

## 如下为字符串转换为字节数组
```
func stringtoslicebyte(buf *tmpBuf, s string) []byte {
	var b []byte
	if buf != nil && len(s) <= len(buf) {
		*buf = tmpBuf{}
		b = buf[:len(s)]
	} else {
		b = rawbyteslice(len(s))
	}
	copy(b, s)
	return b
}
```

## 总结
* 本节我们深入介绍了字符串,字符常量存储于静态存储区,其内容不可以被改变。声明时有`单撇号`或者是`双引号`两种方法
* 字符常量的拼接发生在编译时,变量字符串的拼接发生在运行时。如果拼接后的字符串大小 小于32字节时,会有一个临时的缓存供其使用。如果拼接后的字符串大小 `大于` 32字节时,会请求分配内存
* 需要注意的是,字节数组与字符串的相互转换并不是无损的简单的一个指针的差别。而是涉及到了拷贝！因此相对而言,其仍然是消耗资源的
* 本文还对编译时和运行时涉及到的函数进行了具体的说明
* see you~

## 参考资料
* [项目链接](https://github.com/dreamerjackson/theWayToGolang)
* [作者知乎](https://www.zhihu.com/people/ke-ai-de-xiao-tu-ji-71)
* [blog](https://dreamerjonson.com/)
* [draveness 字符串](https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-string/)