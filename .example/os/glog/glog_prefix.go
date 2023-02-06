package main

import (
	"github.com/rock-rabbit/gf/frame/g"
)

func main() {
	g.Log().SetPrefix("[API]")
	g.Log().Println("hello world")
	g.Log().Error("error occurred")
}
