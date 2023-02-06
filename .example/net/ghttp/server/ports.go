package main

import (
	"github.com/rock-rabbit/gf/net/ghttp"
)

func main() {
	s := ghttp.GetServer()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Writeln("go frame!")
	})
	s.SetPort(8100, 8200, 8300)
	s.Run()
}
