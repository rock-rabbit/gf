package main

import (
	"github.com/rock-rabbit/gf/frame/g"
	"github.com/rock-rabbit/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindStatusHandlerByMap(map[int]ghttp.HandlerFunc{
		403: func(r *ghttp.Request) { r.Response.Writeln("403") },
		404: func(r *ghttp.Request) { r.Response.Writeln("404") },
		500: func(r *ghttp.Request) { r.Response.Writeln("500") },
	})
	s.SetPort(8199)
	s.Run()
}
