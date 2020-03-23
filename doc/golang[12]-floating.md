# golang快速入门[8.3]-深入理解浮点数

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

## 前言
* 在上文中我们学习了go语言中的自动类型推断
* 我们将在本文中深入理解go语言浮点数的存储细节
* 下面的一段简单程序 0.3 + 0.6 结果是什么？有人会天真的认为是0.9，但实际输出却是0.8999999999999999（go 1.13.5）

```go
		var f1 float64 = 0.3
		var f2 float64 = 0.6
		fmt.Println(f1 + f2)
```

* 问题在于大多数小数表示成二进制之后是近似且无限的。 以0.1为例。它可能是你能想到的最简单的十进制之一，但是二进制看起来却非常复杂：0.0001100110011001100...   他是一串连续循环无限的数字（关于如何转换为二进制数以后介绍）。
* 结果的荒诞性告诉我们，必须深入理解浮点数在计算机中的存储方式及其性质，才能正确处理数字的计算。
* golang 与其他很多语言（C、C++、Python）一样，使用了IEEE-754标准存储浮点数。
##  IEEE-754 如何存储浮点数
* IEEE-754规范使用特殊的以2为基数的科学表示法表示浮点数。
```
| 基本的10进制数字 | 科学计数法表示         | 指数表示        |     系数     | 底数 |    指数   |     小数  |
|----------------|---------------------|----------------|-------------|------|----------|----------|
| 700            | 7e+2                | 7 * 10^2       | 7           | 10   | 2        | 0        |
| 4,900,000,000  | 4.9e+9              | 4.9 * 10^9     | 4.9         | 10   | 9        | .9       |
| 5362.63        | 5.36263e+3          | 5.36263 * 10^3 | 5.36263     | 10   | 3        | .36263   |
| -0.00345       | 3.45e-3             | 3.45 * 10^-3   | 3.45        | 10   | -3       | .45      |
| 0.085          | 1.36e-4             | 1.36 * 2^-4    | 1.36        | 2    | -4       | .36      |
```

* 32位的单精度浮点数 与 64位的双精度浮点数的差异
```
| 精度              | 符号位  |  指数位     |  小数位        |偏移量|
|------------------|--------|------------|---------------|------|
| Single (32 Bits) | 1 [31] | 8 [30-23]  | 23 [22-00]    | 127  |
| Double (64 Bits) | 1 [63] | 11 [62-52] | 52 [51-00]    | 1023 |
```
* 符号位： 1 为 负数， 0 为正数。
* 指数位： 存储 指数减去偏移量，偏移量是为了表达负数而设计的。
* 小数位： 存储系数的小数位的准确或者最接近的值。
* 以 数字 0.085 为例。
```
| 符号位 | 指数位(123)    | 小数位 (.36)                  |
|------|----------------|------------------------------|
| 0    | 0111 1011      | 010 1110 0001 0100 0111 1011 |
```
##  小数位的计算
* 以0.36 为例: 010 1110 0001 0100 0111 1011 = 0.36  (第一位数字代表1/2,第二位数字是1/4 ...)
* 分解后的计算步骤为:
```
| Bit | Value   | Fraction  | Decimal          | Total            |
|-----|---------|-----------|------------------|------------------|
| 2   | 4       | 1⁄4       | 0.25             | 0.25             |
| 4   | 16      | 1⁄16      | 0.0625           | 0.3125           |
| 5   | 32      | 1⁄32      | 0.03125          | 0.34375          |
| 6   | 64      | 1⁄64      | 0.015625         | 0.359375         |
| 11  | 2048    | 1⁄2048    | 0.00048828125    | 0.35986328125    |
| 13  | 8192    | 1⁄8192    | 0.0001220703125  | 0.3599853515625  |
| 17  | 131072  | 1⁄131072  | 0.00000762939453 | 0.35999298095703 |
| 18  | 262144  | 1⁄262144  | 0.00000381469727 | 0.3599967956543  |
| 19  | 524288  | 1⁄524288  | 0.00000190734863 | 0.35999870300293 |
| 20  | 1048576 | 1⁄1048576 | 0.00000095367432 | 0.35999965667725 |
| 22  | 4194304 | 1⁄4194304 | 0.00000023841858 | 0.35999989509583 |
| 23  | 8388608 | 1⁄8388608 | 0.00000011920929 | 0.36000001430512 |
```

##  go语言显示浮点数 -  验证之前的理论
* math.Float32bits 可以为我们打印出数字的二进制表示。
* 下面的go代码输出0.085的二进制表达。
* 为了验证之前理论的正确性，根据二进制表示反向推导出其所表示的原始十进制0.085
```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var number float32 = 0.085
	fmt.Printf("Starting Number: %f\n\n", number)
	// Float32bits returns the IEEE 754 binary representation
	bits := math.Float32bits(number)

	binary := fmt.Sprintf("%.32b", bits)

	fmt.Printf("Bit Pattern: %s | %s %s | %s %s %s %s %s %s\n\n",
		binary[0:1],
		binary[1:5], binary[5:9],
		binary[9:12], binary[12:16], binary[16:20],
		binary[20:24], binary[24:28], binary[28:32])

	bias := 127
	sign := bits & (1 << 31)
	exponentRaw := int(bits >> 23)
	exponent := exponentRaw - bias

	var mantissa float64
	for index, bit := range binary[9:32] {
		if bit == 49 {
			position := index + 1
			bitValue := math.Pow(2, float64(position))
			fractional := 1 / bitValue
			mantissa = mantissa + fractional
		}
	}

	value := (1 + mantissa) * math.Pow(2, float64(exponent))

	fmt.Printf("Sign: %d Exponent: %d (%d) Mantissa: %f Value: %f\n\n",
		sign,
		exponentRaw,
		exponent,
		mantissa,
		value)
}

```

