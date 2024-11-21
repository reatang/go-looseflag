package main

import (
	"fmt"

	looseflag "github.com/reatang/looseflag"
)

// go run main.go -a str1 -b -int=123 1111111 222222 333333
// go run main.go -a str1 -int=123 -b -- 1111111 222222 333333

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
