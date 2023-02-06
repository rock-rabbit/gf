package main

import (
	_ "github.com/rock-rabbit/gf/os/gres/testdata/example/boot"

	"github.com/rock-rabbit/gf/frame/g"
	"github.com/rock-rabbit/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/template", func(r *ghttp.Request) {
			r.Response.WriteTplDefault(g.Map{
				"name": "GoFrame",
			})
		})
	})
	s.Run()
}
