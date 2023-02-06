package main

import (
	"context"
	"fmt"
	"github.com/rock-rabbit/gf/frame/g"
	"github.com/rock-rabbit/gf/util/gvalid"
)

func main() {
	g.I18n().SetLanguage("cn")
	err := gvalid.Check(context.TODO(), "", "required", nil)
	fmt.Println(err.String())
}
