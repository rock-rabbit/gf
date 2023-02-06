package main

import (
	"fmt"

	"github.com/rock-rabbit/gf/frame/g"
	"github.com/rock-rabbit/gf/os/gres"
	_ "github.com/rock-rabbit/gf/os/gres/testdata"
)

func main() {
	gres.Dump()

	v := g.View()
	s, err := v.Parse("index.html")
	fmt.Println(err)
	fmt.Println(s)
}
