package main

import (
	"fmt"
	"github.com/rock-rabbit/gf/frame/g"
)

func main() {
	g.Log().PrintStack()

	fmt.Println(g.Log().GetStack())
}
