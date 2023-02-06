package main

import (
	"fmt"

	"github.com/rock-rabbit/gf/frame/g"
	"github.com/rock-rabbit/gf/net/ghttp"
	"github.com/rock-rabbit/gf/os/gres"
	_ "github.com/rock-rabbit/gf/os/gres/testdata/data"
)

func main() {
	gres.Dump()

	//v := g.View()
	//v.SetPath("template/layout1")

	s := g.Server()
	s.SetIndexFolder(true)
	s.SetServerRoot("root")
	s.BindHookHandler("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
		fmt.Println(r.URL.Path, r.IsFileRequest())
	})
	s.BindHandler("/template", func(r *ghttp.Request) {
		r.Response.WriteTpl("layout1/layout.html")
	})
	s.SetPort(8198)
	s.Run()
}
