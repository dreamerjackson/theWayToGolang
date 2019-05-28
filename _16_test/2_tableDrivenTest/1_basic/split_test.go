/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package split

import (
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

//单元测试
func TestSplit(t *testing.T) {
	got := Split("a/b/c", "/")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

//表格驱动测试
func TestSplit2(t *testing.T) {
	type test struct {
		input string
		sep   string
		want  []string
	}

	tests := []test{
		{input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
		{input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
		{input: "abc", sep: "/", want: []string{"abc"}},
	}

	for _, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
//表格驱动测试改进：匿名结构体
func TestSplit3(t *testing.T) {
	tests := []struct {
		input string
		sep   string
		want  []string
	}{
		{input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
		{input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
		{input: "abc", sep: "/", want: []string{"abc"}},
	}

	for _, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
//表格驱动测试改进：打印序号
func TestSplit4(t *testing.T) {
	tests := []struct {
		input string
		sep  string
		want  []string
	}{
		{input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
		{input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
		{input: "abc", sep: "/", want: []string{"abc"}},
		{input: "a/b/c/", sep: "/", want: []string{"a", "b", "c"}},
	}

	for i, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.want, got)
		}
	}
}

//表格驱动测试改进：给每一个测试名字
func TestSplit5(t *testing.T) {
	tests := []struct {
		name  string
		input string
		sep   string
		want  []string
	}{
		{name: "simple", input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
		{name: "wrong sep", input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
		{name: "no sep", input: "abc", sep: "/", want: []string{"abc"}},
		{name: "trailing sep", input: "a/b/c/", sep: "/", want: []string{"a", "b", "c"}},
	}

	for _, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}

//表格驱动测试改进：用map ，测试顺序不一定
func TestSplit6(t *testing.T) {
	tests := map[string]struct {
		input string
		sep   string
		want  []string
	}{
		"simple":       {input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
		"wrong sep":    {input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
		"no sep":       {input: "abc", sep: "/", want: []string{"abc"}},
		"trailing sep": {input: "a/b/c/", sep: "/", want: []string{"a", "b", "c"}},
	}

	for name, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}

///表格驱动测试改进：sub tests.  即便某一个失败，所有的也会测试
//可以执行某一个子用例：go test -run=.*/trailing -v
func TestSplit7(t *testing.T) {
	tests := map[string]struct {
		input string
		sep   string
		want  []string
	}{
		"simple":       {input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
		"wrong sep":    {input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
		"no sep":       {input: "abc", sep: "/", want: []string{"abc"}},
		"trailing sep": {input: "a/b/c/", sep: "/", want: []string{"a", "b", "c"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}

//表格驱动测试改进：%#v相应值的Go语法表示
/*
%#v的一个例子
func main() {
    type T struct {
        I int
    }
    x := []*T{{1}, {2}, {3}}
    y := []*T{{1}, {2}, {4}}
    fmt.Printf("%v %v\n", x, y)
    fmt.Printf("%#v %#v\n", x, y)
}

*/
func TestSplit8(t *testing.T) {
	tests := map[string]struct {
		input string
		sep   string
		want  []string
	}{
		"simple":       {input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
		"wrong sep":    {input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
		"no sep":       {input: "abc", sep: "/", want: []string{"abc"}},
		"trailing sep": {input: "a/b/c/", sep: "/", want: []string{"a", "b", "c"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %#v, got: %#v", tc.want, got)
			}
		})
	}
}

//表格驱动测试改进：cmp.Diff 替代 reflect.DeepEqual
/*
The goal of the cmp library is it is specifically to compare two values.
This is similar to reflect.DeepEqual, but it has more capabilities. Using the cmp pacakge you can, of course, write:
func main() {
    type T struct {
        I int
    }
    x := []*T{{1}, {2}, {3}}
    y := []*T{{1}, {2}, {4}}
    fmt.Println(cmp.Equal(x, y)) // false
}

But far more useful for us with our test function is the
cmp.Diff function which will produce a textual description of what is different between the two values, recursively.
func main() {
    type T struct {
        I int
    }
    x := []*T{{1}, {2}, {3}}
    y := []*T{{1}, {2}, {4}}
    diff := cmp.Diff(x, y)
    fmt.Printf(diff)
}
*/
func TestSplit9(t *testing.T) {
	tests := map[string]struct {
		input string
		sep   string
		want  []string
	}{
		"simple":       {input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
		"wrong sep":    {input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
		"no sep":       {input: "abc", sep: "/", want: []string{"abc"}},
		"trailing sep": {input: "a/b/c/", sep: "/", want: []string{"a", "b", "c"}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

