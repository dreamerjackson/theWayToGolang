# golang快速入门[5.1]-go语言是如何运行的-链接器

## 前文
* [golang快速入门[2.1]-go语言开发环境配置-windows](https://zhuanlan.zhihu.com/p/107659334)
* [golang快速入门[2.2]-go语言开发环境配置-macOS](https://zhuanlan.zhihu.com/p/107661202)
* [golang快速入门[2.3]-go语言开发环境配置-linux](https://zhuanlan.zhihu.com/p/107662649)
* [golang快速入门[3]-go语言helloworld](https://zhuanlan.zhihu.com/p/107664129)
* [golang快速入门[4]-go语言如何编译为机器码](https://zhuanlan.zhihu.com/p/107665043)

## 前言
* 在上一篇文章中,我们详细介绍了go语言编译为机器码经历的：词法分析 => 语法分析 => 类型检查 => 中间代码 => 代码优化 => 生成机器码
* 但是在源代码生成执行程序的过程中，其实还经历了链接等过程。总的来说一个程序的生命周期可以概括为: 编写代码 => 编译 => 链接 => 加载到内存 => 执行
* 在第5章我们将对其进行逐一解释

## 链接(link)
* 我们编写的程序可能会使用其他程序或程序库( library ) 正如我们在helloworld程序中使用的fmt package
* 我们编写的程序必须与这些程序或程序库一起才能够执行
* 链接是将我们编写的程序与我们需要的外部程序组合在一起的过程
* 链接器是系统软件，在系统开发中起着至关重要的作用，因为它可以进行单独的编译。您可以将它分解为更小，更易管理的块，然后分别进行修改和编译，而不是将一个大型应用程序组织为一个整体的源文件。当您更改其中一个模块时，只需重新编译它并重新链接应用程序，而无需重新编译其他源文件。
* 链接分为两种，静态链接与动态链接
* 静态链接的特点在于链接器会将程序中使用的所有库程序复制到最后的可执行文件中。而动态链接只会在最后的可执行文件中存储动态链接库的位置，并在运行时调用。
* 因此静态链接要更快，可移植，因为它不需要在运行它的系统上存在该库。但是在磁盘和内存上占用更多的空间
* 链接发生的过程会在两个地方，一种是静态链接会在编译时的最后一步发生，一种是动态链接在程序加载到内存时发生。

* 下面我们简单对比一下静态链接与动态链接

| 静态链接 | 动态链接 |
|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| 静态链接是将程序中使用的所有库模块复制到最终可执行文件的过程,这是由链接器执行的，并且是编译过程的最后一步。 加载程序后，操作系统会将包含可执行代码和数据的单个文件放入内存,该静态链接文件包括调用程序和被调用程序. | 在动态链接中，外部库（共享库）的地址放置在最终的可执行文件中，而实际链接是在运行时将可执行文件和库都放置在内存中时进行的,动态链接使多个程序可以使用可执行模块的单个副本。 |
| 静态链接由称为链接器的程序执行，是编译程序的最后一步 | 动态链接由操作系统在运行时执行 |
| 静态链接文件的大小明显更大，因为外部程序内置在可执行文件中 | 在动态链接中，共享库中只有一个副本保留在内存中。 这大大减小了可执行程序的大小，从而节省了内存和磁盘空间 |
| 在静态链接中，如果任何外部程序已更改，则必须重新编译并重新链接它们，否则更改将不会反映在现有的可执行文件中 | 在动态链接中不同，只需要更新和重新编译各个共享模块程序即可变动。这是动态链接所提供的最大优势之一 |
| 静态链接的程序每次将其加载到内存中执行时，都会花费恒定的加载时间 | 动态链接中，如果共享库代码已存在于内存中，则可以减少加载时间 |
| 使用静态链接库的程序通常比使用共享库的程序快 | 使用共享库的程序通常比使用静态链接库的程序要慢。 |
| 在静态链接程序中，所有代码都包含在一个可执行模块中。 因此，它们永远不会遇到兼容性问题。 | 动态链接的程序依赖于具有兼容的库。 如果更改了库（例如，新的编译器版本可能更改了库），则可能必须重新设计应用程序以使其与该库的新版本兼容。 如果从系统中删除了一个库，则使用该库的程序将不再起作用。 |

## go语言是静态链接还是动态链接？
* 有时会看到一些比较老的文章说go语言是静态链接的，但这种说法是不准确的
* 现在的go语言不仅支持静态链接也支持动态编译
* 总的来说，go语言在一般默认情况下是静态链接的，但是一些特殊的情况，例如使用了CGO（即引用了C代码）的地方，则会使用操作系统的动态链接库。例如go语言的`net/http`包在默认情况下会应用`libpthread`与 `libc` 的动态链接库，这种情况会导致go语言程序虚拟内存的增加（下一文介绍）
* go语言也支持在`go build`编译时传递参数来指定要生成的链接库的方式,我们可以使用`go help build`命令查看
```
» go help buildmode                                                                                                                                                             jackson@192
	-buildmode=archive
		Build the listed non-main packages into .a files. Packages named
		main are ignored.

	-buildmode=c-archive
		Build the listed main package, plus all packages it imports,
		into a C archive file. The only callable symbols will be those
		functions exported using a cgo //export comment. Requires
		exactly one main package to be listed.

	-buildmode=c-shared
		Build the listed main package, plus all packages it imports,
		into a C shared library. The only callable symbols will
		be those functions exported using a cgo //export comment.
		Requires exactly one main package to be listed.

	-buildmode=default
		Listed main packages are built into executables and listed
		non-main packages are built into .a files (the default
		behavior).

	-buildmode=shared
		Combine all the listed non-main packages into a single shared
		library that will be used when building with the -linkshared
		option. Packages named main are ignored.

	-buildmode=exe
		Build the listed main packages and everything they import into
		executables. Packages not named main are ignored.

	-buildmode=pie
		Build the listed main packages and everything they import into
		position independent executables (PIE). Packages not named
		main are ignored.

	-buildmode=plugin
		Build the listed main packages, plus all packages that they
		import, into a Go plugin. Packages not named main are ignored.
```
* archive:   将非 main package构建为 .a 文件. main 包将被忽略。
* c-archive: 将 main package构建为及其导入的所有package构建为构建到 C 归档文件中
* c-shared:  将mainpackage构建为，以及它们导入的所有package构建到C 动态库中。
* shared:    将所有非 main package合并到一个动态库中，当使用-linkshared参数后，能够使用此动态库
* exe:       将main package和其导入的package构建为成为可执行文件
* 本文不再介绍go如何手动使用动态库这一高级功能，读者只需现在知道go可以实现这一功能即可


## 编译与链接的具体过程
* 下面我们以helloworld程序为例，来说明go语言编译与链接的过程，我们可以使用`go build`命令，`-x`参数代表了打印执行的过程
```
go build  -x main.go
```
输出如下：
```
WORK=/var/folders/g2/0l4g444904vbn8wxnrw0j_980000gn/T/go-build757876739
mkdir -p $WORK/b001/
cat >$WORK/b001/importcfg << 'EOF' # internal
# import config
packagefile fmt=/usr/local/go/pkg/darwin_amd64/fmt.a
packagefile runtime=/usr/local/go/pkg/darwin_amd64/runtime.a
EOF
cd /Users/jackson/go/src/viper/XXX
/usr/local/go/pkg/tool/darwin_amd64/compile -o $WORK/b001/_pkg_.a -trimpath "$WORK/b001=>" -p main -complete -buildid JqleDuJlC1iLMVADicsQ/JqleDuJlC1iLMVADicsQ -goversion go1.13.6 -D _/Users/jackson/go/src/viper/args -importcfg $WORK/b001/importcfg -pack -c=4 ./main.go
/usr/local/go/pkg/tool/darwin_amd64/buildid -w $WORK/b001/_pkg_.a # internal
cp $WORK/b001/_pkg_.a /Users/jackson/Library/Caches/go-build/cf/cf0dc65f39f01c8494192fa8af14570b445f6a25b762edf0b7258c22d6e10dc8-d # internal
cat >$WORK/b001/importcfg.link << 'EOF' # internal
packagefile command-line-arguments=$WORK/b001/_pkg_.a
packagefile fmt=/usr/local/go/pkg/darwin_amd64/fmt.a
packagefile runtime=/usr/local/go/pkg/darwin_amd64/runtime.a
packagefile errors=/usr/local/go/pkg/darwin_amd64/errors.a
...
EOF
mkdir -p $WORK/b001/exe/
cd .
/usr/local/go/pkg/tool/darwin_amd64/link -o $WORK/b001/exe/a.out -importcfg $WORK/b001/importcfg.link -buildmode=exe -buildid=zCU3mCFNeUDzrRM33f4L/JqleDuJlC1iLMVADicsQ/r7xJ7p5GD5T9VONtmxob/zCU3mCFNeUDzrRM33f4L -extld=clang $WORK/b001/_pkg_.a
/usr/local/go/pkg/tool/darwin_amd64/buildid -w $WORK/b001/exe/a.out # internal
mv $WORK/b001/exe/a.out main
rm -r $WORK/b001/
```
* 下面我们对输出进行逐行分析
* 创建了一个临时目录，用于存放临时文件。默认情况下命令结束时自动删除此目录，如果需要保留添加`-work`参数。
```
WORK=/var/folders/g2/0l4g444904vbn8wxnrw0j_980000gn/T/go-build757876739
mkdir -p $WORK/b001/
cat >$WORK/b001/importcfg << 'EOF' # internal
```
* 生成编译配置文件，主要为编译过程需要的外部依赖（如：引用的其他包的函数定义）
```
# import config
packagefile fmt=/usr/local/go/pkg/darwin_amd64/fmt.a
packagefile runtime=/usr/local/go/pkg/darwin_amd64/runtime.a
```
* 编译，生成中间结果`$WORK/b001/_pkg_.a`,

```
/usr/local/go/pkg/tool/darwin_amd64/compile -o $WORK/b001/_pkg_.a -trimpath "$WORK/b001=>" -p main -complete -buildid JqleDuJlC1iLMVADicsQ/JqleDuJlC1iLMVADicsQ -goversion go1.13.6 -D _/Users/jackson/go/src/viper/args -importcfg $WORK/b001/importcfg -pack -c=4 ./main.go
```

* .a文件由compile命令生成，也可以通过[go tool compile](https://golang.org/cmd/compile/)进行调用
* .a类型的文件又叫做目标文件([object file](https://en.wikipedia.org/wiki/Object_code))，其是一个压缩包，内部包含了`__.PKGDEF`、`_go_.o` 两个文件，分别为编译目标文件和链接目标文件
```
$ file _pkg_.a # 检查文件格式
_pkg_.a: current ar archive # 说明是ar格式的打包文件
$ ar x _pkg_.a #解包文件
$ ls
__.PKGDEF  _go_.o
```

* 文件内容由代码导出的函数、变量以及引用的其他包的信息组成。为了弄清这两个文件包含的信息需要查看go编译器实现的相关代码，相关代码在`src/cmd/compile/internal/gc/obj.go`文件中（源码中的文件内容可能随版本更新变化，本系列文章以Go1.13.5版本为准）
* 下面代码中生成ar文件，ar文件 是一种非常简单的打包文件格式，广泛用于linux中静态链接库文件中，文件以 字符串`"!<arch>\n"`开头。随后跟着60字节的文件头部（包含文件名、修改时间等信息），之后跟着文件内容。因为ar文件格式简单，Go编译器直接在函数中实现了ar打包过程。
* startArchiveEntry用于预留ar文件头信息位置（60字节），finishArchiveEntry用于写入文件头信息，因为文件头信息中包含文件大小，在写入完成之前文件大小未知，所以分两步完成。
```
func dumpobj1(outfile string, mode int) {
	bout, err := bio.Create(outfile)
	if err != nil {
		flusherrors()
		fmt.Printf("can't create %s: %v\n", outfile, err)
		errorexit()
	}
	defer bout.Close()
	bout.WriteString("!<arch>\n")

	if mode&modeCompilerObj != 0 {
		start := startArchiveEntry(bout)
		dumpCompilerObj(bout)
		finishArchiveEntry(bout, start, "__.PKGDEF")
	}
	if mode&modeLinkerObj != 0 {
		start := startArchiveEntry(bout)
		dumpLinkerObj(bout)
		finishArchiveEntry(bout, start, "_go_.o")
	}
}

```
* 生成链接配置文件，主要为需要链接的其他依赖
```
cat >$WORK/b001/importcfg.link << 'EOF' # internal
packagefile command-line-arguments=$WORK/b001/_pkg_.a
packagefile fmt=/usr/local/go/pkg/darwin_amd64/fmt.a
packagefile runtime=/usr/local/go/pkg/darwin_amd64/runtime.a
packagefile errors=/usr/local/go/pkg/darwin_amd64/errors.a
...
EOF
```

* 执行链接器，生成最终可执行文件`main`,同时可执行文件会拷贝到当前路径，最后删除临时文件
```
/usr/local/go/pkg/tool/darwin_amd64/link -o $WORK/b001/exe/a.out -importcfg $WORK/b001/importcfg.link -buildmode=exe -buildid=zCU3mCFNeUDzrRM33f4L/JqleDuJlC1iLMVADicsQ/r7xJ7p5GD5T9VONtmxob/zCU3mCFNeUDzrRM33f4L -extld=clang $WORK/b001/_pkg_.a
/usr/local/go/pkg/tool/darwin_amd64/buildid -w $WORK/b001/exe/a.out # internal
mv $WORK/b001/exe/a.out main
rm -r $WORK/b001/
```

## 总结
* 在本文中，我们介绍了go程序从源代码到运行需要经历的重要一环——链接，并介绍了静态链接与动态链接
* 在本文中，我们用一个例子介绍了编译与链接的具体过程
* 在下文中，我们将介绍go语言的内存分配

## 参考资料
* [项目链接](https://github.com/dreamerjackson/theWayToGolang)
* [作者知乎](https://www.zhihu.com/people/ke-ai-de-xiao-tu-ji-71)
* [blog](https://dreamerjonson.com/)
* [wiki obj code](https://en.wikipedia.org/wiki/Object_code)
* [golang Command compile](https://golang.org/cmd/compile/)
* [golang Command Link](https://golang.org/cmd/link/)
* [初探 Go 的编译命令执行过程](https://halfrost.com/go_command/)
* [How does the go build command work ?](https://dave.cheney.net/2013/10/15/how-does-the-go-build-command-work)
* [Golang编译器漫谈（1）编译器和连接器](https://hao.io/2020/01/golang%e7%bc%96%e8%af%91%e5%99%a8%e6%bc%ab%e8%b0%88%ef%bc%881%ef%bc%89%e7%bc%96%e8%af%91%e5%99%a8%e5%92%8c%e8%bf%9e%e6%8e%a5%e5%99%a8/)
* [What are the differences between static and dynamic (shared) library linking?](http://cs-fundamentals.com/tech-interview/c/difference-between-static-and-dynamic-linking.php)

