package main

import (
	"github.com/rock-rabbit/gf/frame/g"
	"github.com/rock-rabbit/gf/i18n/gi18n"
	"github.com/rock-rabbit/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(func(r *ghttp.Request) {
			r.SetCtx(gi18n.WithLanguage(r.Context(), r.GetString("lang", "zh-CN")))
			r.Middleware.Next()
		})
		group.ALL("/", func(r *ghttp.Request) {
			r.Response.WriteTplContent(`{#hello}{#world}!`)
		})
	})
	s.SetPort(8199)
	s.Run()
}
