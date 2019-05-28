

## reflect的基本功能TypeOf和ValueOf
既然反射就是用来检测存储在接口变量内部(值value；类型concrete type) pair对的一种机制。那么在Golang的reflect反射包中有什么样的方式可以让我们直接获取到变量内部的信息呢？ 它提供了两种类型（或者说两个方法）让我们可以很容易的访问接口变量内容，分别是reflect.ValueOf() 和 reflect.TypeOf()，看看官方的解释
// ValueOf returns a new Value initialized to the concrete value
// stored in the interface i.  ValueOf(nil) returns the zero
func ValueOf(i interface{}) Value {...}

// TypeOf returns the reflection Type that represents the dynamic type of i.
// If i is a nil interface value, TypeOf returns nil.
func TypeOf(i interface{}) Type {...}


## What is reflection?
Reflection is the ability of a program to inspect its variables and values at run time and find their type. You might not understand what this means but that's alright. You will get a clear understanding of reflection by the end of this tutorial, so stay with me.

## What is the need to inspect a variable and find its type?
The first question anyone gets when learning about reflection is why do we even need to inspect a variable and find its type at runtime when each and every variable in our program is defined by us and we know its type at compile time itself. Well this is true most of the times, but not always.

## Let me explain what I mean. Let's write a simple program.
```
package main

import (
    "fmt"
)

func main() {
    i := 10
    fmt.Printf("%d %T", i, i)
}
```

In the program above, the type of i is known at compile time and we print it in the next line. Nothing magical here.

Now let's understand the need to know the type of a variable at run time. Let's say we want to write a simple function which will take a struct as argument and will create a SQL insert query using it.

Consider the following program,

```go
package main

import (
    "fmt"
)

type order struct {
    ordId      int
    customerId int
}

func main() {
    o := order{
        ordId:      1234,
        customerId: 567,
    }
    fmt.Println(o)
}
```
We need to write a function which will take the struct o in the program above as an argument and return the following SQL insert query,

insert into order values(1234, 567)
This function is simple to write. Lets do that now.

```go
package main

import (
    "fmt"
)

type order struct {
    ordId      int
    customerId int
}

func createQuery(o order) string {
    i := fmt.Sprintf("insert into order values(%d, %d)", o.ordId, o.customerId)
    return i
}

func main() {
    o := order{
        ordId:      1234,
        customerId: 567,
    }
    fmt.Println(createQuery(o))
}
```


The createQuery function in line no. 12 creates the insert query by using the ordId and customerId fields of o. This program will output,

```
insert into order values(1234, 567)
```
Now lets take our query creator to the next level. What if we want to generalize our query creator and make it work on any struct. Let me explain what I mean using a program.

```
package main

type order struct {
    ordId      int
    customerId int
}

type employee struct {
    name string
    id int
    address string
    salary int
    country string
}

func createQuery(q interface{}) string {
}

func main() {

}
```
Our objective is to finish the createQuery function in line no. 16 of the above program so that it takes any struct as argument and creates an insert query based on the struct fields.

For example, if we pass the struct below
```
o := order {
    ordId: 1234,
    customerId: 567
}
``

```
Our createQuery function should return,

insert into order values (1234, 567)
```
Similarly if we pass

e := employee {
       name: "Naveen",
       id: 565,
       address: "Science Park Road, Singapore",
       salary: 90000,
       country: "Singapore",
   }
```

it should return,
```
insert into employee values("Naveen", 565, "Science Park Road, Singapore", 90000, "Singapore")
```

Since the createQuery function should work with any struct, it takes a interface{} as argument. For simplicity, we will only deal with structs that contain fields of type string and int but this can be extended for any type.

The createQuery function should work on any struct. The only way to write this function is to examine the type of the struct argument passed to it at run time, find its fields and then create the query. This is where reflection is useful. In the next steps of the tutorial, we will learn how we can achieve this using the reflect package.


## reflect package
The reflect package implements run-time reflection in Go. The reflect package helps to identify the underlying concrete type and the value of a interface{} variable. This is exactly what we need. The createQuery function takes a interface{} argument and the query needs to be created based on the concrete type and value of the interface{} argument. This is exactly what the reflect package helps in doing.
There are a few types and methods in the reflect package which we need to know first before writing our generic query generator program. Lets look at them one by one.

## reflect.Type and reflect.Value
The concrete type of interface{} is represented by reflect.Type and the underlying value is represented by reflect.Value. There are two functions reflect.TypeOf() and reflect.ValueOf() which return the reflect.Type and reflect.Value respectively. These two types are the base to create our query generator. Let's write a simple example to understand these two types.

```go
package main

import (
    "fmt"
    "reflect"
)

type order struct {
    ordId      int
    customerId int
}

func createQuery(q interface{}) {
    t := reflect.TypeOf(q)
    v := reflect.ValueOf(q)
    fmt.Println("Type ", t)
    fmt.Println("Value ", v)


}
func main() {
    o := order{
        ordId:      456,
        customerId: 56,
    }
    createQuery(o)

}

```

In the program above, the createQuery function in line no. 13 takes a interface{} as argument. The function reflect.TypeOf in line no. 14 takes a interface{} as argument and returns the reflect.Type containing the concrete type of the interface{} argument passed. Similarly the reflect.ValueOf function in line no. 15 takes a interface{} as argument and returns the reflect.Value which contains the underlying value of the interface{} argument passed.