* 输出：
```
Starting Number: 0.085000
Bit Pattern: 0 | 0111 1011 | 010 1110 0001 0100 0111 1011
Sign: 0 Exponent: 123 (-4) Mantissa: 0.360000 Value: 0.085000
```

##  经典问题：如何判断一个浮点数其实存储的是整数
* 思考10秒钟....
* 下面是一段判断浮点数是否为整数的go代码实现，我们接下来逐行分析函数。它可以加深对于浮点数的理解
```go
func IsInt(bits uint32, bias int) {
    exponent := int(bits >> 23) - bias - 23
    coefficient := (bits & ((1 << 23) - 1)) | (1 << 23)
    intTest := (coefficient & (1 << uint32(-exponent) - 1))

    fmt.Printf("\nExponent: %d Coefficient: %d IntTest: %d\n",
        exponent,
        coefficient,
        intTest)

    if exponent < -23 {
        fmt.Printf("NOT INTEGER\n")
        return
    }

    if exponent < 0 && intTest != 0 {
        fmt.Printf("NOT INTEGER\n")
        return
    }

    fmt.Printf("INTEGER\n")
}
```

* 要保证是整数，一个重要的条件是必须要指数位大于127，如果指数位为127，代表指数为0. 指数位大于127，代表指数大于0， 反之小于0.下面我们以数字234523为例子：
```
Starting Number: 234523.000000
Bit Pattern: 0 | 1001 0000 | 110 0101 0000 0110 1100 0000
Sign: 0 Exponent: 144 (17) Mantissa: 0.789268 Value: 234523.000000
Exponent: -6 Coefficient: 15009472 IntTest: 0
INTEGER

```

* 第一步,计算指数。 由于  多减去了23，所以在第一个判断中 判断条件为  exponent < -23
```
exponent := int(bits >> 23) - bias - 23
```
* 第二步,`(bits & ((1 << 23) - 1))` 计算小数位。
```
coefficient := (bits & ((1 << 23) - 1)) | (1 << 23)

Bits:                   01001000011001010000011011000000
(1 << 23) - 1:          00000000011111111111111111111111
bits & ((1 << 23) - 1): 00000000011001010000011011000000
```

* `| (1 << 23)`` 代表 将1加在前方。
```
bits & ((1 << 23) - 1): 00000000011001010000011011000000
(1 << 23):              00000000100000000000000000000000
coefficient:            00000000111001010000011011000000
```

1 + 小数 = 系数。

* 第三步，计算intTest 只有当指数的倍数可以弥补最小的小数位的时候，才是一个整数。 如下，指数是17位，其不能够弥补最后6位的小数。即不能弥补1/2^18 的小数。 由于2^18位之后为0.所以是整数。
```
exponent:                     (144 - 127 - 23) = -6
1 << uint32(-exponent):       000000
(1 << uint32(-exponent)) - 1: 111111

coefficient:                 00000000111001010000011011000000
1 << uint32(-exponent)) - 1: 00000000000000000000000000111111
intTest:                     00000000000000000000000000000000

```

##  扩展阅读：概念：Normal number and denormal (or subnormal) number
* wiki的解释是：
```
In computing, a normal number is a non-zero number in a floating-point representation which is within the balanced range supported by a given floating-point format: it is a floating point number that can be represented without leading zeros in its significand.
```
什么意思呢？在IEEE-754中指数位有一个偏移量，偏移量是为了表达负数而设计的。 比如单精度中的0.085，实际的指数是 -3， 存储到指数位是123。
所以表达的负数就是有上限的。这个上限就是2^-126。 如果比这个负数还要小，例如2^-127,这个时候应该表达为0.1 * 2 ^ -126.   这时系数变为了不是1为前导的数，这个数就叫做denormal (or subnormal) number。
正常的系数是以1为前导的数就叫做Normal number。
##  扩展阅读：概念：精度
精度是一个非常复杂的概念，
在这里笔者讨论的是2进制浮点数的10进制精度。
精度为d表示的是在一个范围内，如果我们将d位10进制（按照科学计数法表达）转换为二进制。再将二进制转换为d位10进制。数据不损失意味着在此范围内是有d精度的。
精度的原因在于，数据在进制之间相互转换时，是不能够精准匹配的，而是匹配到一个最近的数。
在这里暂时不深入探讨，而是给出结论：
float32的精度为6-8位，
float64的精度为15-17位
并且精度是动态变化的，不同的范围可能有不同的精度。这里简单提示一下是由于 2的幂 与 10的幂之间的交错是不同的。

## 参考资料
* [项目链接](https://github.com/dreamerjackson/theWayToGolang)
* [作者知乎](https://www.zhihu.com/people/ke-ai-de-xiao-tu-ji-71)
* [blog](https://dreamerjonson.com/)
* [Type inference](https://en.wikipedia.org/wiki/Type_inference)
* [Rob Pike:Less is exponentially more](https://commandcenter.blogspot.com/2012/06/)
* [Type inference for go](http://fileadmin.cs.lth.se/cs/Education/EDAN70/CompilerProjects/2015/Reports/GigovicMalmros.pdf)
