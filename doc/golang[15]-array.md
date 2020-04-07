# golang快速入门[9.2]-深入数组用法、陷阱与编译时

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
* 在本节我们将介绍go语言中重要的数据类型——数组
* 数组是一个重要的数据类型,通常会与go语言另一个重要的结构：切片作对比。
* go语言中数组与其他语言有在显著的不同，包括其不能够进行添加,以及值拷贝的特性。在这一小节中,将会详细介绍。

## 数组的声明与定义
```
//声明三种方式
var arr [3]int
var arr2  = [4]int{1,2,3,4}
arr4 :=[...]int{2,3,4}
```

## 简单获取数组类型
```
fmt.Printf("类型arr3: %T,类型arr4: %T\n",arr3,arr4)
```

## 获取数组长度与通过下标获取
```
len(arr3)
arr3[2]
```
## 编译时
* 数组在编译时的数据类型为`TARRAY`,通过`NewArray`函数进行创建,AST节点的Op操作：`OARRAYLIT`
```
// NewArray returns a new fixed-length array Type.
func NewArray(elem *Type, bound int64) *Type {
	if bound < 0 {
		Fatalf("NewArray: invalid bound %v", bound)
	}
	t := New(TARRAY)
	t.Extra = &Array{Elem: elem, Bound: bound}
	t.SetNotInHeap(elem.NotInHeap())
	return t
}
```
* 内部的Array结构存储了数组中的类型以及数组的大小
```
// Array contains Type fields specific to array types.
type Array struct {
	Elem  *Type // element type
	Bound int64 // number of elements; <0 if unknown yet
}
```

* 数组的声明中,存在一个语法糖。`[...]int{2,3,4}`。 其实质与一般的数组声明类似的。
* 对于字面量的初始化方式,在编译时,通过`typecheckcomplit` 函数循环字面量分别进行赋值。
```
func typecheckcomplit(n *Node) (res *Node) {
	nl := n.List.Slice()
		for i2, l := range nl {
	        i++
			if i > length {
				length = i
				if checkBounds && length > t.NumElem() {
					setlineno(l)
					yyerror("array index %d out of bounds [0:%d]", length-1, t.NumElem())
					checkBounds = false
				}
			}
		}

		if t.IsDDDArray() {
			t.SetNumElem(length)
		}
	}
}
```
* 抽象的表达就是：
```
a:=[3]int{2,3,4}
变为
var arr [3]int
a[0] = 2
a[1] = 3
a[2] = 4
```
* 如果`t.IsDDDArray`判断到是语法糖的形式进行的数组初始化,那么会将其长度设置到数组中`t.SetNumElem(length)`.
* 在编译期的优化阶段,还会进行重要的优化。在函数`anylit`中,当数组的长度小于4时,在运行时会在栈中进行初始化`initKindDynamic`。当数组的长度大于4,会在静态区初始化数组`initKindStatic`.
```
func anylit(n *Node, var_ *Node, init *Nodes) {
	t := n.Type
	switch n.Op {
	case OSTRUCTLIT, OARRAYLIT:
		if !t.IsStruct() && !t.IsArray() {
			Fatalf("anylit: not struct/array")
		}

		if var_.isSimpleName() && n.List.Len() > 4 {
			...
			fixedlit(ctxt, initKindStatic, n, vstat, init)

			// copy static to var
			a := nod(OAS, var_, vstat)

			a = typecheck(a, ctxStmt)
			a = walkexpr(a, init)
			init.Append(a)

			// add expressions to automatic
			fixedlit(inInitFunction, initKindDynamic, n, var_, init)
			break
		}
}
```
* 他们都是通过`fixedlit`函数实现的。
```
func fixedlit(ctxt initContext, kind initKind, n *Node, var_ *Node, init *Nodes) {
	for _, r := range n.List.Slice() {
    // build list of assignments: var[index] = expr
    setlineno(a)
    a = nod(OAS, a, value)
    a = typecheck(a, ctxStmt)
	switch n.Op {
	    ...
		switch kind {
		case initKindStatic:
			genAsStatic(a)
		case initKindDynamic, initKindLocalCode:
			a = orderStmtInPlace(a, map[string][]*Node{})
			a = walkstmt(a)
			init.Append(a)
		default:
			Fatalf("fixedlit: bad kind %d", kind)
		}

	}
}


```

