package main

import (
	"time"

	"github.com/rock-rabbit/gf/frame/g"
	"github.com/rock-rabbit/gf/net/ghttp"
	"github.com/rock-rabbit/gf/os/gproc"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Writeln("哈喽！")
	})
	s.BindHandler("/pid", func(r *ghttp.Request) {
		r.Response.Writeln(gproc.Pid())
	})
	s.BindHandler("/sleep", func(r *ghttp.Request) {
		r.Response.Writeln(gproc.Pid())
		time.Sleep(10 * time.Second)
		r.Response.Writeln(gproc.Pid())
	})
	s.EnableAdmin()
	s.SetPort(8199)
	s.Run()
}
