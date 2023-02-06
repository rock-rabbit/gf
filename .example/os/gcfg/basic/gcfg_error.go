package main

import (
	"fmt"

	"github.com/rock-rabbit/gf/frame/g"
)

func main() {
	fmt.Println(g.Config().Get("none"))
}
