## 引言

下面的一段简单程序 0.3 + 0.6 结果是什么？有人会天真的认为是0.9，但实际输出却是0.8999999999999999（go 1.13.5）

```go
		var f1 float64 = 0.3
		var f2 float64 = 0.6
		fmt.Println(f1 + f2)
```

问题在于大多数小数表示成二进制之后是近似且无限的。 以0.1为例。它可能是你能想到的最简单的十进制之一，但是二进制看起来却非常复杂：
0.0001100110011001100...   他是一串连续循环无限的数字（关于如何转换为二进制数以后介绍）。
结果的荒诞性告诉我们，必须深入理解浮点数在计算机中的存储方式及其性质，才能正确处理数字的计算。
golang 与其他很多语言（C、C++、Python）一样，使用了IEEE-754标准存储浮点数。
##  IEEE-754 如何存储浮点数
（How does the IEEE-754 standard store a floating point number in a binary format?）
IEEE-754规范使用特殊的以2为基数的科学表示法表示浮点数。
```
| 基本的10进制数字 | 科学计数法表示 | 指数表示   | 系数 | 底数  | 指数  | 小数 |
|----------------|---------------------|---------------|-------------|------|----------|----------|
| 700            | 7e+2                | 7 * 10^2       | 7           | 10   | 2        | 0        |
| 4,900,000,000  | 4.9e+9              | 4.9 * 10^9     | 4.9         | 10   | 9        | .9       |
| 5362.63        | 5.36263e+3          | 5.36263 * 10^3 | 5.36263     | 10   | 3        | .36263   |
| -0.00345       | 3.45e-3             | 3.45 * 10^-3   | 3.45        | 10   | -3       | .45      |
| 0.085          | 1.36e-4             | 1.36 * 2^-4    | 1.36        | 2    | -4       | .36      |
```

32位的单精度浮点数 与 64位的双精度浮点数的差异
```
| 精度             | 符号位 |  指数位   |  小数位        |偏移量|
|------------------|--------|------------|---------------|------|
| Single (32 Bits) | 1 [31] | 8 [30-23]  | 23 [22-00]    | 127  |
| Double (64 Bits) | 1 [63] | 11 [62-52] | 52 [51-00]    | 1023 |
```
符号位： 1 为 负数， 0 为正数。
指数位： 存储 指数减去偏移量，偏移量是为了表达负数而设计的。
小数位： 存储系数的小数位的准确或者最接近的值。

以 数字 0.085 为例。
```
| 符号位 | 指数位(123)  | 小数位 (.36)                 |
|------|----------------|------------------------------|
| 0    | 0111 1011      | 010 1110 0001 0100 0111 1011 |
```
##  小数位的计算
以0.36 为例:
010 1110 0001 0100 0111 1011 = 0.36  (第一位数字代表1/2,第二位数字是1/4 ...)
分解后的计算步骤为:
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
math.Float32bits 可以为我们打印出数字的二进制表示。
下面的go代码输出0.085的二进制表达。
为了验证之前理论的正确性，根据二进制表示反向推导出其所表示的原始十进制0.085
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

输出：
```
Starting Number: 0.085000
Bit Pattern: 0 | 0111 1011 | 010 1110 0001 0100 0111 1011
Sign: 0 Exponent: 123 (-4) Mantissa: 0.360000 Value: 0.085000
```



##  经典问题：如何判断一个浮点数其实存储的是整数

思考10秒钟....

下面是一段判断浮点数是否为整数的go代码实现，我们接下来逐行分析函数。
它可以加深对于浮点数的理解
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

1、要保证是整数，一个重要的条件是必须要指数位大于127，如果指数位为127，代表指数为0. 指数位大于127，代表指数大于0， 反之小于0.

下面我们以数字234523为例子：
```
Starting Number: 234523.000000
Bit Pattern: 0 | 1001 0000 | 110 0101 0000 0110 1100 0000
Sign: 0 Exponent: 144 (17) Mantissa: 0.789268 Value: 234523.000000
Exponent: -6 Coefficient: 15009472 IntTest: 0
INTEGER

```

第一步,计算指数。 由于  多减去了23，所以在第一个判断中 判断条件为  exponent < -23
    exponent := int(bits >> 23) - bias - 23

