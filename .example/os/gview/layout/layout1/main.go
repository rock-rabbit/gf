package main

import (
	"github.com/rock-rabbit/gf/frame/g"
	"github.com/rock-rabbit/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.WriteTpl("layout.html", g.Map{
			"header":    "This is header",
			"container": "This is container",
			"footer":    "This is footer",
		})
	})
	s.SetPort(8199)
	s.Run()
}
