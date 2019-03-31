/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import "fmt"

type Income interface {
	calculate() float64 //计算收入总额
	source() string     //用来说明收入来源
}

//固定账单项目
type FixedBilling struct {
	projectName  string  //工程项目
	biddedAmount float64 //项目招标总额
}

//定时生产项目(定时和材料项目)
type TimeAndMaterial struct {
	projectName string
	workHours   float64 //工作时长
	hourlyRate  float64 //每小时工资率
}

//固定收入项目
func (f FixedBilling) calculate() float64 {
	return f.biddedAmount
}

func (f FixedBilling) source() string {
	return f.projectName
}

//定时收入项目
func (t TimeAndMaterial) calculate() float64 {
	return t.workHours * t.hourlyRate
}

func (t TimeAndMaterial) source() string {
	return t.projectName
}

//通过广告点击获得收入
type Advertisement struct {
	adName         string
	clickCount     int
	incomePerclick float64
}

func (a Advertisement) calculate() float64 {
	return float64(a.clickCount) * a.incomePerclick
}

func (a Advertisement) source() string {
	return a.adName
}

func main() {
	p1 := FixedBilling{"项目1", 5000}
	p2 := FixedBilling{"项目2", 10000}
	p3 := TimeAndMaterial{"项目3", 100, 40}
	p4 := TimeAndMaterial{"项目4", 250, 20}
	p5 := Advertisement{"广告1", 10000, 0.1}
	p6 := Advertisement{"广告2", 20000, 0.05}

	ic := []Income{p1, p2, p3, p4, p5, p6}
	fmt.Println(calculateNetIncome(ic))
}

//计算净收入
func calculateNetIncome(ic []Income) float64 {
	netincome := 0.0
	for _, income := range ic {
		fmt.Printf("收入来源：%s ，收入金额：%.2f \n", income.source(), income.calculate())
		netincome += income.calculate()
	}
	return netincome
}

//说明：
// 没有对calculateNetIncome函数做任何更改，尽管添加了新的收入方式。全靠多态性而起作用。
// 由于新的Advertisement类型也实现了Income接口，可以将它添加到ic切片中。
// calculateNetIncome函数在没有任何更改的情况下工作，因为它可以调用Advertisement类型的calculate()和source()方法。
