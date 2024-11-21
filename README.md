# go-looseflag

不兼容go flag，以宽松的方式解析命令行参数的库。

该库主要面向：主命令调用子命令已不可修改，但是需要开发一个适配的新子命令。

## 功能

无须预定义参数，直接获取命令行参数的值

```go
package main

import (
	"fmt"

	looseflag "github.com/reatang/go-looseflag"
)

// go run main.go -a str1 -b -int=123 1111111 222222 333333
// go run main.go -a=str1 -int 123 -b -- 1111111 222222 333333
// go run main.go -a=str1 -int 123 -b 1111111 222222 333333

func main() {
    looseflag.CommandLine.SetBoolArgs("b")
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
