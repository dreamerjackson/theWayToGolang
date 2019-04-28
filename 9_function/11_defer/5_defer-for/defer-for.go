/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */



 package main
//由于defer会在函数结束时执行，所以循环里的defer是非常危险的，下面的程序可能会因为用完了文件描述符而崩溃。
/*
for _, filename := range filenames {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close() // NOTE: risky; could run out of file descriptors
// ...process f...
}
*/

//上面的程序可以修改为,包裹在一个函数中：
/*
for _, filename := range filenames {
	if err := doFile(filename); err != nil {
	return err
	}
}
func doFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	// ...process f...
}
*/