/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package add_test

import (
	"testing"
	"theWayToGolang/16_test/1_unitTest"
)

func TestAdd(t *testing.T) {
	sum := add.Add(1,2)
	if sum == 3 {
		t.Log("the result is ok")
	} else {
		t.Fatal("the result is wrong")
	}
}