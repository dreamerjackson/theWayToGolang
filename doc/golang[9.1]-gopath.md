# golang快速入门[7.1]-项目与依赖管理-gopath

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

## 前言
* 在之前文章中，我们介绍了go语言开发环境的配置
* 在本章中,我们将介绍go语言的项目结构、项目管理以及依赖管理。在本文中，我们关注`gopath`
## gopath是什么
* 在go语言开发环境配置文章中，我们介绍了配置`gopath`与`goroot`环境变量的步骤，但是并没有对其进行深入解释。可以在终端输入 `go env` 或者`go env gopath`查看具体的配置
```
C:\Windows\system32> go env
set GO111MODULE=
set GOARCH=amd64
set GOBIN=
set GOCACHE=C:\Users\jackson\AppData\Local\go-build
set GOENV=C:\Users\jackson\AppData\Roaming\go\env
set GOEXE=.exe
set GOFLAGS=
set GOHOSTARCH=amd64
set GOHOSTOS=windows
set GONOPROXY=
set GONOSUMDB=
set GOOS=windows
set GOPATH=C:\Users\jackson\go
set GOPRIVATE=
set GOPROXY=https://proxy.golang.org,direct
set GOROOT=c:\go
...
```
* 在go1.8之后,如果不指定gopath，则gopath是默认的。
    + 在mac,linux下为`$HOME/go`
    + windows 下为`%USERPROFILE%\go`
* `gopath` 可以理解为go语言的工作空间,内部存储了`src`,`bin`,`pkg` 三个文件夹
```
go/
├── bin
├── pkg
└── src
```

* `$GOPATH/bin`目录存储了通过`go install` 安装的二进制文件。操作系统使用$PATH环境变量来查找无完整路径即可执行的二进制应用程序。建议将此目录添加到全局$PATH变量中
* `$GOPATH/pkg`目录中,会存储预编译的obj文件(文件名根据操作系统的不同而不同,例如mac下为`darwin_amd64`)，以加快程序的后续编译。大多数开发人员不需要访问此目录。后面还会介绍，pkg下的mod文件还会存储`go module`的依赖。
* `$GOPATH/src`目录 存储我们项目的go代码。通常包含许多版本控制存储库（例如，由Git管理），每个存储库都包含了一个或多个package，每个package都在一个目录中包含一个或多个Go源文件。

* 因此,整个路径看起来像是：
```
go/
├── bin
     └── main.exe
├── pkg
     ├── darwin_amd64
     └── mod
└── src
    ├── github.com
    │   ├── tylfin
    │   │   ├── dynatomic
    │   │   └── geospy
    │   └── uudashr
    │       └── gopkgs
    └── golang.org
        └── x
            └── tools
```

* gopath具有多个作用，当我们想从github或其他地方获取go项目代码时,我们可以使用`go get`指令。 此时程序会默认的将代码存储到`$GOPATH/src`目录中。例如拉取`go get github.com/dreamerjackson/theWayToGolang`时,目录结构如下：
```
go/
├── bin
├── pkg
└── src
    └── github.com
           └── dreamerjackson
                     └── theWayToGolang
```
* 当我们使用`go get`的`-u`参数时，会将该项目以及项目所依赖的所有其他项目一并下载到`$GOPATH/src`目录中
* gopath的另一个功能是明确package的导入位置。前文我们介绍过，go代码通过package进行组织,在helloworl程序中，我们导入了go语言内置的`fmt` package.当我们要导入第三方时应该怎么做呢？其实如果我们在项目中导入了一个第三方包,例如
```
import "blue/red"
```
实际引用的是`$GOPATH/src/blue/red` 文件中的代码。

同理,如果导入为
```
import "github.com/gobuffalo/buffalo"
```
实际引用的是`$GOPATH/src/github.com/gobuffalo/buffalo` 文件中的代码。

## 下面我们用一个例子来说明导入第三方包
* 首先在`$GOPATH/src`中新建一个文件夹`mymath`，在文件夹中新建一个文件`add.go`
```
» mkdir mymath
» cd mymath
» touch add.go
```
add.go的内容如下，要注意导出的函数必须是首字母大写的，这是go语言的规则。
```
package mymath

func Add(a int, b int) int {
	return a + b
}

```
接着在`$GOPATH/src`中创建一个main.go文件，即可以完整的调用`mymath`包中的`add`函数。
```
package main

import (
	"fmt"
	"mymath"
)

func main() {
	result := mymath.Add(1, 2)
	fmt.Println("result:", result)
}
```
* 通过`go run main.go` 即可运行输出`1+2`的结果`3`
## gopath的优劣
* 相比于其他语言繁琐的配置,go语言中的工作空间`gopath`配置相对简单，容易理解
* gopath使得在文件系统组织整个代码更加简洁、结构化，但是限制在单一的工作空间中。
* gopath并没有解决版本依赖的问题,而将其留给了其他工具去实现。正因为如此,gopath中的代码就是一个唯一的master分支,并且强制使用各个模块最新的代码。
## 总结
* 本文介绍了gopath的含义、功能、优劣、以及如何通过GOPATH来组织项目,导入第三方库。
* 在go1.13之后，go官方已经开始全面拥抱`go module`.我们在下文中，将介绍`go module`的原理和用法，以及如何通过`go module`进行go语言的依赖管理与项目组织。

## 参考资料
* [项目链接](https://github.com/dreamerjackson/theWayToGolang)
* [作者知乎](https://www.zhihu.com/people/ke-ai-de-xiao-tu-ji-71)
* [blog](https://dreamerjonson.com/)
* [I still ❤️ you, GOPATH](https://divan.dev/posts/gopath/)
* [How to Write Go Code (with GOPATH)](https://golang.org/doc/gopath_code.html)
* [How to Write Go Code ](https://golang.org/doc/code.html#Organization)