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

##

* 编译时新建一个切片
* 切片内元素的类型是在编译期间确定的
```
func NewSlice(elem *Type) *Type {
	if t := elem.Cache.slice; t != nil {
		if t.Elem() != elem {
			Fatalf("elem mismatch")
		}
		return t
	}

	t := New(TSLICE)
	t.Extra = Slice{Elem: elem}
	elem.Cache.slice = t
	return t
}
```
* 切片的类型
```
// Slice contains Type fields specific to slice types.
type Slice struct {
	Elem *Type // element type
}

```
* 运行时数据结构
```
type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}
```

## 初始化
```
slice := []int{1, 2, 3}
slice := make([]int, 10)
```

n.right 存储切片的长度,op为`OLITERAL`


##
当我们使用字面量 []int{1, 2, 3} 创建新的切片时，会创建一个array数组(`[3]int{1,2,3}`)存储于静态区中。同时会创建一个变量,
```
var vstat [3]int
vstat[0] = 1
vstat[1] = 2
vstat[2] = 3
var vauto *[3]int = new([3]int)
*vauto = vstat
slice := vauto[:]
```
```
	// recipe for var = []t{...}
	// 1. make a static array
	//	var vstat [...]t
	// 2. assign (data statements) the constant part
	//	vstat = constpart{}
	// 3. make an auto pointer to array and allocate heap to it
	//	var vauto *[...]t = new([...]t)
	// 4. copy the static array to the auto array
	//	*vauto = vstat
	// 5. for each dynamic part assign to the array
	//	vauto[i] = dynamic part
	// 6. assign slice of allocated heap to var
	//	var = vauto[:]
```
* 核心逻辑位于`slicelit`
```
func slicelit(ctxt initContext, n *Node, var_ *Node, init *Nodes)
```


## make 初始化
* 使用`make` 关键字,在typecheck1 检查阶段,op类型为`OMAKE`,

```
case OMAKE:
func typecheck1(n *Node, top int) (res *Node) {
switch t.Etype {
default:
    yyerror("cannot make type %v", t)
    n.Type = nil
    return n

case TSLICE:
    if i >= len(args) {
        yyerror("missing len argument to make(%v)", t)
        n.Type = nil
        return n
    }

    l = args[i]
    i++
    l = typecheck(l, ctxExpr)
    var r *Node
    if i < len(args) {
        r = args[i]
        i++
        r = typecheck(r, ctxExpr)
    }

    if l.Type == nil || (r != nil && r.Type == nil) {
        n.Type = nil
        return n
    }
    if !checkmake(t, "len", l) || r != nil && !checkmake(t, "cap", r) {
        n.Type = nil
        return n
    }
    if Isconst(l, CTINT) && r != nil && Isconst(r, CTINT) && l.Val().U.(*Mpint).Cmp(r.Val().U.(*Mpint)) > 0 {
        yyerror("len larger than cap in make(%v)", t)
        n.Type = nil
        return n
    }

    n.Left = l
    n.Right = r
    n.Op = OMAKESLICE
```

* op操作变为`OMAKESLICE`,并且左节点存储长度3, 右节点存储容量4
* 当数组长度比较小时,新建一个数组
类似于:
```
arr [4]int
r := arr[:3]
```
* 当数值长度比较大时,

```
case OMAKESLICE:
    l := n.Left
    r := n.Right
    if r == nil {
        r = safeexpr(l, init)
        l = r
    }
    t := n.Type
    if n.Esc == EscNone {
        if !isSmallMakeSlice(n) {
            Fatalf("non-small OMAKESLICE with EscNone: %v", n)
        }
        // var arr [r]T
        // n = arr[:l]
        i := indexconst(r)
        if i < 0 {
            Fatalf("walkexpr: invalid index %v", r)
        }
        t = types.NewArray(t.Elem(), i) // [r]T
        var_ := temp(t)
        a := nod(OAS, var_, nil) // zero temp
        a = typecheck(a, ctxStmt)
        init.Append(a)
        r := nod(OSLICE, var_, nil) // arr[:l]
        r.SetSliceBounds(nil, l, nil)
        r = conv(r, n.Type) // in case n.Type is named.
        r = typecheck(r, ctxExpr)
        r = walkexpr(r, init)
        n = r
    } else {
        // n escapes; set up a call to makeslice.
        // When len and cap can fit into int, use makeslice instead of
        // makeslice64, which is faster and shorter on 32 bit platforms.

        if t.Elem().NotInHeap() {
            yyerror("%v is go:notinheap; heap allocation disallowed", t.Elem())
        }

        len, cap := l, r

        fnname := "makeslice64"
        argtype := types.Types[TINT64]

        // Type checking guarantees that TIDEAL len/cap are positive and fit in an int.
        // The case of len or cap overflow when converting TUINT or TUINTPTR to TINT
        // will be handled by the negative range checks in makeslice during runtime.
        if (len.Type.IsKind(TIDEAL) || maxintval[len.Type.Etype].Cmp(maxintval[TUINT]) <= 0) &&
            (cap.Type.IsKind(TIDEAL) || maxintval[cap.Type.Etype].Cmp(maxintval[TUINT]) <= 0) {
            fnname = "makeslice"
            argtype = types.Types[TINT]
        }

        m := nod(OSLICEHEADER, nil, nil)
        m.Type = t

        fn := syslook(fnname)
        m.Left = mkcall1(fn, types.Types[TUNSAFEPTR], init, typename(t.Elem()), conv(len, argtype), conv(cap, argtype))
        m.Left.SetNonNil(true)
        m.List.Set2(conv(len, types.Types[TINT]), conv(cap, types.Types[TINT]))

        m = typecheck(m, ctxExpr)
        m = walkexpr(m, init)
        n = m
    }

```
* 调用运行时runtime.makeslice
## 参考资料
* [项目链接](https://github.com/dreamerjackson/theWayToGolang)
* [作者知乎](https://www.zhihu.com/people/ke-ai-de-xiao-tu-ji-71)
* [blog](https://dreamerjonson.com/)
* [draveness 字符串](https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-string/)