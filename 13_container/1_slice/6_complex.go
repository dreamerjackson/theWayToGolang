/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"

func f(s []string, level int) {
	if level > 5 {
		return
	}
	s = append(s, fmt.Sprint(level))
	f(s, level+1)
	fmt.Println("level:", level, "slice:", s)
}

func main() {
	f(nil, 0)
}

/*
输出结果：

level: 5 slice: [0 1 2 3 4 5]
level: 4 slice: [0 1 2 3 4]
level: 3 slice: [0 1 2 3]
level: 2 slice: [0 1 2]
level: 1 slice: [0 1]
level: 0 slice: [0]


参考资料：
https://dave.cheney.net/2018/07/12/slices-from-the-ground-up
*/