# go-looseflag

不兼容go flag，以宽松的方式解析命令行参数的库。Incompatible with go flag, library that parses command-line parameters in a loose manner

## 功能

无须预定义参数，直接获取命令行参数的值

```go
@ -0,0 +1,31 @@
package main

import (
	"fmt"

	looseflag "github.com/reatang/go-looseflag"
)

// go run main.go -a str1 -b -int=123 1111111 222222 333333
// go run main.go -a=str1 -int 123 -b -- 1111111 222222 333333

func main() {
	err := looseflag.Parse()
	if err != nil {
		panic(err)
	}

	if v, ok := looseflag.GetString("a"); ok {
		fmt.Printf("%v\n", v)
	}

	if v, ok := looseflag.GetInt("int"); ok {
		fmt.Printf("%v\n", v)
	}

	if v, ok := looseflag.GetBool("b"); ok {
		fmt.Printf("%v\n", v)
	}

	fmt.Println(looseflag.Args())
}

```

#### 碎碎念

这是一个写golang compiler代理被逼疯后的产物.
