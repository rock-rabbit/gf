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
	v.SetPath("files/template/layout2")
	s, err := v.Parse("layout.html", g.Map{
		"mainTpl": "main/main1.html",
	})
	fmt.Println(err)
	fmt.Println(s)
}
