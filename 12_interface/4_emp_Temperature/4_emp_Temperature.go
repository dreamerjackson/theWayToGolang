/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"fmt"
	"flag"
)
//flag包中的Value是一个接口，其中的set方法表明命令行参数应该如何解析
//如  ./4_emp_Temperature -temp -18C
//会将-18c传递到set函数中解析，解析-18C为 -18 并返回

/*
$ ./4_emp_Temperature -temp 212°F
100°C
$ ./4_emp_Temperature -temp 273.15K
invalid value "273.15K" for flag -temp: invalid temperature "273.15K"
Usage of ./4_emp_Temperature:
-temp value
the temperature (default 20°C)
$ ./4_emp_Temperature -help
Usage of ./4_emp_Temperature:
-temp value
the temperature (default 20°C)
*/

/*
//!+flagvalue
package flag

// Value is the interface to the value stored in a flag.
type Value interface {
	String() string
	Set(string) error
}
//!-flagvalue
*/

type Celsius float64
type Fahrenheit float64

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9.0/5.0 + 32.0) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

type celsiusFlag struct{Celsius}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}



var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}