### file block
下面的代码无效，因为import 是file block 。不能跨文件

```go
// f1.go
package main

import "fmt"
// f2.go  无效
package main

func f() {
  fmt.Println("Hello World")
}
```