第二步，
(bits & ((1 << 23) - 1)) 计算小数位。
```
coefficient := (bits & ((1 << 23) - 1)) | (1 << 23)

Bits:                   01001000011001010000011011000000
(1 << 23) - 1:          00000000011111111111111111111111
bits & ((1 << 23) - 1): 00000000011001010000011011000000
```
| (1 << 23) 代表 将1加在前方。
```
bits & ((1 << 23) - 1): 00000000011001010000011011000000
(1 << 23):              00000000100000000000000000000000
coefficient:            00000000111001010000011011000000
```

1 + 小数 = 系数。

第三步，计算intTest 只有当指数的倍数可以弥补最小的小数位的时候，才是一个整数。
如下，指数是17位，其不能够弥补最后6位的小数。即不能弥补1/2^18 的小数。
由于2^18位之后为0.所以是整数。
```
exponent:                     (144 - 127 - 23) = -6
1 << uint32(-exponent):       000000
(1 << uint32(-exponent)) - 1: 111111

coefficient:                 00000000111001010000011011000000
1 << uint32(-exponent)) - 1: 00000000000000000000000000111111
intTest:                     00000000000000000000000000000000

```




##  概念：Normal number and denormal (or subnormal) number
wiki的解释是：
```
In computing, a normal number is a non-zero number in a floating-point representation which is within the balanced range supported by a given floating-point format: it is a floating point number that can be represented without leading zeros in its significand.
```
什么意思呢？在IEEE-754中指数位有一个偏移量，偏移量是为了表达负数而设计的。 比如单精度中的0.085，实际的指数是 -3， 存储到指数位是123。
所以表达的负数就是有上限的。这个上限就是2^-126。 如果比这个负数还要小，例如2^-127,这个时候应该表达为0.1 * 2 ^ -126.   这时系数变为了不是1为前导的数，这个数就叫做denormal (or subnormal) number。
正常的系数是以1为前导的数就叫做Normal number。
##  概念：精度
    精度是一个非常复杂的概念，
    在这里笔者讨论的是2进制浮点数的10进制精度。
    精度为d表示的是在一个范围内，如果我们将d位10进制（按照科学计数法表达）转换为二进制。再将二进制转换为d位10进制。数据不损失意味着在此范围内是有d精度的。
    精度的原因在于，数据在进制之间相互转换时，是不能够精准匹配的，而是匹配到一个最近的数。如图所示：
    <center>{% asset_img 1.png %}</center>
    <center><div>精度转换</div></center>

    在这里暂时不深入探讨，而是给出结论：
    float32的精度为6-8位，
    float64的精度为15-17位
    并且精度是动态变化的，不同的范围可能有不同的精度。这里简单提示一下是由于 2的幂 与 10的幂之间的交错是不同的。
##  golang decimal 包详解
链接:https://github.com/shopspring/decimal
取名叫decimal，顾名思意是将浮点数转换为10进制表达。decimal包使用math/big包存储大整数并进行大整数的计算。比如对于字符串 "123.45" 我们可以将其转换为12345这个大整数，以及-2代表指数。参考decimal结构体：
```
type Decimal struct {
	value *big.Int
	exp int32
}
```

在本小节中，我们不会探讨math/big 是如何进行大整数运算的，而是探讨decimal包一个非常重要的函数：
NewFromFloat(value float64) Decimal ，其主要调用了下面的函数：
```
func newFromFloat(val float64, bits uint64, flt *floatInfo) Decimal {
	if math.IsNaN(val) || math.IsInf(val, 0) {
		panic(fmt.Sprintf("Cannot create a Decimal from %v", val))
	}
	exp := int(bits>>flt.mantbits) & (1<<flt.expbits - 1)
	mant := bits & (uint64(1)<<flt.mantbits - 1)
	switch exp {
	case 0:
		exp++
	default:
		mant |= uint64(1) << flt.mantbits
	}
	exp += flt.bias
	var d decimal
	d.Assign(mant)
	d.Shift(exp - int(flt.mantbits))
	d.neg = bits>>(flt.expbits+flt.mantbits) != 0
	roundShortest(&d, mant, exp, flt)
	if d.nd < 19 {
		tmp := int64(0)
		m := int64(1)
		for i := d.nd - 1; i >= 0; i-- {
			tmp += m * int64(d.d[i]-'0')
			m *= 10
		}
		if d.neg {
			tmp *= -1
		}
		return Decimal{value: big.NewInt(tmp), exp: int32(d.dp) - int32(d.nd)}
	}
	dValue := new(big.Int)
	dValue, ok := dValue.SetString(string(d.d[:d.nd]), 10)
	if ok {
		return Decimal{value: dValue, exp: int32(d.dp) - int32(d.nd)}
	}
	return NewFromFloatWithExponent(val, int32(d.dp)-int32(d.nd))
}
```

