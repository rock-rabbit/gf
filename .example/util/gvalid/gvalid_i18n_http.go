package main

import (
	"github.com/rock-rabbit/gf/net/ghttp"

	"github.com/rock-rabbit/gf/frame/g"
	"github.com/rock-rabbit/gf/i18n/gi18n"
)

func main() {
	type User struct {
		Name    string `v:"required#ReuiredUserName"`
		Type    int    `v:"required#ReuiredUserType"`
		Project string `v:"size:10#MustSize"`
	}
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(func(r *ghttp.Request) {
			lang := r.GetString("lang", "zh-CN")
			r.SetCtx(gi18n.WithLanguage(r.Context(), lang))
			r.Middleware.Next()
		})
		group.GET("/validate", func(r *ghttp.Request) {
			var (
				err  error
				user = User{}
			)
			if err = r.Parse(&user); err != nil {
				r.Response.WriteExit(err)
			}
			r.Response.WriteExit(user)
		})
	})
	s.SetPort(8199)
}
