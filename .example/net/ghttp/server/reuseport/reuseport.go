package main

import (
	"github.com/rock-rabbit/gf/frame/g"
	"github.com/rock-rabbit/gf/net/ghttp"
	"github.com/rock-rabbit/gf/os/glog"
)

func main() {
	s1 := ghttp.GetServer("s1")
	s1.SetPort(8882)
	s1.BindHandler("/", func(r *ghttp.Request) {
		glog.Println("s1")
		r.Response.Writeln("s1")
	})
	s1.Start()

	s2 := ghttp.GetServer("s2")
	s2.SetPort(8882)
	s2.BindHandler("/", func(r *ghttp.Request) {
		glog.Println("s2")
		r.Response.Writeln("s2")
	})
	s2.Start()

	g.Wait()
}
