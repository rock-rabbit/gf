package main

import (
	"github.com/rock-rabbit/gf/frame/g"
)

func main() {
	g.DB().Model("user").Distinct().CountColumn("uid,name")
}
