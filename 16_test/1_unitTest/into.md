终端的项目目录下运行go test -v就可以看到测试结果了。 -v代表打印详细信息

```
➜  hello go test -v
=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
        main_test.go:26: the result is ok
PASS
ok      flysnow.org/hello       0.007s
```


1、含有单元测试代码的go文件必须以_test.go结尾，Go语言测试工具只认符合这个规则的文件
2、单元测试文件名_test.go前面的部分最好是被测试的方法所在go文件的文件名，比如例子中是add_test.go，因为测试的Add函数，在add.go文件里
3、单元测试的函数名必须以Test开头，是可导出公开的函数
4、测试函数的签名必须接收一个指向testing.T类型的指针，并且不能返回任何值
5、函数名最好是Test+要测试的方法函数名，比如例子中是TestAdd，表示测试的是Add这个这个函数