## 数组索引
```
 var a [3]int
 b := a[1]
```
* 数组访问越界是非常严重的错误，Go 语言中对越界的判断是可以在编译期间由静态类型检查完成的，`typecheck1` 函数会对访问数组的索引进行验证：
```
func typecheck1(n *Node, top int) (res *Node) {
	switch n.Op {
	case OINDEX:
		ok |= ctxExpr
		l := n.Left  // array
		r := n.Right // index
		switch n.Left.Type.Etype {
		case TSTRING, TARRAY, TSLICE:
			...
			if n.Right.Type != nil && !n.Right.Type.IsInteger() {
				yyerror("non-integer array index %v", n.Right)
				break
			}
			if !n.Bounded() && Isconst(n.Right, CTINT) {
				x := n.Right.Int64()
				if x < 0 {
					yyerror("invalid array index %v (index must be non-negative)", n.Right)
				} else if n.Left.Type.IsArray() && x >= n.Left.Type.NumElem() {
					yyerror("invalid array index %v (out of bounds for %d-element array)", n.Right, n.Left.Type.NumElem())
				}
			}
		}
	...
	}
}
```
* 访问数组的索引是非整数时会直接报错 —— non-integer array index %v；
* 访问数组的索引是负数时会直接报错 —— "invalid array index %v (index must be non-negative)"；
* 访问数组的索引越界时会直接报错 —— "invalid array index %v (out of bounds for %d-element array)"；
* 数组和字符串的一些简单越界错误都会在编译期间发现，比如我们直接使用整数或者常量访问数组，但是如果使用变量去访问数组或者字符串时,编译器就无法发现对应的错误了，这时就需要在运行时去判断错误。
```
i:= 3
m:= a[i]
```
* Go 语言运行时在发现数组、切片和字符串的越界操作会由运行时的 panicIndex 和 runtime.goPanicIndex 函数触发程序的运行时错误并导致崩溃退出：
```
TEXT runtime·panicIndex(SB),NOSPLIT,$0-8
	MOVL	AX, x+0(FP)
	MOVL	CX, y+4(FP)
	JMP	runtime·goPanicIndex(SB)

func goPanicIndex(x int, y int) {
	panicCheck1(getcallerpc(), "index out of range")
	panic(boundsError{x: int64(x), signed: true, y: y, code: boundsIndex})
}
```

