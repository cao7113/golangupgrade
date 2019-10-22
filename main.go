package main

import (
	"fmt"
	"runtime"

	ptry "github.com/cao7113/golangprivatemod/try"

	try "github.com/cao7113/golang/try"
)

func main() {
	str := fmt.Sprintf("Go version: %s", runtime.Version())
	fmt.Println(str)
	ptry.Try()
	try.Try()
}
