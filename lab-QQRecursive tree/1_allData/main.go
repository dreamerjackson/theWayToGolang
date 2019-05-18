
/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unsafe"
)
var g_pp = make([]*[]*byte,10)
func main() {
	f, err := os.OpenFile("fib.txt", os.O_WRONLY, 0666)
	if err!=nil{
		log.Fatal(err)
	}
	file = bufio.NewWriter(f)
	memtoFile(&g_pp,11)
	file.Flush()
}
//初始化构建内存模型
func init(){
	fmt.Println("初始化开始")
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str:= scanner.Text()

		if len(str) <50{//数据已经整理过，最多50位。
			qq:= getQQ(str)
			//fmt.Println(qq)
			if len(qq)==10  && isAllNum(qq){
				//递归树，将模型构建完毕。
				assign(&g_pp,11,qq,str)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("初始化结束")
}


/*
最重要的递归函数，pp指针。deep为深度，str为字符串，password为密码
*/
func assign(pp *[]*[]*byte, deep int, str string,password string) {

	//最后一级指针时，开辟内存，存储密码。
	if deep == 1 {

		buf:= make([]*byte,10)
		(*pp)[getnum(str[10-deep])] = &buf
		p:= (*string)(unsafe.Pointer((*pp)[getnum(str[10-deep])]))
		*p = password
		return
	}

	//刚开始传递11级指针的地址，deep=11递归下去，可以省略，改为传递11级指针，deep=10。
	if deep == 11 {
		assign((*[]*[]*byte)(unsafe.Pointer(pp)), deep - 1, str,password);
		return
	}
	//判断该指针是否为空。如1131052403，当deep=10时，取出第一个数字1.判断pp[1]是否为Nil，为Nil就说明从来没有出现过第10位为1的qq号。
	//这时就会为pp[1]开辟10个指针大小的内存，初始化为空。
	//如果已经存在就继续递归。
	if (*pp)[getnum(str[10-deep])]!=nil{
		assign((*[]*[]*byte)(unsafe.Pointer((*pp)[getnum(str[10-deep])])), deep - 1, str,password);
	}else {
		buf:= make([]*byte,10)
		(*pp)[getnum(str[10-deep])] = &buf
		assign((*[]*[]*byte)(unsafe.Pointer((*pp)[getnum(str[10-deep])])), deep - 1, str,password);
	}
}

//字符转换为数字
func getnum(u uint8) uint8{
	return u - '0'
}

//qq补齐位数，判断是否为数字，字符转换为数字，数字不足补充0的算法
//获取QQ号 1131052403----qwerty
//截取数字10位，不足的补0，
//对于这个函数的改进，让我可以在查找qq函数时也可以用
func getQQ(s string) string{
	raw:= strings.Split(s,"----")[0]
	length := len(raw)
	if length < 10 {
		raw =  strings.Repeat("0",10-length) + raw
	}
	return raw
}

//判断qq全部为数字
func isAllNum(qq string ) bool{
	for _,ch := range qq{
		if ch < '0' || ch > '9'{
			return false
		}
	}
	return true
}



/*
根据初始化递归树的原理，我们可以明白判断是否存在的意义。
如果在某一位出现了一个数字，其所在的指针为NULL，就说明在该位从来没有出现过这个数字。
相反，如果存在该数字，就首先说明其每一位的指针都不为空。

*/
var flag = true
var findresult string = ""
func isExit(pp *[]*[]*byte, deep int, str string) {

	if deep == 1 {
		if (*pp)[getnum(str[10-deep])]!=nil{
			findresult = *(*string)(unsafe.Pointer((*pp)[getnum(str[10-deep])]))
		}
		return
	}

	if deep == 11 {
		isExit((*[]*[]*byte)(unsafe.Pointer(pp)), deep - 1, str);
		return
	}

	if flag && (*pp)[getnum(str[10-deep])]!=nil{
		isExit((*[]*[]*byte)(unsafe.Pointer((*pp)[getnum(str[10-deep])])), deep - 1, str);
	}else {

		flag = false
		return
	}
}

//查找qq号
func search(){
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		flag = true
		findresult=""

		qq:= input.Text()
		qq= getQQ(qq)
		fmt.Println("搜索qq号：",qq)
		isExit(&g_pp,11,qq)
		if isAllNum(qq) && findresult!=""{
			fmt.Printf("结果为：%s\n",findresult)
		}
	}
}


func add(qq string){


	qqslice := strings.Split(qq,"----")

	if len(qqslice) <2{
		log.Fatal("请输入正确的qq号，例如：1131052403----password")
		return
	}
	qqmian:= qqslice[0]
	qqpassword:= qqslice[1]


	qqlength := len(qqmian)

	if qqlength < 10 {
		qqmian =  strings.Repeat("0",10-qqlength) + qqmian
	}

	if isAllNum(qqmian)==false{
		log.Fatal("不是数字")
		return
	}

	pwdlength:= len(qqpassword)
	if pwdlength <=0{
		log.Fatal("没有设置密码")
		return
	}
	flag = true

	isExit(&g_pp,11,qqmian)

	if flag==true{
		log.Fatal("已经存在账号")
		return
	}

	assign(&g_pp,11,qqmian,qqpassword)

}

var file *bufio.Writer
func memtoFile(pp *[]*[]*byte, deep int) {

	//最后一级指针时，开辟内存，存储密码。
	if deep == 1 {
		for i:=0;i<=9;i++{
			if (*pp)[i]!=nil{

				p:= (*string)(unsafe.Pointer((*pp)[i]))
				file.WriteString(*p+"\n")
			}
		}
		return
	}

	//刚开始传递11级指针的地址，deep=11递归下去，可以省略，改为传递11级指针，deep=10。
	if deep == 11 {
		memtoFile((*[]*[]*byte)(unsafe.Pointer(pp)), deep - 1);
		return
	}

	for i:=0;i<=9;i++{
		if (*pp)[i]!=nil{
			memtoFile((*[]*[]*byte)(unsafe.Pointer((*pp)[i])), deep - 1);
		}
	}
}