* 最后要提到的是,即便数组的索引是变量。在某些时候任然能够在编译时通过优化检测出越界并在运行时报错。
* 例如对于一个简单的代码
```
a := [3]int{1,2,3}
b := 8
_ = a[b]
```
* 我们可以通过如下命令生成ssa.html。显示整个编译时的执行过程。
```
GOSSAFUNC=main GOOS=linux GOARCH=amd64 go tool compile close.go
```
* start阶段为最初生成ssa的阶段,
```
start
b1:-
v1 (?) = InitMem <mem>
v2 (?) = SP <uintptr>
v3 (?) = SB <uintptr>
v4 (15) = VarDef <mem> {arr} v1
v5 (15) = LocalAddr <*[3]int> {arr} v2 v4
v6 (15) = Zero <mem> {[3]int} [24] v5 v4
v7 (?) = Const64 <int> [1]
v8 (15) = LocalAddr <*[3]int> {arr} v2 v6
v9 (?) = Const64 <int> [0]
v10 (?) = Const64 <int> [3]
v11 (15) = PtrIndex <*int> v8 v9
v12 (15) = Store <mem> {int} v11 v7 v6
v13 (?) = Const64 <int> [2]
v14 (15) = LocalAddr <*[3]int> {arr} v2 v12
v15 (15) = PtrIndex <*int> v14 v7
v16 (15) = Store <mem> {int} v15 v13 v12
v17 (15) = LocalAddr <*[3]int> {arr} v2 v16
v18 (15) = PtrIndex <*int> v17 v13
v19 (15) = Store <mem> {int} v18 v10 v16
v20 (?) = Const64 <int> [4] (i[int])
v21 (17) = LocalAddr <*[3]int> {arr} v2 v19
v22 (17) = IsInBounds <bool> v20 v10
If v22 → b2 b3 (likely) (17)
b2: ← b1-
v25 (17) = PtrIndex <*int> v21 v20
v26 (17) = Copy <mem> v19
v27 (17) = Load <int> v25 v26 (elem[int])
Ret v26 (19)
b3: ← b1-
v23 (17) = Copy <mem> v19
v24 (17) = PanicBounds <mem> [0] v20 v10 v23
Exit v24 (17)
```
* 通过函数IsInBounds判断数组长度与索引大小进行对比。`v22 (17) = IsInBounds <bool> v20 v10`,如果失败即执行`v24 (17) = PanicBounds <mem> [0] v20 v10 v23`
* 在`genssa`生成汇编代码的阶段,我们能够看到直接被优化为了`00008 (17) CALL runtime.panicIndex(SB)` 即在运行时直接会触发Panic
```
genssa
# main.go
00000 (14) TEXT "".main(SB), ABIInternal
00001 (14) FUNCDATA $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
00002 (14) FUNCDATA $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
00003 (14) FUNCDATA $2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
v3
00004 (+17) PCDATA $0, $0
v3
00005 (+17) PCDATA $1, $0
v3
00006 (+17) MOVL $4, AX
v19
00007 (17) MOVL $3, CX
v24
00008 (17) CALL runtime.panicIndex(SB)
00009 (17) XCHGL AX, AX
00010 (?) END
```
## 数组的值拷贝问题
* 无论是赋值的`b`还是函数调用中的形参`c`,都是值拷贝的
```
a:= [3]int{1,2,3}
b = a

func Change(c [3]int){
    ...
}
```
我们可以通过简单的打印地址来验证:
```
package main

import "fmt"

func main() {
	a := [5]int{1,2,3,4,5}
	fmt.Printf("a:%p\n",&a)
	b:=a
	CopyArray(a)
	fmt.Printf("b:%p\n",&b)
}
//
func CopyArray( c [5]int){
	fmt.Printf("c:%p\n",&c)
}
```
输出为：
```
a:0xc00001a150
c:0xc00001a1b0
b:0xc00001a180
```
* 说明每一个数组在内存的位置都是不相同的,验证其是值拷贝

## 总结
* 数组是go语言中的特殊类型,其与其他语言不太一样。他不可以添加,但是可以获取值,获取长度。
* 同时,数组的拷贝都是值拷贝,因此不要尽量不要进行大数组的拷贝。
* 常量的下标以及某一些变量的下标的访问越界问题可以在编译时检测到,但是变量的下标的数组越界问题只会在运行时报错。
* 数组的声明中,存在一个语法糖。`[...]int{2,3,4}`,但是本质本没有什么差别
* 在编译期的优化阶段,还会进行重要的优化。当数组的长度小于4时,在运行时会在栈中进行初始化。当数组的长度大于4,会在静态区初始化数组
* 其实我们在go语言中对于数组用得较少，而是更多的使用切片。这是下一节的内容。see you~

## 参考资料
* [项目链接](https://github.com/dreamerjackson/theWayToGolang)
* [作者知乎](https://www.zhihu.com/people/ke-ai-de-xiao-tu-ji-71)
* [blog](https://dreamerjonson.com/)
* [draveness slice](https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-array/)
