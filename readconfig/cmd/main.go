package main

import (
	"fmt"
	"readconfig/internal"
)

func main() {
	x, _ := internal.ReadVal()
	fmt.Println(x)
}