此函数会将浮点数转换为Decimal结构。
读者想象一下这个问题：如果传递的浮点数value 例如0.1 本身就是一个近似数。 为什么decimal包能够解决计算的准确性？
这就是NewFromFloat为我们做的事情。
下面我将对此函数做逐行分析。

```
	//2-4行判断浮点数有效性，不能为NAN或INF
    if math.IsNaN(val) || math.IsInf(val, 0) {
		panic(fmt.Sprintf("Cannot create a Decimal from %v", val))
	}
```

第5行：剥离出IEEE浮点数的指数位
    exp := int(bits>>flt.mantbits) & (1<<flt.expbits - 1)

第6行：剥离出浮点数的系数的小数位
	mant := bits & (uint64(1)<<flt.mantbits - 1)

第7行：如果是指数位为0，代表浮点数是denormal (or subnormal) number；
默认情况下会在mant之前加上1，因为mant只是系数的小数，在前面加上1后，代表真正的小数位。
现在 mant = IEEE浮点数系数 * 2^53

第13行： 加上偏移量，exp现在代表真正的指数。
第14行： 引入了一个中间结构`decimal`
```
type decimal struct {
	d     [800]byte
	nd    int
	dp    int
	neg   bool
	trunc bool
}
```

第15行： 调用d.Assign(mant) , 将mant作为10进制数，存起来。10进制数的每一位都作为一个字符存储到
decimal的byte数组中
```
func (a *decimal) Assign(v uint64) {
	var buf [24]byte

	// Write reversed decimal in buf.
	n := 0
	for v > 0 {
		v1 := v / 10
		v -= 10 * v1
		buf[n] = byte(v + '0')
		n++
		v = v1
	}

	// Reverse again to produce forward decimal in a.d.
	a.nd = 0
	for n--; n >= 0; n-- {
		a.d[a.nd] = buf[n]
		a.nd++
	}
	a.dp = a.nd
	trim(a)
}
```

第16行： 调用shift函数，这个函数非常难理解。
```
func (a *decimal) Shift(k int) {
	switch {
	case a.nd == 0:
	case k > 0:
		for k > maxShift {
			leftShift(a, maxShift)
			k -= maxShift
		}
		leftShift(a, uint(k))
	case k < 0:
		for k < -maxShift {
			rightShift(a, maxShift)
			k += maxShift
		}
		rightShift(a, uint(-k))
	}
}
```

此函数的功能是为了获取此浮点数代表的10进制数据的有效位个数以及小数位个数。
exp是真实的指数，其也是能够覆盖小数部分2进制位的个数。（参考前面如何判断浮点数是整数）
exp - int(flt.mantbits)代表不能被exp覆盖的2进制位的个数
如果exp - int(flt.mantbits)  > 0  代表exp能够完全覆盖小数位 因此 浮点数是一个非常大的整数。这时会调用leftShift(a, uint(k))。否则将调用rightShift(a, uint(-k)), 常规rightShift会调用得更多。因此我们来看看rightShift函数的实现。

第5行： 此for循环将计算浮点数10进制表示的小数部分的有效位为 r-1 。
        n >> k  是一个重要的衡量指标，代表了小数部分与整数部分的分割

第21行：此时整数部分所占的有效位数为a.dp -=（r-1）
第24行：这两个循环做了2件事情：
    1、计算10进制表示的有效位数
    2、将10进制表示存入bytes数组中。例如对于浮点数64.125，现在byte数组存储的前5位就是64125

