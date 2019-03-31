## 结构体的声明
```
//声明Teacher结构体
type Teacher struct {
	name string
	age  int8
	sex  byte
}
```

## 结构体的定义

### 方式1
//1、var声明方式实例化结构体，初始化方式为：对象.属性=值
```
	var t1 Teacher
  t1.name = "jonson"
  t1.age = 35
  t1.sex = 1
```
### 方式2

变量简短声明格式实例化结构体，初始化方式为：对象.属性=值
```
t2 := Teacher{}
t2.name = "olaya"
t2.age = 30
t2.sex = 1
```
### 方式3
3、变量简短声明格式实例化结构体，声明时初始化。初始化方式为：属性:值 。属性:值可以同行，也可以换行。（类似map的用法）
```
t3 := Teacher{
  name: "Josh",
  age:  28,
  sex:  1,
}
t3 = Teacher{name: "Josh2", age: 27, sex: 1}
```
### 方式4
变量简短声明格式实例化结构体，声明时初始化，不写属性名，按属性顺序只写属性值
```
t4 := Teacher{"Ruby", 30, 0}
```
### 方式5
创建指针类型的结构体
```go
t5 := new(Teacher)
(*t5).name = "Running"
(*t5).age = 31
(*t5).sex = 0
t5.name = "Running2"
t5.age = 31
t5.sex = 0
```
### 方式6
匿名结构体
```go
addr := struct {
  province, city string
}{"陕西省", "西安市"}
```