The above program prints,

Type  main.order
Value  {456 56}
From the output, we can see that the program prints the concrete type and the value of the interface.

## reflect.Kind
There is one more important type in the reflection package called Kind.

The types Kind and Type in the reflection package might seem similar but they have a difference which will be clear from the program below.

```go
package main

import (
    "fmt"
    "reflect"
)

type order struct {
    ordId      int
    customerId int
}

func createQuery(q interface{}) {
    t := reflect.TypeOf(q)
    k := t.Kind()
    fmt.Println("Type ", t)
    fmt.Println("Kind ", k)


}
func main() {
    o := order{
        ordId:      456,
        customerId: 56,
    }
    createQuery(o)

}
```


The program above outputs,

Type  main.order
Kind  struct
I think you will now be clear about the differences between the two. Type represents the actual type of the interface{}, in this case main.Order and Kind represents the specific kind of the type. In this case, it's a struct.


## NumField() and Field() methods
The NumField() method returns the number of fields in a struct and the Field(i int) method returns the reflect.Value of the ith field.

```go
package main

import (
    "fmt"
    "reflect"
)

type order struct {
    ordId      int
    customerId int
}

func createQuery(q interface{}) {
    if reflect.ValueOf(q).Kind() == reflect.Struct {
        v := reflect.ValueOf(q)
        fmt.Println("Number of fields", v.NumField())
        for i := 0; i < v.NumField(); i++ {
            fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
        }
    }

}
func main() {
    o := order{
        ordId:      456,
        customerId: 56,
    }
    createQuery(o)
}
```

In the program above, in line no. 14 we first check whether the Kind of q is a struct because the NumField method works only on struct. The rest of the program is self explanatory. This program outputs,

Number of fields 2
Field:0 type:reflect.Value value:456
Field:1 type:reflect.Value value:56


## Int() and String() methods
The methods Int and String help extract the reflect.Value as an int64 and string respectively.
```go
package main

import (
    "fmt"
    "reflect"
)

func main() {
    a := 56
    x := reflect.ValueOf(a).Int()
    fmt.Printf("type:%T value:%v\n", x, x)
    b := "Naveen"
    y := reflect.ValueOf(b).String()
    fmt.Printf("type:%T value:%v\n", y, y)

}
```

In the program above, in line no. 10, we extract the reflect.Value as an int64 and in line no. 13, we extract it as string. This program prints,
```
type:int64 value:56
type:string value:Naveen
```

## Complete Program
Now that we have enough knowledge to finish our query generator, lets go ahead and do it.
```go
package main

import (
    "fmt"
    "reflect"
)

type order struct {
    ordId      int
    customerId int
}

type employee struct {
    name    string
    id      int
    address string
    salary  int
    country string
}

func createQuery(q interface{}) {
    if reflect.ValueOf(q).Kind() == reflect.Struct {
        t := reflect.TypeOf(q).Name()
        query := fmt.Sprintf("insert into %s values(", t)
        v := reflect.ValueOf(q)
        for i := 0; i < v.NumField(); i++ {
            switch v.Field(i).Kind() {
            case reflect.Int:
                if i == 0 {
                    query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
                } else {
                    query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
                }
            case reflect.String:
                if i == 0 {
                    query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
                } else {
                    query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
                }
            default:
                fmt.Println("Unsupported type")
                return
            }
        }
        query = fmt.Sprintf("%s)", query)
        fmt.Println(query)
        return

    }
    fmt.Println("unsupported type")
}

func main() {
    o := order{
        ordId:      456,
        customerId: 56,
    }
    createQuery(o)

    e := employee{
        name:    "Naveen",
        id:      565,
        address: "Coimbatore",
        salary:  90000,
        country: "India",
    }
    createQuery(e)
    i := 90
    createQuery(i)

}
```

In line no. 22, we first check whether the passed argument is a struct. In line no. 23 we get the name of the struct from its reflect.Type using the Name() method. In the next line, we use t and start creating the query.

The case statement in line. 28 checks whether the current field is reflect.Int, if that's the case we extract the value of that field as int64 using the Int() method. The if else statement is used to handle edge cases. Please add logs to understand why it is needed. Similar logic is used to extract the string in line no. 34.

We have also added checks to prevent the program from crashing when unsupported types are passed to the createQuery function. The rest of the program is self explanatory. I recommend adding logs at appropriate places and checking their output to understand this program better.

This program prints,
```
insert into order values(456, 56)
insert into employee values("Naveen", 565, "Coimbatore", 90000, "India")
unsupported type
```
I would leave it as an exercise for the reader to add the field names to the output query. Please try changing the program to print query of the format,

insert into order(ordId, customerId) values(456, 56)
Should reflection be used?
Having shown a practical use of reflection, now comes the real question. Should you be using reflection? I would like to quote Rob Pike's proverb on the use of reflection which answers this question.

Clear is better than clever. Reflection is never clear.

Reflection is a very powerful and advanced concept in Go and it should be used with care. It is very difficult to write clear and maintainable code using reflection. It should be avoided wherever possible and should be used only when absolutely necessary.

This brings us to and end of this tutorial. Hope you enjoyed it. Have a good day.

Like my tutorials? Please show your support by donating. Your donations will help me create more awesome tutorials.