```
func rightShift(a *decimal, k uint) {
	r := 0
	w := 0
	var n uint
	for ; n>>k == 0; r++ {
		if r >= a.nd {
			if n == 0 {
				a.nd = 0
				return
			}
			for n>>k == 0 {
				n = n * 10
				r++
			}
			break
		}
		c := uint(a.d[r])
		n = n*10 + c - '0'
	}
	// 整数部分的有效位数
	a.dp -= r - 1

	var mask uint = (1 << k) - 1
    // 整数部分
	for ; r < a.nd; r++ {
		c := uint(a.d[r])
		dig := n >> k
		n &= mask
		a.d[w] = byte(dig + '0')
		w++
		n = n*10 + c - '0'
	}
    // 小数部分
	for n > 0 {
		dig := n >> k
		n &= mask
		if w < len(a.d) {
			a.d[w] = byte(dig + '0')
			w++
		} else if dig > 0 {
			a.trunc = true
		}
		n = n * 10
	}
    // 有效位
	a.nd = w
	trim(a)
}
```

继续回到newFromFloat函数，第18行，调用了roundShortest函数，
此函数非常关键。其会将浮点数转换为离其最近的十进制数。
这是为什么decimal.NewFromFloat(0.1)能够精准表达0.1的原因。

参考上面的精度，此函数主要考察了2的幂与10的幂之间的交错关系。四舍五入到最接近的10净值。
未完待续，下面只提示几点：
1、2^exp <= d < 10^dp。
2、10进制数之间至少相聚10^(dp-nd)
3、2的幂之间的最小间距至少为2^(exp-mantbits)
4、什么时候d就是最接近2进制的10进制数？
如果10^(dp-nd) > 2^(exp-mantbits)，表明 当十进制下降一个最小位数时，匹配到的是更小的数字  value -  2^(exp-mantbits)，所以是最小的。

```
func roundShortest(d *decimal, mant uint64, exp int, flt *floatInfo) {
		if mant == 0 {
			d.nd = 0
			return
		}
        // d 是否就是最接近的2进制数。
		minexp := flt.bias + 1 // minimum possible exponent
		if exp > minexp && 332*(d.dp-d.nd) >= 100*(exp-int(flt.mantbits)) {
			// The number is already shortest.
			return
		}
        // 计算最接近的大于d的10进制数
		upper := new(decimal)
		upper.Assign(mant*2 + 1)
		upper.Shift(exp - int(flt.mantbits) - 1)

		var mantlo uint64
		var explo int
		if mant > 1<<flt.mantbits || exp == minexp {
			mantlo = mant - 1
			explo = exp
		} else {
			mantlo = mant*2 - 1
			explo = exp - 1
		}
        // 计算最接近的小于d的10进制数
		lower := new(decimal)
		lower.Assign(mantlo*2 + 1)
		lower.Shift(explo - int(flt.mantbits) - 1)

		inclusive := mant%2 == 0

		for i := 0; i < d.nd; i++ {
			l := byte('0') // lower digit
			if i < lower.nd {
				l = lower.d[i]
			}
			m := d.d[i]    // middle digit
			u := byte('0') // upper digit
			if i < upper.nd {
				u = upper.d[i]
			}

			okdown := l != m || inclusive && i+1 == lower.nd

			okup := m != u && (inclusive || m+1 < u || i+1 < upper.nd)
		switch {
		case okdown && okup:
			d.Round(i + 1)
			return
		case okdown:
			d.RoundDown(i + 1)
			return
		case okup:
			d.RoundUp(i + 1)
			return
		}
	}
}

```


继续回到newFromFloat函数，第19行  如果精度小于19，是在int64范围内的，可以使用快速路径，否则使用math/big包进行赋值操作，效率稍微要慢一些。

第36行，正常情况几乎不会发生。如果setstring在异常的情况下会调用NewFromFloatWithExponent 指定精度进行四舍五入截断。

##  参考资料

qq交流群：713385260
未完，待续...

[Why 0.1 Does Not Exist In Floating-Point](https://www.exploringbinary.com/why-0-point-1-does-not-exist-in-floating-point/)

[Normal number](https://en.wikipedia.org/wiki/Normal_number)

[7-bits-are-not-enough-for-2-digit-accuracy](https://www.exploringbinary.com/7-bits-are-not-enough-for-2-digit-accuracy/)

[Decimal Precision of Binary Floating-Point Numbers](https://www.exploringbinary.com/decimal-precision-of-binary-floating-point-numbers/)
