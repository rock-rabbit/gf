package main

import (
	"fmt"

	"github.com/rock-rabbit/gf/text/gstr"
)

func main() {
	fmt.Println(gstr.HideStr("热爱GF热爱生活", 20, "*"))
	fmt.Println(gstr.HideStr("热爱GF热爱生活", 50, "*"))
}
