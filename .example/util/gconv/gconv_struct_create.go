package main

import (
	"github.com/rock-rabbit/gf/frame/g"
	"github.com/rock-rabbit/gf/util/gconv"
)

func main() {
	type User struct {
		Uid  int
		Name string
	}
	user := (*User)(nil)
	params := g.Map{
		"uid":  1,
		"name": "john",
	}
	err := gconv.Struct(params, &user)
	if err != nil {
		panic(err)
	}
	g.Dump(user)
}
