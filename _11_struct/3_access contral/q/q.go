/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "theWayToGolang/_11_struct/3_access contral/p"


//通过大些字母进行权限控制
func main(){

	var _ = pp.T{A: 1} // var _ = pp.T{A: 1, b: 2}  compile error: can't reference a, b

}
